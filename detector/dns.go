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
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/vicanso/cybertect/helper"
	"go.uber.org/zap"
)

type (
	DNS struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Name        string         `json:"name,omitempty"`
		Owner       string         `json:"owner,omitempty" gorm:"index:idx_dns_owner"`
		Status      int            `json:"status,omitempty" gorm:"index:idx_dns_status"`
		Description string         `json:"description,omitempty"`
		Receivers   pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`

		// Server dns server(ip:port)
		Server   string `json:"server,omitempty" gorm:"type:varchar(64);not null;"`
		Hostname string `json:"hostname,omitempty" gorm:"type:varchar(64);not null;"`
		Timeout  string `json:"timeout,omitempty"`
	}
	DNSDetectResult struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Receivers pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`
		Duration  int            `json:"duration,omitempty"`
		Result    int            `json:"result,omitempty" gorm:"index:idx_dns_detect_result_result"`
		Message   string         `json:"message,omitempty"`

		Task     uint           `json:"task,omitempty" gorm:"index:idx_dns_detect_result_task"`
		Server   string         `json:"server,omitempty" gorm:"type:varchar(64)"`
		Hostname string         `json:"hostname,omitempty" gorm:"type:varchar(64)"`
		IPAddrs  pq.StringArray `json:"ipAddrs,omitempty" gorm:"type:text[]"`
	}

	DNSSrv struct{}
)

func init() {
	pgGetClient().AutoMigrate(&DNS{}).
		AutoMigrate(&DNSDetectResult{})
}

// Resolve resolve the hostname
func (d *DNS) Resolve() ([]net.IPAddr, error) {
	dnsServer := d.Server
	if !strings.Contains(dnsServer, ":") {
		dnsServer += ":53"
	}
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dia := net.Dialer{}
			return dia.DialContext(ctx, "udp", dnsServer)
		},
	}
	timeout, _ := time.ParseDuration(d.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return r.LookupIPAddr(ctx, d.Hostname)
}

// Add add dns detector
func (srv *DNSSrv) Add(d *DNS) (err error) {
	err = pgCreate(d)
	return
}

// UpdateByID update the dns detector
func (srv *DNSSrv) UpdateByID(id uint, attrs ...interface{}) (err error) {
	err = pgGetClient().Model(&DNS{
		ID: id,
	}).Update(attrs...).Error
	return
}

// FindByID find the dns detector by id
func (srv *DNSSrv) FindByID(id uint) (data *DNS, err error) {
	data = &DNS{}
	err = pgGetClient().Where("id = ?", id).First(data).Error
	if err != nil {
		return
	}
	return
}

// List list the dns detector
func (srv *DNSSrv) List(params helper.PGQueryParams, args ...interface{}) (data []*DNS, err error) {
	data = make([]*DNS, 0)
	err = pgQuery(params).Find(&data).Error
	return
}

// Count count the dns detector
func (srv *DNSSrv) Count(args ...interface{}) (count int, err error) {
	return pgCount(&DNS{}, args...)
}

// ListResult list the dns detect result
func (srv *DNSSrv) ListResult(params helper.PGQueryParams, args ...interface{}) (data []*DNSDetectResult, err error) {
	data = make([]*DNSDetectResult, 0)
	err = pgQuery(params, args...).Find(&data).Error
	return
}

// CountResult count the dns detect result
func (srv *DNSSrv) CountResult(args ...interface{}) (count int, err error) {
	return pgCount(&DNSDetectResult{}, args...)
}

// Detect do the dns detect
func (srv *DNSSrv) Detect() {
	result := make([]*DNS, 0)
	err := pgGetClient().Where("status = ?", StatusEnabled).Find(&result).Error
	if err != nil {
		logger.Error("get dns detector fail",
			zap.Error(err),
		)
	}
	for _, dns := range result {
		go srv.detectOne(dns)
	}
}

// ValidateOwner validate the owner
func (srv *DNSSrv) ValidateOwner(id uint, owner string) (err error) {
	data, err := srv.FindByID(id)
	if err != nil {
		return
	}
	if data.Owner != owner {
		err = errOwnerInvalid
	}
	return
}

// detectOne detect one
func (srv *DNSSrv) detectOne(dns *DNS) {
	result := DNSDetectResult{
		Server:   dns.Server,
		Hostname: dns.Hostname,
		Task:     dns.ID,
	}
	startedAt := time.Now()
	ipAddrs, err := dns.Resolve()
	duration := int(time.Since(startedAt).Milliseconds())
	if err != nil {
		result.Message = err.Error()
		result.Result = DetectFail
	} else {
		result.Result = DetectSuccess
		result.IPAddrs = make([]string, len(ipAddrs))
		for index, item := range ipAddrs {
			result.IPAddrs[index] = item.String()
		}
	}
	result.Duration = duration
	task := fmt.Sprintf("dns-%d", result.Task)
	if isDetectResultChange(task, result.Result) {
		result.Receivers = dns.Receivers
		go srv.alarm(result)
	}
	err = pgCreate(&result)
	if err != nil {
		logger.Error("dns detect one fail",
			zap.Error(err),
		)
	}
}

// alarm dns detect alarm
func (srv *DNSSrv) alarm(result DNSDetectResult) {
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
		Title: fmt.Sprintf("%s: dns detect %s", status, result.Hostname),
		Content: fmt.Sprintf(`hostname: %s
server: %s
ips: %s
durtaion: %s 
message: %s`,
			result.Hostname,
			result.Server,
			result.IPAddrs,
			duration,
			result.Message,
		),
	}
	doAlarms(data, users)
}
