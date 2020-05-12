// Copyright 2019 tree xie
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

package schedule

import (
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/detector"

	"github.com/robfig/cron/v3"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/service"

	"go.uber.org/zap"
)

func init() {
	c := cron.New()
	_, _ = c.AddFunc("@every 5m", redisCheck)
	_, _ = c.AddFunc("@every 1m", configRefresh)
	_, _ = c.AddFunc("@every 5m", redisStats)
	_, _ = c.AddFunc("@every 1m", pgStats)
	// DNS检测5分钟一次则可
	_, _ = c.AddFunc("@every 5m", dnsCheck)
	_, _ = c.AddFunc("@every 1m", tcpCheck)
	_, _ = c.AddFunc("@every 1m", pingCheck)
	_, _ = c.AddFunc("@every 1m", httpCheck)
	c.Start()
}

func redisCheck() {
	err := helper.RedisPing()
	if err != nil {
		log.Default().Error("redis check fail",
			zap.Error(err),
		)
		service.AlarmError("redis check fail")
	}
}

func configRefresh() {
	configSrv := new(service.ConfigurationSrv)
	err := configSrv.Refresh()
	if err != nil {
		log.Default().Error("config refresh fail",
			zap.Error(err),
		)
		service.AlarmError("config refresh fail")
	}
}

func redisStats() {
	stats := helper.RedisStats()
	helper.GetInfluxSrv().Write(cs.MeasurementRedisStats, stats, nil)
}

func pgStats() {
	stats := helper.PGStats()
	helper.GetInfluxSrv().Write(cs.MeasurementPGStats, stats, nil)
}

func dnsCheck() {
	(&detector.DNSSrv{}).Detect()
}

func tcpCheck() {
	(&detector.TCPSrv{}).Detect()
}

func pingCheck() {
	(&detector.PingSrv{}).Detect()
}

func httpCheck() {
	(&detector.HTTPSrv{}).Detect()
}
