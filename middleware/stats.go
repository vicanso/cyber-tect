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

package middleware

import (
	"github.com/dustin/go-humanize"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
	M "github.com/vicanso/elton/middleware"
	"go.uber.org/zap"
)

func NewStats() elton.Handler {
	return M.NewStats(M.StatsConfig{
		OnStats: func(info *M.StatsInfo, c *elton.Context) {
			// ping 的日志忽略
			if info.URI == "/ping" {
				return
			}
			// TODO 如果需要可以从session中获取账号
			// 此处记录的session id，需要从客户登录记录中获取对应的session id
			// us := service.NewUserSession(c)
			sid := util.GetSessionID(c)
			log.Default().Info("access log",
				// 由客户端设置的uuid
				zap.String("uuid", c.GetRequestHeader("X-UUID")),
				zap.String("ip", info.IP),
				zap.String("sid", sid),
				zap.String("method", info.Method),
				zap.String("route", info.Route),
				zap.String("uri", info.URI),
				zap.Int("status", info.Status),
				zap.Uint32("connecting", info.Connecting),
				zap.String("consuming", info.Consuming.String()),
				zap.String("size", humanize.Bytes(uint64(info.Size))),
				zap.Int("bytes", info.Size),
			)
			tags := map[string]string{
				cs.TagMethod: info.Method,
				cs.TagRoute:  info.Route,
			}
			fields := map[string]interface{}{
				cs.FieldIP:         info.IP,
				cs.FieldSID:        sid,
				cs.FieldURI:        info.URI,
				cs.FieldStatus:     info.Status,
				cs.FieldUse:        int(info.Consuming.Milliseconds()),
				cs.FieldSize:       info.Size,
				cs.FieldProcessing: info.Connecting,
			}
			helper.GetInfluxSrv().Write(cs.MeasurementHTTPStats, tags, fields)
		},
	})
}
