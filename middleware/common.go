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
	"fmt"
	"net/http"
	"strings"
	"time"

	warner "github.com/vicanso/count-warner"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/service"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
	"go.uber.org/zap"
)

const (
	xCaptchaHeader    = "X-Captcha"
	errCommonCategory = "common-validate"
)

// WaitFor 延时响应中间件，设置最少等待多久再响应
func WaitFor(d time.Duration, onlyErrOccurreds ...bool) elton.Handler {
	ns := d.Nanoseconds()
	onlyErrOccurred := false
	if len(onlyErrOccurreds) != 0 {
		onlyErrOccurred = onlyErrOccurreds[0]
	}
	return func(c *elton.Context) (err error) {
		start := time.Now()
		err = c.Next()
		// 如果未出错，而且配置为仅在出错时才等待
		if err == nil && onlyErrOccurred {
			return
		}
		use := time.Now().UnixNano() - start.UnixNano()
		// 无论成功还是失败都wait for
		if use < ns {
			time.Sleep(time.Duration(ns-use) * time.Nanosecond)
		}
		return
	}
}

// ValidateCaptcha 图形难码校验
func ValidateCaptcha(magicalCaptcha string) elton.Handler {
	return func(c *elton.Context) (err error) {
		value := c.GetRequestHeader(xCaptchaHeader)
		if value == "" {
			err = hes.New("图形验证码参数不能为空", errCommonCategory)
			return
		}
		arr := strings.Split(value, ":")
		if len(arr) != 2 {
			err = hes.New(fmt.Sprintf("图形验证码参数长度异常(%d)", len(arr)), errCommonCategory)
			return
		}
		// 如果有配置万能验证码，则判断是否相等
		if magicalCaptcha != "" && arr[1] == magicalCaptcha {
			return c.Next()
		}
		valid, err := service.ValidateCaptcha(c.Context(), arr[0], arr[1])
		if err != nil {
			return err
		}
		if !valid {
			err = hes.New("图形验证码错误", errCommonCategory)
			return
		}
		return c.Next()
	}
}

// NewNoCacheWithCondition 创建no cache的中间件，此中间件根据设置的key value来判断是否设置为no cache
func NewNoCacheWithCondition(key, value string) elton.Handler {
	return func(c *elton.Context) (err error) {
		err = c.Next()
		if c.QueryParam(key) == value {
			c.NoCache()
		}
		return
	}
}

// NewNotFoundHandler 创建404 not found的处理函数
func NewNotFoundHandler() http.HandlerFunc {
	// 对于404的请求，不会执行中间件，一般都是因为攻击之类才会导致大量出现404，
	// 因此可在此处汇总出错IP，针对较频繁出错IP，增加告警信息
	// 如果1分钟同一个IP出现60次404
	warner404 := warner.NewWarner(60*time.Second, 60)
	warner404.ResetOnWarn = true
	warner404.On(func(ip string, _ int) {
		service.AlarmError("too many 404 request, client ip:" + ip)
	})
	go func() {
		// 因为404是根据IP来告警，因此可能存在大量不同的key，因此定时清除过期数据
		for range time.NewTicker(5 * time.Minute).C {
			warner404.ClearExpired()
		}
	}()
	notFoundErrBytes := (&hes.Error{
		Message:    "Not Found",
		StatusCode: http.StatusNotFound,
		Category:   "defaultNotFound",
	}).ToJSON()
	return func(resp http.ResponseWriter, req *http.Request) {
		ip := elton.GetClientIP(req)
		log.Default().Info("404",
			zap.String("ip", ip),
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
		)
		status := http.StatusNotFound
		resp.Header().Set(elton.HeaderContentType, elton.MIMEApplicationJSON)
		resp.WriteHeader(status)
		_, err := resp.Write(notFoundErrBytes)
		if err != nil {
			log.Default().Info("404 response fail",
				zap.String("ip", ip),
				zap.String("uri", req.RequestURI),
				zap.Error(err),
			)
		}
		warner404.Inc(ip, 1)

		tags := map[string]string{
			cs.TagMethod: req.Method,
		}
		fields := map[string]interface{}{
			cs.FieldIP:     ip,
			cs.FieldURI:    req.RequestURI,
			cs.FieldStatus: status,
		}
		helper.GetInfluxSrv().Write(cs.MeasurementHTTPStats, tags, fields)
	}
}

// NewMethodNotAllowedHandler 创建method not allowed的处理函数
func NewMethodNotAllowedHandler() http.HandlerFunc {
	methodNotAllowedErrBytes := (&hes.Error{
		Message:    "Method Not Allowed",
		StatusCode: http.StatusMethodNotAllowed,
		Category:   "defaultMethodNotAllowed",
	}).ToJSON()
	return func(resp http.ResponseWriter, req *http.Request) {
		ip := elton.GetClientIP(req)
		log.Default().Info("method not allowed",
			zap.String("ip", ip),
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
		)
		resp.Header().Set(elton.HeaderContentType, elton.MIMEApplicationJSON)
		status := http.StatusMethodNotAllowed
		resp.WriteHeader(status)
		_, err := resp.Write(methodNotAllowedErrBytes)
		if err != nil {
			log.Default().Info("method not allowed response fail",
				zap.String("ip", ip),
				zap.String("uri", req.RequestURI),
				zap.Error(err),
			)
		}
		tags := map[string]string{
			cs.TagMethod: req.Method,
		}
		fields := map[string]interface{}{
			cs.FieldIP:     ip,
			cs.FieldURI:    req.RequestURI,
			cs.FieldStatus: status,
		}
		helper.GetInfluxSrv().Write(cs.MeasurementHTTPStats, tags, fields)
	}
}
