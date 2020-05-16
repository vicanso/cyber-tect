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
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/vicanso/cybertect/helper"
	"go.uber.org/zap"
)

type (
	Ping struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Name        string         `json:"name,omitempty"`
		Owner       string         `json:"owner,omitempty" gorm:"index:idx_ping_owner"`
		Status      int            `json:"status,omitempty" gorm:"index:idx_ping_status"`
		Description string         `json:"description,omitempty"`
		Receivers   pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`

		Network string `json:"network,omitempty"`
		IP      string `json:"ip,omitempty"`
		Timeout string `json:"timeout,omitempty"`
	}

	PingDetectResult struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Receivers pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`
		Duration  int            `json:"duration,omitempty"`
		Result    int            `json:"result,omitempty" gorm:"index:idx_ping_detect_result_result"`
		Message   string         `json:"message,omitempty"`

		Task    uint   `json:"task,omitempty" grom:"index:idx_ping_detect_result_task"`
		Network string `json:"network,omitempty"`
		IP      string `json:"ip,omitempty"`
	}

	PingSrv struct{}
)

func init() {
	pgGetClient().AutoMigrate(&Ping{}).
		AutoMigrate(&PingDetectResult{})
}

// Ping ping the ip
func (p *Ping) Ping() error {
	network := p.Network
	if network == "" {
		network = "ip4:icmp"
	}
	timeout, _ := time.ParseDuration(p.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	return portCheck(network, p.IP, 0, timeout)
}

// Add add ping detector
func (srv *PingSrv) Add(p *Ping) (err error) {
	err = pgCreate(p)
	return
}

// UpdateByID update the ping detector
func (srv *PingSrv) UpdateByID(id uint, attrs ...interface{}) (err error) {
	err = pgGetClient().Model(&Ping{
		ID: id,
	}).Update(attrs...).Error
	return
}

// FindByID find by the ping detector by id
func (srv *PingSrv) FindByID(id uint) (data *Ping, err error) {
	data = &Ping{}
	err = pgGetClient().Where("id = ?", id).First(data).Error
	if err != nil {
		return
	}
	return
}

// List list the ping detector
func (srv *PingSrv) List(params helper.PGQueryParams, args ...interface{}) (data []*Ping, err error) {
	data = make([]*Ping, 0)
	err = pgQuery(params, args...).Find(&data).Error
	return
}

// Count count the ping detector
func (srv *PingSrv) Count(args ...interface{}) (count int, err error) {
	return pgCount(&Ping{}, args...)
}

// ListResult list the ping detect result
func (srv *PingSrv) ListResult(params helper.PGQueryParams, args ...interface{}) (data []*PingDetectResult, err error) {
	data = make([]*PingDetectResult, 0)
	err = pgQuery(params, args...).Find(&data).Error
	return
}

// CountResult count the ping detect result
func (srv *PingSrv) CountResult(args ...interface{}) (count int, err error) {
	return pgCount(&PingDetectResult{}, args...)
}

// ValidateOwner validate the owner
func (srv *PingSrv) ValidateOwner(id uint, owner string) (err error) {
	data, err := srv.FindByID(id)
	if err != nil {
		return
	}
	if data.Owner != owner {
		err = errOwnerInvalid
	}
	return
}

// Detect do the ping detect
func (srv *PingSrv) Detect() {
	result := make([]*Ping, 0)
	err := pgGetClient().Where("status = ?", StatusEnabled).Find(&result).Error
	if err != nil {
		logger.Error("get ping detector fail",
			zap.Error(err),
		)
	}
	for _, ping := range result {
		go srv.detectOne(ping)
	}
}

// detectOne detect one ping
func (srv *PingSrv) detectOne(ping *Ping) {
	result := PingDetectResult{
		IP:      ping.IP,
		Network: ping.Network,
		Task:    ping.ID,
	}
	startedAt := time.Now()
	err := ping.Ping()
	duration := int(time.Since(startedAt).Milliseconds())
	if err != nil {
		result.Message = err.Error()
		result.Result = DetectFail
	} else {
		result.Result = DetectSuccess
	}
	// 设置最少时间为1
	if duration == 0 {
		duration = 1
	}
	result.Duration = duration
	task := fmt.Sprintf("ping-%d", result.Task)
	if isDetectResultChange(task, result.Result) {
		result.Receivers = ping.Receivers
		go srv.alarm(result)
	}
	err = pgCreate(&result)
	if err != nil {
		logger.Error("tcp detect one fail",
			zap.Error(err),
		)
	}
}

// alarm ping detect alarm
func (srv *PingSrv) alarm(result PingDetectResult) {
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
		Title: fmt.Sprintf("%s: ping detect %s", status, result.IP),
		Content: fmt.Sprintf(`network: %s
ip: %s
durtaion: %s 
message: %s`,
			result.Network,
			result.IP,
			duration,
			result.Message,
		),
	}
	doAlarms(data, users)
}
