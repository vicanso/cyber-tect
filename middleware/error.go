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
	"bytes"
	"net/http"

	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/service"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
	"go.uber.org/zap"
)

// New Error handler
func NewError() elton.Handler {
	return func(c *elton.Context) error {
		err := c.Next()
		if err == nil {
			return nil
		}
		uri := c.Request.RequestURI
		he, ok := err.(*hes.Error)
		if !ok {
			// 如果不是以http error的形式返回的error则为非主动抛出错误
			log.Default().Error("unexpected error",
				zap.String("method", c.Request.Method),
				zap.String("route", c.Route),
				zap.String("uri", uri),
				zap.Error(err),
			)
			he = hes.NewWithError(err)
			he.StatusCode = http.StatusInternalServerError
			he.Exception = true
		} else {
			// 避免修改了原有的error对象
			he = he.Clone()
		}
		if he.StatusCode == 0 {
			he.StatusCode = http.StatusInternalServerError
		}
		if he.Extra == nil {
			he.Extra = make(map[string]interface{})
		}
		account := ""
		tid := util.GetTrackID(c)
		us := service.NewUserSession(c)
		if us != nil && us.IsLogin() {
			account = us.MustGetInfo().Account
		}
		ip := c.RealIP()
		sid := util.GetSessionID(c)

		he.Extra["route"] = c.Route
		// 记录用户相关信息
		fields := map[string]interface{}{
			cs.FieldStatus:    he.StatusCode,
			cs.FieldError:     he.Error(),
			cs.FieldURI:       uri,
			cs.FieldException: he.Exception,
			cs.FieldIP:        ip,
			cs.FieldSID:       sid,
			cs.FieldTID:       tid,
		}
		if account != "" {
			fields[cs.FieldAccount] = account
		}
		tags := map[string]string{
			cs.TagMethod: c.Request.Method,
			cs.TagRoute:  c.Route,
		}
		if he.Category != "" {
			tags[cs.TagCategory] = he.Category
		}

		helper.GetInfluxSrv().Write(cs.MeasurementHTTPError, tags, fields)
		c.StatusCode = he.StatusCode
		c.BodyBuffer = bytes.NewBuffer(he.ToJSON())
		return nil
	}
}
