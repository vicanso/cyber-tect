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
	TCP struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Owner       string         `json:"owner,omitempty" gorm:"index:idx_tcp_owner"`
		Status      int            `json:"status,omitempty" gorm:"index:idx_tcp_status"`
		Description string         `json:"description,omitempty"`
		Receivers   pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`

		Network string `json:"network,omitempty"`
		IP      string `json:"ip,omitempty"`
		Port    int    `json:"port,omitempty"`
		Timeout string `json:"timeout,omitempty"`
	}

	TCPDetectResult struct {
		ID        uint       `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

		Receivers pq.StringArray `json:"receivers,omitempty" gorm:"type:text[]"`
		Duration  int            `json:"duration,omitempty"`
		Result    int            `json:"result,omitempty"`
		Message   string         `json:"message,omitempty"`

		Task    uint   `json:"task,omitempty" gorm:"index:idx_tcp_detect_result_task"`
		Network string `json:"network,omitempty"`
		IP      string `json:"ip,omitempty"`
		Port    int    `json:"port,omitempty"`
	}

	TCPSrv struct{}
)

func init() {
	pgGetClient().AutoMigrate(&TCP{}).
		AutoMigrate(&TCPDetectResult{})
}

// Check check the ip:port
func (t *TCP) Check() error {
	network := t.Network
	if network == "" {
		network = "tcp"
	}
	timeout, _ := time.ParseDuration(t.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	return portCheck(network, t.IP, t.Port, timeout)
}

// Add add tcp detector
func (srv *TCPSrv) Add(t *TCP) (err error) {
	err = pgCreate(t)
	return
}

// UpdateByID update the tcp detector
func (srv *TCPSrv) UpdateByID(id uint, attrs ...interface{}) (err error) {
	err = pgGetClient().Model(&TCP{
		ID: id,
	}).Update(attrs...).Error
	return
}

// FindByID find by the tcp detector by id
func (srv *TCPSrv) FindByID(id uint) (data *TCP, err error) {
	data = &TCP{}
	err = pgGetClient().Where("id = ?", id).First(data).Error
	if err != nil {
		return
	}
	return
}

// List list the tcp detector
func (srv *TCPSrv) List(params helper.PGQueryParams, args ...interface{}) (data []*TCP, err error) {
	data = make([]*TCP, 0)
	err = pgQuery(params).Find(&data).Error
	return
}

// Count count the http detector
func (srv *TCPSrv) Count(args ...interface{}) (count int, err error) {
	return pgCount(&TCP{}, args...)
}

// ListResult list the tcp detect result
func (srv *TCPSrv) ListResult(params helper.PGQueryParams, args ...interface{}) (data []*TCPDetectResult, err error) {
	data = make([]*TCPDetectResult, 0)
	err = pgQuery(params, args...).Find(&data).Error
	return
}

// CountResult count the tcp detect result
func (srv *TCPSrv) CountResult(args ...interface{}) (count int, err error) {
	return pgCount(&TCPDetectResult{}, args...)
}

// Detect do the tcp detect
func (srv *TCPSrv) Detect() {
	result := make([]*TCP, 0)
	err := pgGetClient().Where("status = ?", StatusEnabled).Find(&result).Error
	if err != nil {
		logger.Error("get tcp detector fail",
			zap.Error(err),
		)
	}
	for _, tcp := range result {
		go srv.detectOne(tcp)
	}
}

// ValidateOwner validate the owner
func (srv *TCPSrv) ValidateOwner(id uint, owner string) (err error) {
	data, err := srv.FindByID(id)
	if err != nil {
		return
	}
	if data.Owner != owner {
		err = errOwnerInvalid
	}
	return
}

// detectOne detect one tcp
func (srv *TCPSrv) detectOne(tcp *TCP) {
	result := TCPDetectResult{
		Network: tcp.Network,
		IP:      tcp.IP,
		Port:    tcp.Port,
		Task:    tcp.ID,
	}
	startedAt := time.Now()
	err := tcp.Check()
	duration := int(time.Since(startedAt).Milliseconds())
	if err != nil {
		result.Message = err.Error()
		result.Result = DetectFail
	} else {
		result.Result = DetectSuccess
	}
	result.Duration = duration
	task := fmt.Sprintf("dns-%d", result.Task)
	if isDetectResultChange(task, result.Result) {
		result.Receivers = tcp.Receivers
		go srv.alarm(result)
	}
	err = pgCreate(&result)
	if err != nil {
		logger.Error("tcp detect one fail",
			zap.Error(err),
		)
	}
}

// alarm tcp detect alarm
func (srv *TCPSrv) alarm(result TCPDetectResult) {
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
		Title: fmt.Sprintf("%s: tcp detect %s:%d", status, result.IP, result.Port),
		Content: fmt.Sprintf(`network: %s
ip: %s
port: %d
durtaion: %s 
message: %s`,
			result.Network,
			result.IP,
			result.Port,
			duration,
			result.Message,
		),
	}
	doAlarms(data, users)
}
