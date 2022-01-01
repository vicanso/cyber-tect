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
	"net"
	"strings"
	"time"

	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/dnsdetector"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	parallel "github.com/vicanso/go-parallel"
)

type (
	DNSSrv struct{}
)

// check dns check
func (srv *DNSSrv) check(ctx context.Context, host, server string, timeout time.Duration) ([]net.IPAddr, error) {
	if !strings.Contains(server, ":") {
		server += ":53"
	}
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, _ string) (net.Conn, error) {
			dia := net.Dialer{}
			return dia.DialContext(ctx, network, server)
		},
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return r.LookupIPAddr(ctx, host)
}

// detect dns detect
func (srv *DNSSrv) detect(ctx context.Context, config *ent.DNSDetector) (*ent.DNSDetectorResult, error) {
	timeout, _ := time.ParseDuration(config.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	result := schema.DetectorResultSuccess
	subResults := make(schema.DNSDetectorSubResults, 0)
	maxDuration := 0
	messages := make([]string, 0)

	for _, server := range config.Servers {
		startedAt := time.Now()
		addrs, err := srv.check(ctx, config.Host, server, timeout)
		subResult := schema.DNSDetectorSubResult{
			Server:   server,
			Duration: ceilToMs(time.Since(startedAt)),
		}
		// 如果检测成功
		if err == nil {
			if len(addrs) != 0 {
				ipList := make([]string, len(addrs))
				for index, addr := range addrs {
					ipList[index] = addr.String()
				}
				subResult.IPS = ipList
			} else {
				err = errors.New("no ip for the host")
			}
			// 检测IP是否均在预期列表中
			for _, ip := range subResult.IPS {
				if !util.ContainsString(config.Ips, ip) {
					err = fmt.Errorf("ip(%s) is not matched", ip)
				}
			}
		}

		if err != nil {
			subResult.Result = schema.DetectorResultFail
			subResult.Message = fmt.Sprintf("%s(%s), %s", config.Host, server, err.Error())
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
	return getEntClient().DNSDetectorResult.Create().
		SetTask(config.ID).
		SetResult(schema.DetectorResult(result)).
		SetHost(config.Host).
		SetResults(subResults).
		SetMaxDuration(maxDuration).
		SetMessages(messages).
		Save(ctx)
}

func (srv *DNSSrv) doAlarm(ctx context.Context, name string, receivers []string, result *ent.DNSDetectorResult) {
	// 如果无结果，忽略
	if result == nil {
		return
	}
	doAlarm(ctx, alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("dns-%d", result.Task),
		IsSuccess: result.Result == schema.DetectorResultSuccess,
		Messages:  result.Messages,
	})
}

// Detect do dns detect
func (srv *DNSSrv) Detect(ctx context.Context, count int64) error {
	result, err := getEntClient().DNSDetector.Query().
		Where(dnsdetector.StatusEQ(schema.StatusEnabled)).
		All(ctx)
	if err != nil {
		return err
	}

	pErr := parallel.Parallel(func(index int) error {
		item := result[index]
		detectResult, err := srv.detect(ctx, item)
		srv.doAlarm(ctx, item.Name, item.Receivers, detectResult)
		return err
	}, len(result), detectorConfig.Concurrency)
	// 如果parallel检测失败，则转换
	if pErr != nil {
		err = convertParallelError(pErr, "dns detect fail")
	}
	if err != nil {
		return err
	}
	return nil
}
