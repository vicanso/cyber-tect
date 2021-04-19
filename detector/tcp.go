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
	"fmt"
	"strings"
	"time"

	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/tcpdetector"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/go-parallel"
)

type (
	TCPSrv struct{}
)

// check tcp check
func (srv *TCPSrv) check(addr string, timeout time.Duration) error {
	return portCheck("", addr, timeout)
}

// detect tcp detect
func (srv *TCPSrv) detect(config *ent.TCPDetector) (tcpDetectorResult *ent.TCPDetectorResult, err error) {
	timeout, _ := time.ParseDuration(config.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	result := schema.DetectorResultSuccess
	subResults := make(schema.TCPDetectorSubResults, 0)
	maxDuration := 0
	messages := make([]string, 0)

	for _, addr := range config.Addrs {
		startedAt := time.Now()
		err = srv.check(addr, timeout)
		subResult := schema.TCPDetectorSubResult{
			Addr:     addr,
			Duration: int(time.Since(startedAt).Milliseconds()),
		}
		if err != nil {
			subResult.Result = schema.DetectorResultFail
			subResult.Message = fmt.Sprintf("%s, %s", addr, err.Error())
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
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return getEntClient().TCPDetectorResult.Create().
		SetTask(config.ID).
		SetResult(int8(result)).
		SetAddrs(strings.Join(config.Addrs, ",")).
		SetMaxDuration(maxDuration).
		SetResults(subResults).
		SetMessages(messages).
		Save(ctx)
}

func (srv *TCPSrv) doAlarm(name string, receivers []string, result *ent.TCPDetectorResult) {
	// 如果无结果，忽略
	if result == nil {
		return
	}
	doAlarm(alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("tcp-%d", result.Task),
		IsSuccess: result.Result == int8(schema.DetectorResultSuccess),
		Messages:  result.Messages,
	})
}

// Detect do tcp detect
func (srv *TCPSrv) Detect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	result, err := getEntClient().TCPDetector.Query().
		Where(tcpdetector.StatusEQ(schema.StatusEnabled)).
		All(ctx)
	if err != nil {
		return
	}

	pErr := parallel.Parallel(len(result), detectorConfig.Concurrency, func(index int) error {
		item := result[index]
		detectResult, err := srv.detect(item)
		srv.doAlarm(item.Name, item.Receivers, detectResult)
		return err
	})
	// 如果parallel检测失败，则转换为http error
	if pErr != nil {
		err = convertParallelError(pErr, "tcp detect fail")
	}
	if err != nil {
		return
	}
	return
}
