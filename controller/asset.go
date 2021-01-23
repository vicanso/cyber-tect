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

// 通过packr2将静态文件打包，此controller提供各静态文件的响应处理，
// 主要是管理系统的前端代码，对于资源等（如图片）尽可能不要打包进入程序

package controller

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/elton"
	M "github.com/vicanso/elton/middleware"
)

type (
	// assetCtrl asset ctrl
	assetCtrl  struct{}
	staticFile struct {
		box *packr.Box
	}
)

var (
	assetBox = packr.New("asset", "../web/dist")
)

// Exists 判断文件是否存在
func (sf *staticFile) Exists(file string) bool {
	return sf.box.Has(file)
}

// Get 获取文件内容
func (sf *staticFile) Get(file string) ([]byte, error) {
	return sf.box.Find(file)
}

// Stat 获取文件stat信息
func (sf *staticFile) Stat(file string) os.FileInfo {
	// 文件打包至程序中，因此无file info
	return nil
}

// NewReader 创建读取文件的reader
func (sf *staticFile) NewReader(file string) (io.Reader, error) {
	buf, err := sf.Get(file)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buf), nil
}

func init() {
	g := router.NewGroup("")
	ctrl := assetCtrl{}
	g.GET("/", ctrl.getIndex)
	g.GET("/favicon.{ext}", ctrl.getFavIcon)

	sf := &staticFile{
		box: assetBox,
	}
	g.GET("/static/*", M.NewStaticServe(sf, M.StaticServeConfig{
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
	buf, err := assetBox.Find(file)
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
