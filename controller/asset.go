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

// 主要是管理系统的前端代码，对于资源等（如图片）尽可能不要打包进入程序

package controller

import (
	"bytes"
	"time"

	"github.com/vicanso/cybertect/asset"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/elton"
	M "github.com/vicanso/elton/middleware"
)

type (
	// assetCtrl asset ctrl
	assetCtrl struct{}
)

var assetFS = M.NewEmbedStaticFS(asset.GetFS(), "dist")

func init() {
	g := router.NewGroup("")
	ctrl := assetCtrl{}
	g.GET("/", ctrl.getIndex)
	g.GET("/favicon.{ext}", ctrl.getFavIcon)

	g.GET("/static/*", M.NewStaticServe(assetFS, M.StaticServeConfig{
		// 客户端缓存一年
		MaxAge: 365 * 24 * time.Hour,
		// 缓存服务器缓存一个小时
		SMaxAge:             time.Hour,
		DenyQueryString:     true,
		DisableLastModified: true,
		EnableStrongETag:    true,
		// 如果静态文件都有版本号，可以指定immutable
		Immutable: true,
	}))
}

// 静态文件响应
func sendFile(c *elton.Context, file string) (err error) {
	// 因为静态文件打包至程序中，直接读取
	buf, err := assetFS.Get(file)
	if err != nil {
		return
	}
	// 根据文件后续设置类型
	c.SetContentTypeByExt(file)
	c.BodyBuffer = bytes.NewBuffer(buf)
	return
}

// getIndex 首页
func (*assetCtrl) getIndex(c *elton.Context) (err error) {
	err = sendFile(c, "index.html")
	if err != nil {
		return
	}
	c.CacheMaxAge(10 * time.Second)
	return
}

// getFavIcon 图标
func (*assetCtrl) getFavIcon(c *elton.Context) (err error) {
	err = sendFile(c, "favicon.png")
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Hour, 10*time.Minute)
	return
}
