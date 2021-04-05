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

// 公共的处理函数，包括程序基本信息、性能指标等

package controller

import (
	"bytes"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/service"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
)

type (
	commonCtrl struct{}

	// applicationInfoResp 应用信息响应
	applicationInfoResp struct {
		// 版本号
		Version string `json:"version,omitempty"`
		// 构建时间
		BuildedAt string `json:"buildedAt,omitempty"`
		// 运行时长
		Uptime string `json:"uptime,omitempty"`
		// os类型
		OS string `json:"os,omitempty"`
		// go版本
		GO string `json:"go,omitempty"`
		// 架构类型
		ARCH string `json:"arch,omitempty"`
		// 运行环境配置
		ENV string `json:"env,omitempty"`
	}
	// routersResp 路由列表响应
	routersResp struct {
		// 路由信息
		Routers []elton.RouterInfo `json:"routers,omitempty"`
	}
	// statusListResp 状态列表响应
	statusListResp struct {
		Statuses []*schema.StatusInfo `json:"statuses,omitempty"`
	}
	// randomKeysResp 随机字符
	randomKeysResp struct {
		Keys []string `json:"keys,omitempty"`
	}
)

const (
	errCommonCategory = "common"
)

var (
	// applicationStartedAt 应用启动时间
	applicationStartedAt = time.Now()
)

func init() {
	ctrl := commonCtrl{}
	router.NewGroup("").GET("/ping", ctrl.ping)
	g := router.NewGroup("/commons")

	g.GET("/application", ctrl.getApplicationInfo)
	g.GET("/routers", ctrl.getRouters)
	g.GET("/captcha", ctrl.getCaptcha)
	g.GET("/performance", ctrl.getPerformance)
	g.GET("/schema-statuses", ctrl.listStatus)
	g.GET("/random-keys", ctrl.getRandomKeys)
	// 获取系统prof指标
	g.GET(
		"/prof",
		loadUserSession,
		shouldBeAdmin,
		ctrl.getProf,
	)
}

// ping 用于检测服务是否可用
func (*commonCtrl) ping(c *elton.Context) error {
	if !service.ApplicationIsRunning() {
		return hes.NewWithStatusCode("应用服务不可用", http.StatusServiceUnavailable, errCommonCategory)
	}
	c.BodyBuffer = bytes.NewBufferString("pong")
	return nil
}

// getApplicationInfo 获取应用信息
func (*commonCtrl) getApplicationInfo(c *elton.Context) (err error) {
	c.CacheMaxAge(time.Minute)
	c.Body = &applicationInfoResp{
		Version:   service.GetApplicationVersion(),
		BuildedAt: service.GetApplicationBuildedAt(),
		Uptime:    humanize.Time(applicationStartedAt),
		OS:        runtime.GOOS,
		GO:        runtime.Version(),
		ARCH:      runtime.GOARCH,
		ENV:       config.GetENV(),
	}
	return
}

// getRouters 获取系统的路由
func (*commonCtrl) getRouters(c *elton.Context) (err error) {
	c.CacheMaxAge(time.Minute)
	c.Body = &routersResp{
		Routers: c.Elton().GetRouters(),
	}
	return
}

// getCaptcha 获取图形验证码
func (*commonCtrl) getCaptcha(c *elton.Context) (err error) {
	bgColor := c.QueryParam("bg")
	fontColor := c.QueryParam("color")
	if bgColor == "" {
		bgColor = "255,255,255"
	}
	if fontColor == "" {
		fontColor = "102,102,102"
	}
	info, err := service.GetCaptcha(c.Context(), fontColor, bgColor)
	if err != nil {
		return
	}
	// 防止此字段未设置好，序列化至前端
	info.Value = ""
	c.NoStore()
	c.Body = &info
	return
}

// getPerformance 获取应用性能指标
func (*commonCtrl) getPerformance(c *elton.Context) (err error) {
	p := service.GetPerformance()
	c.Body = &p
	return
}

// listStatus 获取状态列表
func (*commonCtrl) listStatus(c *elton.Context) (err error) {
	c.CacheMaxAge(5 * time.Minute)
	c.Body = &statusListResp{
		Statuses: schema.GetStatusList(),
	}
	return
}

// getRandomKeys 获取随机字符串
func (*commonCtrl) getRandomKeys(c *elton.Context) (err error) {
	n, _ := strconv.Atoi(c.QueryParam("n"))
	size, _ := strconv.Atoi(c.QueryParam("size"))
	if size < 1 {
		size = 10
	}
	if n < 1 {
		n = 1
	}
	result := make([]string, n)
	for index := 0; index < n; index++ {
		result[index] = util.RandomString(size)
	}
	c.Body = &randomKeysResp{
		Keys: result,
	}
	return
}

// getProf 获取prof信息
func (*commonCtrl) getProf(c *elton.Context) (err error) {
	d := 30 * time.Second
	v := c.QueryParam("d")
	if v != "" {
		d, err = time.ParseDuration(v)
		if err != nil {
			return
		}
	}
	result, err := profSrv.Get(d)
	if err != nil {
		return
	}
	c.SetHeader(elton.HeaderContentType, elton.MIMEBinary)
	c.SetHeader("Content-Disposition", `attachment; filename="gprof"`)
	c.BodyBuffer = result
	return
}
