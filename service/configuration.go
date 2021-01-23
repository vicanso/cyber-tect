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
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/configuration"
	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
	"go.uber.org/zap"
)

type (
	// ConfigurationSrv 配置的相关函数
	ConfigurationSrv struct{}

	// SessionInterceptorData session拦截的数据
	SessionInterceptorData struct {
		Message       string   `json:"message,omitempty"`
		AllowAccounts []string `json:"allowAccounts,omitempty"`
	}

	// CurrentValidConfiguration 当前有效配置
	CurrentValidConfiguration struct {
		UpdatedAt          time.Time               `json:"updatedAt,omitempty"`
		MockTime           string                  `json:"mockTime,omitempty"`
		IPBlockList        []string                `json:"ipBlockList,omitempty"`
		SignedKeys         []string                `json:"signedKeys,omitempty"`
		RouterConcurrency  map[string]uint32       `json:"routerConcurrency,omitempty"`
		RouterMock         map[string]RouterConfig `json:"routerMock,omitempty"`
		SessionInterceptor *SessionInterceptorData `json:"sessionInterceptor,omitempty"`
	}
)

var (
	sessionSignedKeys = new(elton.RWMutexSignedKeys)
	// sessionInterceptorConfig session拦截的配置
	sessionInterceptorConfig = new(sync.Map)
)

// 配置刷新时间
var configurationRefreshedAt time.Time

const (
	sessionInterceptorKey = "sessionInterceptor"
)

func init() {
	sessionConfig := config.GetSessionConfig()
	// session中用于cookie的signed keys
	sessionSignedKeys.SetKeys(sessionConfig.Keys)
}

// GetSignedKeys 获取用于cookie加密的key列表
func GetSignedKeys() elton.SignedKeysGenerator {
	return sessionSignedKeys
}

// GetCurrentValidConfiguration 获取当前有效配置
func GetCurrentValidConfiguration() *CurrentValidConfiguration {
	interData, _ := GetSessionInterceptorData()
	result := &CurrentValidConfiguration{
		UpdatedAt:         configurationRefreshedAt,
		MockTime:          util.GetMockTime(),
		IPBlockList:       GetIPBlockList(),
		SignedKeys:        sessionSignedKeys.GetKeys(),
		RouterConcurrency: GetRouterConcurrency(),
		RouterMock:        GetRouterMockConfig(),
	}
	if interData != nil {
		v := *interData
		result.SessionInterceptor = &v
	}
	return result
}

// GetSessionInterceptorMessage 获取session拦截的配置信息
func GetSessionInterceptorData() (*SessionInterceptorData, bool) {
	value, ok := sessionInterceptorConfig.Load(sessionInterceptorKey)
	if !ok {
		return nil, false
	}
	data, ok := value.(*SessionInterceptorData)
	if !ok {
		return nil, false
	}
	return data, true
}

// available 获取可用的配置
func (*ConfigurationSrv) available() ([]*ent.Configuration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	now := util.Now()
	return helper.EntGetClient().Configuration.Query().
		Where(configuration.Status(schema.StatusEnabled)).
		Where(configuration.StartedAtLT(now)).
		Where(configuration.EndedAtGT(now)).
		Order(ent.Desc(configuration.FieldUpdatedAt)).
		All(ctx)
}

// Refresh 刷新配置
func (srv *ConfigurationSrv) Refresh() (err error) {
	configs, err := srv.available()
	if err != nil {
		return
	}
	configurationRefreshedAt = time.Now()
	var mockTimeConfig *ent.Configuration
	routerConcurrencyConfigs := make([]string, 0)
	routerConfigs := make([]string, 0)
	var signedKeys []string
	blockIPList := make([]string, 0)
	sessionInterceptorValue := ""
	for _, item := range configs {
		switch item.Category {
		case schema.ConfigurationCategoryMockTime:
			// 由于排序是按更新时间，因此取最新的记录
			if mockTimeConfig == nil {
				mockTimeConfig = item
			}
		case schema.ConfigurationCategoryBlockIP:
			blockIPList = append(blockIPList, item.Data)
		case schema.ConfigurationCategorySignedKey:
			// 按更新时间排序，因此如果已获取则不需要再更新
			if len(signedKeys) == 0 {
				signedKeys = strings.Split(item.Data, ",")
			}
		case schema.ConfigurationCategoryRouterConcurrency:
			routerConcurrencyConfigs = append(routerConcurrencyConfigs, item.Data)
		case schema.ConfigurationCategoryRouter:
			routerConfigs = append(routerConfigs, item.Data)
		case schema.ConfigurationCategorySessionInterceptor:
			// 按更新时间排序，因此如果已获取则不需要再更新
			if sessionInterceptorValue == "" {
				sessionInterceptorValue = item.Data
			}
		}
	}

	// 设置session interceptor的拦截信息
	if sessionInterceptorValue == "" {
		sessionInterceptorConfig.Delete(sessionInterceptorKey)
	} else {
		interData := &SessionInterceptorData{}
		err := json.Unmarshal([]byte(sessionInterceptorValue), interData)
		if err != nil {
			log.Default().Error("session interceptor config is invalid",
				zap.Error(err),
			)
			AlarmError("session interceptor config is invalid:" + err.Error())
		}
		sessionInterceptorConfig.Store(sessionInterceptorKey, interData)
	}

	// 如果未配置mock time，则设置为空
	if mockTimeConfig == nil {
		util.SetMockTime("")
	} else {
		util.SetMockTime(mockTimeConfig.Data)
	}

	// 如果数据库中未配置，则使用默认配置
	if len(signedKeys) == 0 {
		sessionConfig := config.GetSessionConfig()
		sessionSignedKeys.SetKeys(sessionConfig.Keys)
	} else {
		sessionSignedKeys.SetKeys(signedKeys)
	}

	// 更新router configs
	updateRouterMockConfigs(routerConfigs)

	// 重置IP拦截列表
	ResetIPBlocker(blockIPList)

	// 重置路由并发限制
	ResetRouterConcurrency(routerConcurrencyConfigs)
	return
}
