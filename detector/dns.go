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
	"net"
	"strings"
	"time"

	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/dnsdetector"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/go-parallel"
	"github.com/vicanso/hes"
)

type (
	DNSSrv struct{}
)

// check dns check
func (srv *DNSSrv) check(host, server string, timeout time.Duration) ([]net.IPAddr, error) {
	if !strings.Contains(server, ":") {
		server += ":53"
	}
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dia := net.Dialer{}
			return dia.DialContext(ctx, "udp", server)
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return r.LookupIPAddr(ctx, host)
}

// detect dns detect
func (srv *DNSSrv) detect(config *ent.DNSDetector) (dnsDetectorResult *ent.DNSDetectorResult, err error) {
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
		addrs, err := srv.check(config.Host, server, timeout)
		subResult := schema.DNSDetectorSubResult{
			Server:   server,
			Duration: int(time.Since(startedAt).Milliseconds()),
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
				err = hes.New("no ip for the host")
			}
			// 检测IP是否均在预期列表中
			for _, ip := range subResult.IPS {
				if !util.ContainsString(config.Ips, ip) {
					err = hes.New(fmt.Sprintf("ip(%s) is not matched", ip))
				}
			}
		}

		if err != nil {
			subResult.Result = schema.DetectorResultFail
			subResult.Message = err.Error()
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
	return getEntClient().DNSDetectorResult.Create().
		SetTask(config.ID).
		SetResult(int8(result)).
		SetHost(config.Host).
		SetResults(subResults).
		SetMaxDuration(maxDuration).
		SetMessages(messages).
		Save(ctx)
}

func (srv *DNSSrv) doAlarm(name string, receivers []string, result *ent.DNSDetectorResult) {
	// 如果无结果，忽略
	if result == nil {
		return
	}
	doAlarm(alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("dns-%d", result.Task),
		IsSuccess: result.Result == int8(schema.DetectorResultSuccess),
		Messages:  result.Messages,
	})
}

// Detect do dns detect
func (srv *DNSSrv) Detect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	result, err := getEntClient().DNSDetector.Query().
		Where(dnsdetector.StatusEQ(schema.StatusEnabled)).
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
		err = convertParallelError(pErr, "dns detect fail")
	}
	if err != nil {
		return
	}
	return
}
