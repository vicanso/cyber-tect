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
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"

	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/util"
	parallel "github.com/vicanso/go-parallel"
	"github.com/vicanso/hes"
	HT "github.com/vicanso/http-trace"
)

const userAgentChrome = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36"
const acceptEncodingChrome = "gzip"

// warningCertExpiredDuration 证书过期告警时间（7天）
var warningCertExpiredDuration = 7 * 24 * time.Hour

type (
	HTTPSrv struct{}
)

// check 执行一次http检测
func (srv *HTTPSrv) check(url, ip string, timeout time.Duration) (ht *HT.HTTPTrace, err error) {
	var dialContext func(ctx context.Context, network, addr string) (net.Conn, error)
	// 自定义dns解析（更新为0.0.0.0)表示不指定IP
	if ip != "" && ip != "0.0.0.0" {
		dialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			dia := &net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}
			// IPV6
			if strings.Contains(ip, ":") {
				ip = "[" + ip + "]"
			}
			sepIndex := strings.LastIndex(addr, ":")
			return dia.DialContext(ctx, network, ip+addr[sepIndex:])
		}
	}
	// 每次都使用新的client，避免复用
	client := &http.Client{
		Timeout: timeout,
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			ForceAttemptHTTP2: true,
			// 设置较短时间，不复用
			IdleConnTimeout:       1 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DialContext:           dialContext,
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", userAgentChrome)
	req.Header.Set("Accept-Encoding", acceptEncodingChrome)
	trace, ht := HT.NewClientTrace()
	defer ht.Finish()
	ctx := context.Background()
	ctx = httptrace.WithClientTrace(ctx, trace)
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest {
		err = &hes.Error{
			StatusCode: resp.StatusCode,
			Message:    util.CutRune(string(buf), 500),
		}
		return
	}
	return
}

// fillSubResult 填充相关检测结果
func (srv *HTTPSrv) fillSubResult(ht *HT.HTTPTrace, subResult *schema.HTTPDetectorSubResult) {
	stats := ht.Stats()
	subResult.Duration = int(stats.Total.Milliseconds())
	subResult.DNSLookup = int(stats.DNSLookup.Milliseconds())
	subResult.TCPConnection = int(stats.TCPConnection.Milliseconds())
	subResult.TLSHandshake = int(stats.TLSHandshake.Milliseconds())
	subResult.ServerProcessing = int(stats.ServerProcessing.Milliseconds())
	subResult.ContentTransfer = int(stats.ContentTransfer.Milliseconds())
	if ht.Addr != "" {
		subResult.Addr = ht.Addr
	}
	subResult.Addrs = ht.Addrs
	subResult.Protocol = ht.Protocol
	subResult.TLSVersion = ht.TLSVersion
	subResult.TLSCipherSuite = ht.TLSCipherSuite
	if len(ht.Certificates) != 0 {
		subResult.CertificateDNSNames = ht.Certificates[0].DNSNames
		endDate := ht.Certificates[0].NotAfter
		subResult.CertificateExpirationDates = []string{
			ht.Certificates[0].NotBefore.String(),
			endDate.String(),
		}
		// 如果证书准备过期，设置为失败
		if endDate.UnixNano() < time.Now().UnixNano()+warningCertExpiredDuration.Nanoseconds() {
			subResult.Result = schema.DetectorResultFail
			subResult.Message = fmt.Sprintf("证书将于%s过期", endDate.String())
		}
	}
}

// detect http detect
func (srv *HTTPSrv) detect(config *ent.HTTPDetector) (httpDetectorResult *ent.HTTPDetectorResult, err error) {
	timeout, _ := time.ParseDuration(config.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	result := schema.DetectorResultSuccess
	subResults := make(schema.HTTPDetectorSubResults, 0)
	maxDuration := 0
	messages := make([]string, 0)

	for _, ip := range config.Ips {
		ht, err := srv.check(config.URL, ip, timeout)
		subResult := schema.HTTPDetectorSubResult{
			Addr: ip,
		}
		if err != nil {
			subResult.Result = schema.DetectorResultFail
			subResult.Message = err.Error()
			result = schema.DetectorResultFail
			messages = append(messages, subResult.Message)
		} else {
			subResult.Result = schema.DetectorResultSuccess
		}
		if ht != nil {
			srv.fillSubResult(ht, &subResult)
		}
		if subResult.Duration > maxDuration {
			maxDuration = subResult.Duration
		}
		subResults = append(subResults, &subResult)
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return getEntClient().HTTPDetectorResult.Create().
		SetTask(config.ID).
		SetURL(config.URL).
		SetResult(int8(result)).
		SetResults(subResults).
		SetMaxDuration(maxDuration).
		SetMessages(messages).
		Save(ctx)
}

func (srv *HTTPSrv) doAlarm(name string, receivers []string, result *ent.HTTPDetectorResult) {
	// 如果无结果，忽略
	if result == nil {
		return
	}
	doAlarm(alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("http-%d", result.Task),
		IsSuccess: result.Result == int8(schema.DetectorResultSuccess),
		Messages:  result.Messages,
	})
}

// Detect do http detect
func (srv *HTTPSrv) Detect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	result, err := getEntClient().HTTPDetector.Query().
		Where(httpdetector.StatusEQ(schema.StatusEnabled)).
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
		err = convertParallelError(pErr, "http detect fail")
	}
	if err != nil {
		return
	}

	return
}
