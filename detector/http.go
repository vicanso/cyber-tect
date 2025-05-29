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
	"compress/gzip"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	parallel "github.com/vicanso/go-parallel"
	"github.com/vicanso/hes"
	HT "github.com/vicanso/http-trace"
	"golang.org/x/net/http/httpproxy"
)

const userAgentChrome = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36"
const acceptEncodingChrome = "gzip"

// warningCertExpiredDuration 证书过期告警时间（7天）
var warningCertExpiredDuration = 7 * 24 * time.Hour

// 未指定IP时使用
var nilIPAddr = "0.0.0.0"

type (
	HTTPSrv         struct{}
	httpCheckParams struct {
		url     string
		ip      string
		timeout time.Duration
		script  string
		proxy   string
	}
	httpCheckResult struct {
		ht   *HT.HTTPTrace
		body []byte
	}
)

// check 执行一次http检测
func (srv *HTTPSrv) check(ctx context.Context, params httpCheckParams) (result *httpCheckResult, err error) {
	ip := params.ip
	requestURL := params.url
	timeout := params.timeout
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var dialContext func(ctx context.Context, network, addr string) (net.Conn, error)
	// 自定义dns解析（更新为0.0.0.0)表示不指定IP
	if ip != "" && ip != nilIPAddr {
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
	transport := &http.Transport{
		ForceAttemptHTTP2: true,
		// 设置较短时间，不复用
		IdleConnTimeout:       1 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DialContext:           dialContext,
	}
	// 如果有定义代理，则通过代理检测服务
	if params.proxy != "" {
		fn := (&httpproxy.Config{
			HTTPProxy:  params.proxy,
			HTTPSProxy: params.proxy,
		}).ProxyFunc()
		transport.Proxy = func(r *http.Request) (*url.URL, error) {
			return fn(r.URL)
		}
	}
	client := &http.Client{
		Timeout: timeout,
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", userAgentChrome)
	req.Header.Set("Accept-Encoding", acceptEncodingChrome)
	trace, ht := HT.NewClientTrace()
	result = &httpCheckResult{
		ht: ht,
	}
	defer ht.Finish()
	ctx = httptrace.WithClientTrace(ctx, trace)
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var r io.Reader
	if resp.Header.Get("Content-Encoding") == "gzip" {
		r, _ = gzip.NewReader(resp.Body)
	} else {
		r = resp.Body
	}
	if ht.Protocol == "" {
		ht.Protocol = resp.Proto
	}
	ht.Protocol = strings.ToUpper(ht.Protocol)
	buf, _ := io.ReadAll(r)

	result.body = buf
	// < 200 或者 >= 400 均认为失败
	if resp.StatusCode >= http.StatusBadRequest || resp.StatusCode < http.StatusOK {
		err = &hes.Error{
			StatusCode: resp.StatusCode,
			Message:    util.CutRune(string(buf), 500),
		}
		return
	}
	if len(ht.Certificates) != 0 {
		endDate := ht.Certificates[0].NotAfter
		// 如果证书准备过期，设置为失败
		if endDate.UnixNano() < time.Now().UnixNano()+warningCertExpiredDuration.Nanoseconds() {
			err = fmt.Errorf("证书将于%s过期", endDate.String())
			return
		}
	}

	// 如果未配置检测脚本，则直接返回
	if params.script == "" {
		return
	}

	vm := goja.New()
	var value interface{}
	// 如果响应数据是json，转换为map
	if strings.Contains(resp.Header.Get("Content-Type"), "json") {
		m := make(map[string]interface{})
		// 转换错误忽略
		err = json.Unmarshal(buf, &m)
		if err != nil {
			return
		}
		value = m
	} else {
		value = string(buf)
	}
	err = vm.Set("resp", value)
	if err != nil {
		return
	}
	script := fmt.Sprintf(`(function(){
		%s
	})(resp);`, params.script)
	_, err = vm.RunString(script)
	if err != nil {
		return
	}

	return
}

// fillSubResult 填充相关检测结果
func (srv *HTTPSrv) fillSubResult(ht *HT.HTTPTrace, subResult *schema.HTTPDetectorSubResult) {
	stats := ht.Stats()
	subResult.Duration = ceilToMs(stats.Total)
	subResult.DNSLookup = ceilToMs(stats.DNSLookup)
	subResult.TCPConnection = ceilToMs(stats.TCPConnection)
	subResult.TLSHandshake = ceilToMs(stats.TLSHandshake)
	subResult.ServerProcessing = ceilToMs(stats.ServerProcessing)
	subResult.ContentTransfer = ceilToMs(stats.ContentTransfer)
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
	}
}

// detect http detect
func (srv *HTTPSrv) detect(ctx context.Context, config *ent.HTTPDetector) (httpDetectorResult *ent.HTTPDetectorResult, err error) {
	timeout, _ := time.ParseDuration(config.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	result := schema.DetectorResultSuccess
	subResults := make(schema.HTTPDetectorSubResults, 0)
	maxDuration := 0
	messages := make([]string, 0)

	type httpCheck struct {
		ip    string
		proxy string
	}
	checkList := make([]httpCheck, 0)
	for _, ip := range config.Ips {
		if len(config.Proxies) == 0 {
			checkList = append(checkList, httpCheck{
				ip: ip,
			})
			continue
		}

		for _, proxy := range config.Proxies {
			checkList = append(checkList, httpCheck{
				ip:    ip,
				proxy: proxy,
			})
		}
	}

	for _, check := range checkList {
		ip := check.ip
		proxy := check.proxy
		url := config.URL
		if config.RandomQueryString == 1 {
			str := util.GenXID()
			if strings.Contains(url, "?") {
				url += fmt.Sprintf("&%s", str)
			} else {
				url += fmt.Sprintf("?%s", str)
			}
		}
		checkResult, err := srv.check(ctx, httpCheckParams{
			url:     url,
			ip:      ip,
			timeout: timeout,
			script:  config.Script,
			proxy:   proxy,
		})
		subResult := schema.HTTPDetectorSubResult{
			Addr:  ip,
			Proxy: proxy,
		}
		var ht *HT.HTTPTrace
		if checkResult != nil {
			ht = checkResult.ht
			m := md5.New()
			m.Write(checkResult.body)
			subResult.Hash = hex.EncodeToString(m.Sum(nil))
		}
		if err != nil {
			subResult.Result = schema.DetectorResultFail
			if (ip == "" || ip == nilIPAddr) && ht != nil {
				ip = ht.Addr
			}
			if ip == "" && len(ht.Addrs) != 0 {
				ip = ht.Addrs[0]
			}
			subResult.Message = fmt.Sprintf("%s, %s(%s)", err.Error(), url, ip)
			if ht != nil {
				subResult.Message += fmt.Sprintf(" stats=%s", ht.Stats().String())
			}
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
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	return getEntClient().HTTPDetectorResult.Create().
		SetTask(config.ID).
		SetURL(config.URL).
		SetResult(schema.DetectorResult(result)).
		SetResults(subResults).
		SetMaxDuration(maxDuration).
		SetMessages(messages).
		Save(ctx)
}

func (srv *HTTPSrv) doAlarm(ctx context.Context, name string, receivers []string, result *ent.HTTPDetectorResult) {
	// 如果无结果，忽略
	if result == nil {
		return
	}
	doAlarm(ctx, alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("http-%d", result.Task),
		IsSuccess: result.Result == schema.DetectorResultSuccess,
		Messages:  result.Messages,
	})
}

// Detect do http detect
func (srv *HTTPSrv) Detect(ctx context.Context, count int64) (err error) {
	result, err := getEntClient().HTTPDetector.Query().
		Where(httpdetector.StatusEQ(schema.StatusEnabled)).
		All(ctx)
	if err != nil {
		return
	}
	// 如果未配置检测，则直接返回
	if len(result) == 0 {
		return nil
	}
	pErr := parallel.Parallel(func(index int) error {
		item := result[index]
		if !shouldDoDetect(item.Interval, count) {
			return nil
		}
		detectResult, err := srv.detect(ctx, item)
		srv.doAlarm(ctx, item.Name, item.Receivers, detectResult)
		return err
	}, len(result), detectorConfig.Concurrency)
	// 如果parallel检测失败，则转换为http error
	if pErr != nil {
		err = convertParallelError(pErr, "http detect fail")
	}
	if err != nil {
		return
	}

	return
}
