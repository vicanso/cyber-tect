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
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/email"
	"github.com/vicanso/cybertect/ent/dnsdetectorresult"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/ent/pingdetectorresult"
	"github.com/vicanso/cybertect/ent/tcpdetectorresult"
	"github.com/vicanso/cybertect/ent/user"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/go-axios"
	parallel "github.com/vicanso/go-parallel"
	"github.com/vicanso/hes"
	"go.uber.org/atomic"
)

const defaultTimeout = 3 * time.Second

var getEntClient = helper.EntGetClient

var detectorConfig = config.MustGetDetectorConfig()

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

func isMatchAlarmCount(count int) bool {
	// 前三次每次告警
	if count <= 3 {
		return true
	}
	// <=60时，每10次发送一次
	if count <= 60 && count%10 == 0 {
		return true
	}
	if count > 60 && count%60 == 0 {
		return true
	}
	return false
}

func isWeComRobot(url string) bool {
	return strings.HasPrefix(url, "https://qyapi.weixin.qq.com")
}

func doAlarm(ctx context.Context, detail alarmDetail) {
	value, _ := taskFailCountMap.LoadOrStore(detail.Task, atomic.NewUint32(0))
	failCount, ok := value.(*atomic.Uint32)
	if !ok {
		return
	}
	currentCount := failCount.Load()
	newCount := uint32(0)
	// 成功则失败次数重置为0
	if detail.IsSuccess {
		failCount.Store(0)
	} else {
		newCount = failCount.Inc()
	}
	// 状态未变化
	if currentCount == newCount {
		return
	}

	// 如果非前几次失败，则后续的告警不再每次发送
	// 由于使用邮件告警，容易忽略，因此前几次均发送
	if !isMatchAlarmCount(int(newCount)) {
		return
	}
	// 如果状态变化，而且此次是success
	title := ""
	message := ""
	if detail.IsSuccess {
		title = detail.Name + "(success)"
	} else {
		message = strings.Join(detail.Messages, "\n")
		title = detail.Name + "(fail)"
		if message == "" {
			message = "检测失败，未知异常"
		}
	}
	users, err := helper.EntGetClient().User.Query().
		Where(user.AccountIn(detail.Receivers...)).
		Select(user.FieldEmail, user.FieldAlarmURL).
		All(context.Background())
	if err != nil {
		log.Error(ctx).
			Err(err).
			Msg("get receiver email fail")
		return
	}
	log.Info(ctx).
		Str("name", title).
		Str("message", message).
		Msg("")

	emails := make([]string, 0)
	for _, item := range users {
		// 如果有配置了alarm地址
		if item.AlarmURL != "" {
			go func() {
				ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()
				conf := &axios.Config{
					Context: ctx,
					Method:  http.MethodPost,
					URL:     item.AlarmURL,
					Body: map[string]string{
						"title":   title,
						"message": message,
					},
				}
				// 如果是企业微信机器人
				if isWeComRobot(item.AlarmURL) {
					buf := []byte(message)
					maxSize := 4096
					// 直接截短
					if len(buf) > maxSize {
						message = string(buf[:maxSize])
					}

					t, _ := util.ChinaNow()

					content := fmt.Sprintf("时间：%s\n标题：%s\n内容：%s", t, title, message)
					conf.Body = map[string]interface{}{
						"msgtype": "markdown",
						"markdown": map[string]string{
							"content": content,
						},
					}
				}
				_, err := axios.Request(conf)
				sendAlarmCategory := "sendAlarm"
				if err != nil {
					log.Error(ctx1).
						Str("category", sendAlarmCategory).
						Err(err)
				} else {
					log.Info(ctx1).
						Str("category", sendAlarmCategory).
						Send()
				}
			}()
		} else if item.Email != "" {
			emails = append(emails, item.Email)
		}
	}
	if len(emails) == 0 {
		return
	}

	email.Send(ctx, title, message, emails...)
}

func convertParallelError(err error, message string) error {
	if err == nil {
		return nil
	}
	errs := hes.Error{
		Message: message,
	}
	pErr, _ := err.(*parallel.Errors)
	if pErr != nil {
		for _, e := range pErr.Errs {
			errs.Add(e)
		}
	}
	return &errs
}

func shouldDoDetect(intervalValue string, count int64) bool {
	interval, _ := time.ParseDuration(intervalValue)

	configInterval, _ := time.ParseDuration(detectorConfig.Interval)
	// 如果间隔<=1s，则每次都检测
	// 或者检测间隔少于应用配置检测检测
	if interval <= time.Second ||
		configInterval <= time.Second ||
		interval < configInterval {
		return true
	}
	value := int64(interval / configInterval)
	fmt.Println(count % value)
	return count%value == 0
}

// RemoveExpiredDetectorResult 清除过期数据
func RemoveExpiredDetectorResult() (err error) {
	// 如果不需要清除则直接返回
	if detectorConfig.Expired <= 0 {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	t := time.Now().Add(-detectorConfig.Expired)
	_, err = getEntClient().DNSDetectorResult.
		Delete().
		Where(dnsdetectorresult.UpdatedAtLT(t)).
		Exec(ctx)
	if err != nil {
		return
	}
	_, err = getEntClient().HTTPDetectorResult.
		Delete().
		Where(httpdetectorresult.UpdatedAtLT(t)).
		Exec(ctx)
	if err != nil {
		return
	}

	_, err = getEntClient().TCPDetectorResult.
		Delete().
		Where(tcpdetectorresult.UpdatedAtLT(t)).
		Exec(ctx)
	if err != nil {
		return
	}
	_, err = getEntClient().PingDetectorResult.
		Delete().
		Where(pingdetectorresult.UpdatedAtLT(t)).
		Exec(ctx)
	if err != nil {
		return
	}
	return
}
func ceilToMs(d time.Duration) int {
	if d == 0 {
		return 0
	}
	offset := 0
	if d%time.Millisecond != 0 {
		offset++
	}
	return int(d.Milliseconds()) + offset
}
