// Copyright 2021 tree xie
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

package interceptor

import (
	"encoding/json"
	"strings"

	"go.uber.org/atomic"
)

type (
	//  session拦截的数据
	SessionData struct {
		Enabled       bool     `json:"enabled"`
		Message       string   `json:"message"`
		AllowAccount  string   `json:"allowAccount"`
		AllowAccounts []string `json:"allowAccounts"`
		AllowRoutes   []string `json:"allowRoutes"`
	}
)

//  session拦截的配置
var sessionConfig = atomic.Value{}

//  获取session拦截的配置信息
func GetSessionData() (*SessionData, bool) {
	value := sessionConfig.Load()
	if value == nil {
		return nil, false
	}
	data, ok := value.(*SessionData)
	// 如果数据类型不对或非启用状态
	if !ok || !data.Enabled {
		return nil, false
	}
	return data, true
}

// 更新session拦截配置
func UpdateSessionConfig(value string) (err error) {
	// 如果为空则清除
	if value == "" {
		sessionConfig.Store(&SessionData{})
		return
	}

	interData := &SessionData{}
	err = json.Unmarshal([]byte(value), interData)
	if err != nil {
		return
	}
	if len(interData.AllowAccounts) == 0 && interData.AllowAccount != "" {
		interData.AllowAccounts = strings.Split(interData.AllowAccount, ",")
	}
	// 设置为启用
	interData.Enabled = true
	sessionConfig.Store(interData)
	return
}
