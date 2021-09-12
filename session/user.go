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

package session

import (
	"context"
	"encoding/json"

	"github.com/vicanso/elton"
	se "github.com/vicanso/elton-session"
	session "github.com/vicanso/elton-session"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/util"
)

const (
	// UserSessionInfoKey user session info
	UserSessionInfoKey = "user-session-info"
)

type (
	// UserInfo 用户session中的信息
	UserInfo struct {
		// 登录时使用的Token，此字段不返回
		Token string `json:"token"`
		// 用户账号
		Account string `json:"account"`
		// 用户ID
		ID int `json:"id"`
		// 用户角色列表
		Roles []string `json:"roles"`
		// 用户分组列表
		Groups []string `json:"groups"`
		// Session信息更新时间
		UpdatedAt string `json:"updatedAt"`
		// Session信息创建时间
		LoginAt string `json:"loginAt"`
	}
	// UserSession 用户session
	UserSession struct {
		unmarshalDone bool
		se            *se.Session
		info          UserInfo
	}
)

// GetUserInfo 获取用户信息
// 避免修改了session中的数据，因此返回非指针的形式
func (us *UserSession) GetInfo() (UserInfo, error) {
	info := us.info
	if us.unmarshalDone {
		return info, nil
	}
	data := us.se.GetString(UserSessionInfoKey)
	if data == "" {
		data = "{}"
	}
	info = UserInfo{}
	err := json.Unmarshal([]byte(data), &info)
	if err != nil {
		return info, err
	}
	us.info = info
	us.unmarshalDone = true
	return info, err
}

// MustGetInfo 获取用户信息，如果信息获取失败则触发panic，
// 如果前置中间件已保证是登录状态，可以使用此函数，否则禁止使用
func (us *UserSession) MustGetInfo() UserInfo {
	info, err := us.GetInfo()
	if err != nil {
		panic(err)
	}
	return info
}

// IsLogin 判断用户是否已登录
func (us *UserSession) IsLogin() bool {
	info, err := us.GetInfo()
	if err != nil {
		return false
	}
	return info.Account != ""
}

// SetInfo 设置用户信息
func (us *UserSession) SetInfo(ctx context.Context, info UserInfo) error {
	// 登录时设置登录时间
	if info.Account != "" && info.LoginAt == "" {
		info.LoginAt = util.NowString()
	}
	info.UpdatedAt = util.NowString()
	us.info = info
	us.unmarshalDone = true
	buf, err := json.Marshal(&info)
	if err != nil {
		return err
	}
	err = us.se.Set(ctx, UserSessionInfoKey, string(buf))
	if err != nil {
		return err
	}
	return nil
}

// Destroy 清除用户session
func (us *UserSession) Destroy(ctx context.Context) error {
	return us.se.Destroy(ctx)
}

// Refresh 刷新用户session ttl
func (us *UserSession) Refresh(ctx context.Context) error {
	return us.se.Refresh(ctx)
}

// NewUserSession 创建新的用户session对象
func NewUserSession(c *elton.Context) *UserSession {
	if data, ok := c.Get(cs.UserSession); ok {
		us, ok := data.(*UserSession)
		if ok {
			return us
		}
	}
	se, ok := session.Get(c)
	if !ok {
		return nil
	}
	us := &UserSession{
		se: se,
	}
	c.Set(cs.UserSession, us)

	return us
}
