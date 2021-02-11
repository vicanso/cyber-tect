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

package service

import (
	"encoding/json"

	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
	jwt "github.com/vicanso/elton-jwt"
)

const (
	// UserSessionInfoKey user session info
	UserSessionInfoKey = "user-session-info"
)

type (
	// UserSession 用户session中的信息
	UserSession struct {
		ctx       *elton.Context
		Token     string   `json:"token,omitempty"`
		Account   string   `json:"account,omitempty"`
		ID        int      `json:"id,omitempty"`
		Roles     []string `json:"roles,omitempty"`
		Groups    []string `json:"groups,omitempty"`
		UpdatedAt string   `json:"updatedAt,omitempty"`
		LoginAt   string   `json:"loginAt,omitempty"`
	}
)

// GetInfo 获取用户信息，
func (us *UserSession) GetInfo() (info UserSession) {
	return *us
}

// IsLogin 判断用户是否已登录
func (us *UserSession) IsLogin() bool {
	info := us.GetInfo()
	return info.Account != ""
}

// SetInfo 设置用户信息
func (us *UserSession) SetInfo(info UserSession) (err error) {
	// 登录时设置登录时间
	if info.Account != "" && info.LoginAt == "" {
		info.LoginAt = util.NowString()
	}
	info.UpdatedAt = util.NowString()
	buf, err := json.Marshal(&info)
	if err != nil {
		return
	}
	us.ctx.Set(jwt.DefaultKey, string(buf))
	return
}

// Destroy 清除用户session
func (us *UserSession) Destroy() error {
	us.ctx.Set(jwt.DefaultKey, "")
	return nil
}

// Refresh 刷新用户session ttl
func (us *UserSession) Refresh() error {
	info := us.GetInfo()
	us.SetInfo(info)
	return nil
}

// NewUserSession 创建新的用户session对象
func NewUserSession(c *elton.Context) *UserSession {
	data := c.GetString(jwt.DefaultKey)
	us := &UserSession{}
	// 忽略错误
	_ = json.Unmarshal([]byte(data), us)
	us.ctx = c
	return us
}
