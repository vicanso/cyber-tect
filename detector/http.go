// Copyright 2020 tree xie
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
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"

	"github.com/vicanso/cybertect/helper"

	"github.com/lib/pq"
	HT "github.com/vicanso/http-trace"
	"go.uber.org/zap"
)

const userAgentChrome = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36"
const acceptEncodingChrome = "gzip, deflate, br"

type (
	HTTP struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Owner       string         `json:"owner,omitempty" gorm:"index:idx_http_owner"`
		Status      int            `json:"status,omitempty" gorm:"index:idx_http_status"`
		Description string         `json:"description,omitempty"`
		Receivers   pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`

		// IP 如果配置此IP，则将dns对域名解析为此IP
		IP      string `json:"ip,omitempty"`
		URL     string `json:"url,omitempty"`
		Timeout string `json:"timeout,omitempty"`
	}

	HTTPDetectResult struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Receivers pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`
		Duration  int            `json:"duration,omitempty"`
		Result    int            `json:"result,omitempty" gorm:"index:idx_http_detect_result_result"`
		Message   string         `json:"message,omitempty"`

		Task       uint   `json:"task,omitempty" gorm:"index:idx_http_detect_result_task"`
		IP         string `json:"ip,omitempty"`
		URL        string `json:"url,omitempty"`
		StatusCode int    `json:"statusCode,omitempty"`

		Addrs                      pq.StringArray `json:"addrs,omitempty" gorm:"type:text[]"`
		Addr                       string         `json:"addr,omitempty"`
		Protocol                   string         `json:"protocol,omitempty"`
		TLSVersion                 string         `json:"tlsVersion,omitempty"`
		TLSCipherSuite             string         `json:"tlsCipherSuite,omitempty"`
		CertificateDNSNames        pq.StringArray `json:"certificateDNSNames,omitempty" gorm:"type:text[]"`
		CertificateExpirationDates pq.StringArray `json:"certificateExpirationDates,omitempty" gorm:"type:text[]"`

		DNSLookup        int `json:"dnsLookup,omitempty"`
		TCPConnection    int `json:"tcpConnection,omitempty"`
		TLSHandshake     int `json:"tlsHandshake,omitempty"`
		ServerProcessing int `json:"serverProcessing,omitempty"`
		ContentTransfer  int `json:"contentTransfer,omitempty"`
	}

	HTTPSrv struct{}
)

func init() {
	pgGetClient().AutoMigrate(&HTTP{}).
		AutoMigrate(&HTTPDetectResult{})
}

func (h *HTTP) Check() (resp *http.Response, ht *HT.HTTPTrace, err error) {
	timeout, _ := time.ParseDuration(h.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}

	var dialContext func(ctx context.Context, network, addr string) (net.Conn, error)
	// 自定义dns解析（更新为0.0.0.0)表示不指定IP
	if h.IP != "" {
		dialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			dia := &net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}
			sepIndex := strings.LastIndex(addr, ":")
			return dia.DialContext(ctx, network, h.IP+addr[sepIndex:])
		}
	}
	// 每次都使用新的client，避免复用
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
			// 设置较短时间，不复用
			IdleConnTimeout:       1 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DialContext:           dialContext,
		},
	}

	req, err := http.NewRequest(http.MethodGet, h.URL, nil)
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

	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest {
		body := string(buf)
		if len(body) > 500 {
			body = body[:500]
		}
		err = errors.New("http response with status:" + resp.Status + " body:" + body)
		return
	}
	return
}

// Add add http detector
func (srv *HTTPSrv) Add(h *HTTP) (err error) {
	err = pgCreate(h)
	return
}

// UpdateByID update the http detector
func (srv *HTTPSrv) UpdateByID(id uint, attrs ...interface{}) (err error) {
	err = pgGetClient().Model(&HTTP{
		ID: id,
	}).Update(attrs...).Error
	return
}

// FindByID find by the http detector by id
func (srv *HTTPSrv) FindByID(id uint) (data *HTTP, err error) {
	data = &HTTP{}
	err = pgGetClient().Where("id = ?", id).First(data).Error
	if err != nil {
		return
	}
	return
}

// List list the http detector
func (srv *HTTPSrv) List(params helper.PGQueryParams, args ...interface{}) (data []*HTTP, err error) {
	data = make([]*HTTP, 0)
	err = pgQuery(params, args...).Find(&data).Error
	return
}

// Count count the http detector
func (srv *HTTPSrv) Count(args ...interface{}) (count int, err error) {
	return pgCount(&HTTP{}, args...)
}

// ListResult list the http detect result
func (srv *HTTPSrv) ListResult(params helper.PGQueryParams, args ...interface{}) (data []*HTTPDetectResult, err error) {
	data = make([]*HTTPDetectResult, 0)
	err = pgQuery(params, args...).Find(&data).Error
	return
}

// CountResult count the http detect result
func (srv *HTTPSrv) CountResult(args ...interface{}) (count int, err error) {
	return pgCount(&HTTPDetectResult{}, args...)
}

// Detect do the http detect
func (srv *HTTPSrv) Detect() {
	result := make([]*HTTP, 0)
	err := pgGetClient().Where("status = ?", StatusEnabled).Find(&result).Error
	if err != nil {
		logger.Error("get http detector fail",
			zap.Error(err),
		)
	}
	for _, h := range result {
		go srv.detectOne(h)
	}
}

// ValidateOwner validate the owner
func (srv *HTTPSrv) ValidateOwner(id uint, owner string) (err error) {
	data, err := srv.FindByID(id)
	if err != nil {
		return
	}
	if data.Owner != owner {
		err = errOwnerInvalid
	}
	return
}

func (srv *HTTPSrv) detectOne(h *HTTP) {
	result := HTTPDetectResult{
		IP:   h.IP,
		URL:  h.URL,
		Task: h.ID,
	}
	startedAt := time.Now()
	resp, ht, err := h.Check()
	duration := int(time.Since(startedAt).Milliseconds())
	if err != nil {
		result.Result = DetectFail
		result.Message = err.Error()
	} else {
		result.Result = DetectSuccess
	}
	result.Duration = duration
	if resp != nil {
		result.StatusCode = resp.StatusCode
	}
	if ht != nil {
		stats := ht.Stats()
		result.DNSLookup = int(stats.DNSLookup.Milliseconds())
		result.TCPConnection = int(stats.TCPConnection.Milliseconds())
		result.TLSHandshake = int(stats.TLSHandshake.Milliseconds())
		result.ServerProcessing = int(stats.ServerProcessing.Milliseconds())
		result.ContentTransfer = int(stats.ContentTransfer.Milliseconds())

		result.Addrs = ht.Addrs
		result.Addr = ht.Addr
		result.Protocol = ht.Protocol
		result.TLSVersion = ht.TLSVersion
		result.TLSCipherSuite = ht.TLSCipherSuite
		if len(ht.Certificates) != 0 {
			result.CertificateDNSNames = ht.Certificates[0].DNSNames
			result.CertificateExpirationDates = []string{
				ht.Certificates[0].NotBefore.String(),
				ht.Certificates[0].NotAfter.String(),
			}
		}
	}
	task := fmt.Sprintf("http-%d", result.Task)
	if isDetectResultChange(task, result.Result) {
		result.Receivers = h.Receivers
		go srv.alarm(result)
	}

	err = pgCreate(&result)
	if err != nil {
		logger.Error("http detect one fail",
			zap.Error(err),
		)
	}
}

func (srv *HTTPSrv) alarm(result HTTPDetectResult) {
	users, err := getReceivers(result.Receivers)
	if err != nil {
		logger.Error("get user list fail",
			zap.Any("ids", result.Receivers),
			zap.Error(err),
		)
	}
	duration := formatMs(result.Duration)
	status := "Success"
	if result.Result == DetectFail {
		status = "Fail"
	}
	data := Alarm{
		Title: fmt.Sprintf("%s: http detect %s", status, result.URL),
		Content: fmt.Sprintf(`url: %s
ip: %s
statusCode: %d
addrs: %s
addr: %s
protocol: %s
tlsVersion: %s
certificateDNSNames: %s
certificateExpirationDates: %s
dnsLookup: %s
tcpConnection: %s
tlsHandshake: %s
serverProcessing: %s
contentTransfer: %s
durtaion: %s 
message: %s`,
			result.URL,
			result.IP,
			result.StatusCode,
			result.Addrs,
			result.Addr,
			result.Protocol,
			result.TLSVersion,
			result.CertificateDNSNames,
			result.CertificateExpirationDates,
			formatMs(result.DNSLookup),
			formatMs(result.TCPConnection),
			formatMs(result.TLSHandshake),
			formatMs(result.ServerProcessing),
			formatMs(result.ContentTransfer),
			duration,
			result.Message,
		),
	}
	doAlarms(data, users)
}
