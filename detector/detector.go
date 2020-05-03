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
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/groupcache/lru"
	"github.com/lib/pq"
	"github.com/vicanso/cyber-tect/helper"
	"github.com/vicanso/cyber-tect/log"
	"github.com/vicanso/cyber-tect/service"
	"github.com/vicanso/hes"
	"go.uber.org/zap"
)

var (
	pgCreate    = helper.PGCreate
	pgGetClient = helper.PGGetClient

	logger = log.Default()

	latestDetectResultLRU = lru.New(10 * 1024)
	// 用户服务
	userSrv = new(service.UserSrv)

	errOwnerInvalid = &hes.Error{
		Message:    "You should owner the detect",
		StatusCode: http.StatusForbidden,
	}
)

const (
	defaultTimeout = 3 * time.Second
)

const (
	StatusUnknown = iota
	StatusEnabled
	StatusDisabled
)

const (
	DetectSuccess = iota + 1
	DetectFail
)

type Alarm struct {
	Title   string
	Content string
}

func isDetectResultChange(task string, currentStatus int) bool {
	value, ok := latestDetectResultLRU.Get(task)
	latestDetectResultLRU.Add(task, currentStatus)
	// 如果一开始状态不存在，而且当前状态是失败，则认为状态变化
	if !ok && currentStatus == DetectFail {
		return true
	}
	status, ok := value.(int)
	// 如果转换失败，则认为未变化
	if !ok {
		return false
	}
	return status != currentStatus
}

// portCheck the port check
func portCheck(network, ip string, port int, timeout time.Duration) (err error) {
	addr := ip
	if port != 0 {
		addr = net.JoinHostPort(ip, strconv.Itoa(port))
	}
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return
	}
	defer conn.Close()
	return
}

// doAlarms 发送告警，根据各用户选择的告警配置发送
func doAlarms(alarm Alarm, users []*service.User) {
	logger.Info("detect alarm",
		zap.String("title", alarm.Title),
		zap.String("content", alarm.Content),
	)
	emails := make([]string, 0)
	for _, user := range users {
		// 暂时仅支持email告警
		if user.Email != "" {
			emails = append(emails, user.Email)
		}
	}
	emailAlarm(alarm, emails)
}

// emailAlarm 邮件告警
func emailAlarm(alarm Alarm, emails []string) {
	if len(emails) == 0 {
		return
	}
	service.SendMail(alarm.Title, alarm.Content, emails)
}

// getReceivers get receivers by id
func getReceivers(receivers pq.Int64Array) (users []*service.User, err error) {
	ids := make([]uint, len(receivers))
	for index, id := range receivers {
		ids[index] = uint(id)
	}

	users, err = userSrv.List(service.UserQueryParams{
		IDList: ids,
	})
	return
}

// formatMs format ms to string
func formatMs(ms int) string {
	return (time.Duration(ms) * time.Millisecond).String()
}
