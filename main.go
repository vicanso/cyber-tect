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

/*
Package main Forest Server

	This should demonstrate all the possible comment annotations
	that are available to turn go code into a fully compliant swagger 2.0 spec

Host: localhost
BasePath: /
Version: 1.0.0
Schemes: http

Consumes:
- application/json

Produces:
- application/json

swagger:meta
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
	warner "github.com/vicanso/count-warner"
	"github.com/vicanso/cyber-tect/config"
	_ "github.com/vicanso/cyber-tect/controller"
	"github.com/vicanso/cyber-tect/cs"
	"github.com/vicanso/cyber-tect/helper"
	"github.com/vicanso/cyber-tect/log"
	"github.com/vicanso/cyber-tect/middleware"
	"github.com/vicanso/cyber-tect/router"
	_ "github.com/vicanso/cyber-tect/schedule"
	"github.com/vicanso/cyber-tect/service"
	"github.com/vicanso/cyber-tect/util"
	"github.com/vicanso/elton"
	M "github.com/vicanso/elton/middleware"
	"github.com/vicanso/hes"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

var (
	// Version version of tiny
	Version string
	// BuildAt build at
	BuildAt string
	// GO go version
	GO string
)

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(func(format string, args ...interface{}) {
		value := fmt.Sprintf(format, args...)
		log.Default().Info(value)
	}))
}

// 相关依赖服务的校验，主要是数据库等
func dependServiceCheck() (err error) {
	err = helper.RedisPing()
	if err != nil {
		return
	}
	configSrv := new(service.ConfigurationSrv)
	err = configSrv.Refresh()
	if err != nil {
		return
	}
	return
}

func main() {
	showVersion := flag.Bool("v", false, "show version")

	flag.Parse()
	if *showVersion {
		fmt.Printf("version %s\nbuild at %s\n%s\n", Version, BuildAt, GO)
		return
	}
	defer func() {
		// 关闭influxdb，flush统计数据
		helper.GetInfluxSrv().Close()
	}()

	logger := log.Default()
	e := elton.New()

	e.SignedKeys = service.GetSignedKeys()
	e.GenerateID = func() string {
		return util.RandomString(8)
	}

	// 未处理的error才会触发
	// 如果1分钟出现超过5次未处理异常
	// exception的warner只有一个key，因此无需定时清除
	warnerException := warner.NewWarner(60*time.Second, 5)
	warnerException.ResetOnWarn = true
	warnerException.On(func(_ string, _ warner.Count) {
		service.AlarmError("too many uncaught exception")
	})
	e.OnError(func(c *elton.Context, err error) {
		if !util.IsProduction() {
			he, ok := err.(*hes.Error)
			if ok {
				if he.Extra == nil {
					he.Extra = make(map[string]interface{})
				}
				he.Extra["stack"] = util.GetStack(5)
			}
		}

		// 可以针对实际场景输出更多的日志信息
		logger.DPanic("exception",
			zap.String("ip", c.RealIP()),
			zap.String("uri", c.Request.RequestURI),
			zap.Error(err),
		)
		warnerException.Inc("exception", 1)
	})
	// 对于404的请求，不会执行中间件，一般都是因为攻击之类才会导致大量出现404，
	// 因此可在此处汇总出错IP，针对较频繁出错IP，增加告警信息
	// 如果1分钟同一个IP出现60次404
	warner404 := warner.NewWarner(60*time.Second, 60)
	warner404.ResetOnWarn = true
	warner404.On(func(ip string, _ warner.Count) {
		service.AlarmError("too many 404 request, client ip:" + ip)
	})
	go func() {
		// 因为404是根据IP来告警，因此可能存在大量不同的key，因此定时清除过期数据
		for range time.NewTicker(5 * time.Minute).C {
			warner404.ClearExpired()
		}
	}()

	e.NotFoundHandler = func(resp http.ResponseWriter, req *http.Request) {
		ip := elton.GetClientIP(req)
		logger.Info("404",
			zap.String("ip", ip),
			zap.String("uri", req.RequestURI),
		)
		resp.Header().Set(elton.HeaderContentType, elton.MIMEApplicationJSON)
		resp.WriteHeader(http.StatusNotFound)
		_, err := resp.Write([]byte(`{"statusCode": 404,"message": "Not found"}`))
		if err != nil {
			logger.Info("404 response fail",
				zap.String("ip", ip),
				zap.String("uri", req.RequestURI),
				zap.Error(err),
			)
		}
		warner404.Inc(ip, 1)
	}

	// 捕捉panic异常，避免程序崩溃
	e.Use(M.NewRecover())

	e.Use(middleware.NewEntry())

	// 接口相关统计信息
	e.Use(M.NewStats(M.StatsConfig{
		OnStats: func(info *M.StatsInfo, c *elton.Context) {
			// ping 的日志忽略
			if info.URI == "/ping" {
				return
			}
			sid := util.GetSessionID(c)
			logger.Info("access log",
				zap.String("id", info.CID),
				zap.String("ip", info.IP),
				zap.String("sid", sid),
				zap.String("method", info.Method),
				zap.String("route", info.Route),
				zap.String("uri", info.URI),
				zap.Int("status", info.Status),
				zap.Uint32("connecting", info.Connecting),
				zap.String("consuming", info.Consuming.String()),
				zap.String("size", humanize.Bytes(uint64(info.Size))),
			)
			tags := map[string]string{
				"method": info.Method,
				"route":  info.Route,
			}
			fields := map[string]interface{}{
				"id":         info.CID,
				"ip":         info.IP,
				"sid":        sid,
				"uri":        info.URI,
				"status":     info.Status,
				"use":        info.Consuming.Milliseconds(),
				"size":       info.Size,
				"connecting": info.Connecting,
			}
			helper.GetInfluxSrv().Write(cs.MeasurementHTTP, fields, tags)
		},
	}))

	// 错误处理，将错误转换为json响应
	e.Use(M.NewError(M.ErrorConfig{
		ResponseType: "json",
	}))

	// IP限制
	e.Use(middleware.NewIPBlock())

	// 根据配置对路由mock返回
	e.Use(middleware.NewRouterMocker())

	// 路由并发限制
	routerLimitConfig := config.GetRouterConcurrentLimit()
	if len(routerLimitConfig) != 0 {
		e.Use(M.NewRCL(M.RCLConfig{
			Limiter: M.NewLocalLimiter(routerLimitConfig),
		}))
	}

	// etag与fresh的处理
	e.Use(M.NewDefaultFresh())
	e.Use(M.NewDefaultETag())

	// 对响应数据 c.Body 转换为相应的json响应
	e.Use(M.NewDefaultResponder())

	// 读取读取body的数的，转换为json bytes
	e.Use(M.NewDefaultBodyParser())

	// 初始化路由
	for _, g := range router.GetGroups() {
		e.AddGroup(g)
	}

	err := dependServiceCheck()
	if err != nil {
		service.AlarmError("check depend service fail, " + err.Error())
		// 可以针对实际场景输出更多的日志信息
		logger.DPanic("exception",
			zap.Error(err),
		)
		panic(err)
	}
	logger.Info("start to linstening...",
		zap.String("listen", config.GetListen()),
	)
	err = e.ListenAndServe(config.GetListen())
	if err != nil {
		service.AlarmError("listen and serve fail, " + err.Error())
		panic(err)
	}
}
