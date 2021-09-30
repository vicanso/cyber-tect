// Copyright 2021 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package detector

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-ping/ping"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/pingdetector"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/go-parallel"
)

type (
	PingSrv struct{}
)

// check ping check
func (srv *PingSrv) check(ctx context.Context, ip string, timeout time.Duration) (err error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	pinger.Timeout = timeout
	pinger.Interval = 100 * time.Millisecond
	err = pinger.Run()
	if err != nil {
		return
	}
	loss := int(pinger.Statistics().PacketLoss)
	if loss > 50 {
		msg := fmt.Sprintf("Too many packets are lost, loss:%d%%", loss)
		err = errors.New(msg)
		return
	}
	return
}

// detect ping detect
func (srv *PingSrv) detect(ctx context.Context, config *ent.PingDetector) (pingDetectorResult *ent.PingDetectorResult, err error) {
	timeout, _ := time.ParseDuration(config.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	result := schema.DetectorResultSuccess
	subResults := make(schema.PingDetectorSubResults, 0)
	maxDuration := 0
	messages := make([]string, 0)

	for _, ip := range config.Ips {
		startedAt := time.Now()
		err = srv.check(ctx, ip, timeout)
		subResult := schema.PingDetectorSubResult{
			IP:       ip,
			Duration: int(time.Since(startedAt).Milliseconds()),
		}
		if err != nil {
			subResult.Result = schema.DetectorResultFail
			subResult.Message = fmt.Sprintf("%s, %s", ip, err.Error())
			result = schema.DetectorResultFail
			messages = append(messages, subResult.Message)
		} else {
			subResult.Result = schema.DetectorResultSuccess
		}
		if subResult.Duration > maxDuration {
			maxDuration = subResult.Duration
		}
		subResults = append(subResults, &subResult)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	return getEntClient().PingDetectorResult.Create().
		SetTask(config.ID).
		SetResult(schema.DetectorResult(result)).
		SetIps(strings.Join(config.Ips, ",")).
		SetMaxDuration(maxDuration).
		SetResults(subResults).
		SetMessages(messages).
		Save(ctx)
}

func (srv *PingSrv) doAlarm(ctx context.Context, name string, receivers []string, result *ent.PingDetectorResult) {
	// 如果无结果，忽略
	if result == nil {
		return
	}
	doAlarm(ctx, alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("ping-%d", result.Task),
		IsSuccess: result.Result == schema.DetectorResultSuccess,
		Messages:  result.Messages,
	})
}

// Detect do ping detect
func (srv *PingSrv) Detect(ctx context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	result, err := getEntClient().PingDetector.Query().
		Where(pingdetector.StatusEQ(schema.StatusEnabled)).
		All(ctx)
	if err != nil {
		return
	}
	pErr := parallel.Parallel(func(index int) error {
		item := result[index]
		detectResult, err := srv.detect(ctx, item)
		srv.doAlarm(ctx, item.Name, item.Receivers, detectResult)
		return err
	}, len(result), detectorConfig.Concurrency)
	// 如果parallel检测失败，则转换为http error
	if pErr != nil {
		err = convertParallelError(pErr, "ping detect fail")
	}
	if err != nil {
		return
	}
	return
}
