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

package controller

import (
	"bytes"
	"strconv"

	"github.com/vicanso/cyber-tect/service"
	"github.com/vicanso/cyber-tect/util"

	"github.com/vicanso/elton"
	"github.com/vicanso/cyber-tect/router"
)

type (
	commonCtrl struct{}
)

func init() {
	ctrl := commonCtrl{}
	g := router.NewGroup("")

	g.GET("/ping", ctrl.ping)

	g.GET("/commons/ip-location", ctrl.location)

	g.GET("/commons/routers", ctrl.routers)

	g.GET("/commons/random-keys", ctrl.randomKeys)

	g.GET("/commons/captcha", ctrl.captcha)

	g.GET("/commons/performance", ctrl.getPerformance)
}

// 服务检测ping的响应
// swagger:response pingResponse
// nolint
type pongResponse struct {

	// in: body
	Payload string
}

// swagger:route GET /ping common ping
// ping
//
// 服务正常启动后则返回`pong`，主要用于反向代理的health check
// Responses:
// 	200: pingResponse
// Produces:
// 	- plain/text
func (ctrl commonCtrl) ping(c *elton.Context) error {
	c.BodyBuffer = bytes.NewBufferString("pong")
	return nil
}

// IP定位信息
// swagger:response locationResponse
// nolint
type locationResponse struct {

	// in: body
	Payload *service.Location
}

// swagger:route GET /commons/ip-location common commonsIPLocation
// ip2Location
//
// 从客户的真实IP地址获取定位信息
// Responses:
// 	200: locationResponse
func (ctrl commonCtrl) location(c *elton.Context) (err error) {
	info, err := service.GetLocationByIP(c.RealIP(), c)
	if err != nil {
		return
	}
	c.Body = info
	return
}

func (ctrl commonCtrl) routers(c *elton.Context) (err error) {
	c.Body = map[string]interface{}{
		"routers": c.Elton().Routers,
	}
	return
}

func (ctrl commonCtrl) randomKeys(c *elton.Context) (err error) {
	n, _ := strconv.Atoi(c.QueryParam("n"))
	size, _ := strconv.Atoi(c.QueryParam("size"))
	if size < 1 {
		size = 1
	}
	if n < 1 {
		n = 1
	}
	result := make([]string, size)
	for index := 0; index < size; index++ {
		result[index] = util.RandomString(n)
	}
	c.Body = map[string][]string{
		"keys": result,
	}
	return
}

func (ctrl commonCtrl) captcha(c *elton.Context) (err error) {
	bgColor := c.QueryParam("bg")
	fontColor := c.QueryParam("color")
	if bgColor == "" {
		bgColor = "255,255,255"
	}
	if fontColor == "" {
		fontColor = "102,102,102"
	}
	info, err := service.GetCaptcha(fontColor, bgColor)
	if err != nil {
		return
	}
	// c.SetContentTypeByExt(".jpeg")
	// c.Body = info.Data
	c.NoStore()
	c.Body = info
	return
}

func (ctrl commonCtrl) getPerformance(c *elton.Context) (err error) {
	c.Body = service.GetPerformance()
	return
}
