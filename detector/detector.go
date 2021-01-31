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
	"net"
	"strings"
	"sync"
	"time"

	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/service"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

const defaultTimeout = 3 * time.Second

var getEntClient = helper.EntGetClient

// 记录失败次数
var taskFailCountMap = sync.Map{}

type alarmDetail struct {
	Name      string
	Receivers []string
	Task      string
	IsSuccess bool
	Messages  []string
}

// portCheck the port check
func portCheck(network, addr string, timeout time.Duration) (err error) {
	if network == "" {
		network = "tcp"
	}
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return
	}
	defer conn.Close()
	return
}

func doAlarm(detail alarmDetail) {
	value, _ := taskFailCountMap.LoadOrStore(detail.Task, atomic.NewUint32(0))
	failCount, ok := value.(*atomic.Uint32)
	if !ok {
		return
	}
	currentCount := failCount.Load()
	newCount := uint32(0)
	if !detail.IsSuccess {
		newCount = failCount.Inc()
	}
	// 状态未变化
	if currentCount == newCount {
		return
	}
	// 如果状态变化，而且此次是success
	title := ""
	message := ""
	if detail.IsSuccess {
		title = detail.Name + "(success)"
	} else if newCount == 1 || newCount%3 == 0 {
		// 如果是首次失败，或者每10次失败均发送系统失败信息
		message = strings.Join(detail.Messages, ",")
		title = detail.Name + "(fail)"
		if message == "" {
			message = "检测失败，未知异常"
		}
	}

	log.Default().Info("detect alarm",
		zap.String("name", title),
		zap.String("message", message),
	)
	service.SendEmail(title, message, detail.Receivers...)

}
