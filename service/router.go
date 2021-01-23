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
	"regexp"
	"strconv"
	"strings"
	"sync"
	syncAtomic "sync/atomic"
	"unsafe"

	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/elton"
	"go.uber.org/atomic"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
)

type (
	// RouterConfig 路由配置信息
	RouterConfig struct {
		Route      string `json:"route,omitempty"`
		Method     string `json:"method,omitempty"`
		Status     int    `json:"status,omitempty"`
		CotentType string `json:"cotentType,omitempty"`
		Response   string `json:"response,omitempty"`
		// DelaySeconds 延时，单位秒
		DelaySeconds int    `json:"delaySeconds,omitempty"`
		URL          string `json:"url,omitempty"`
	}
	routerConcurrencyConfig struct {
		Route  string `json:"route,omitempty"`
		Method string `json:"method,omitempty"`
		Max    uint32 `json:"max,omitempty"`
		// RateLimit 频率限制，如100/s
		RateLimit string `json:"rateLimit,omitempty"`
	}
	// routerRateLimit 路由频率限制
	routerRateLimit struct {
		Limiter ratelimit.Limiter
	}
	// RouterConcurrency 路由并发配置
	RouterConcurrency struct {
		Route     string
		Method    string
		Current   atomic.Uint32
		Max       atomic.Uint32
		RateLimit atomic.String
		// limit 保存routerRateLimit对象
		limit unsafe.Pointer
	}
	// rcLimiter 路由请求限制
	rcLimiter struct {
		m map[string]*RouterConcurrency
	}
)

var (
	routerMutex          = new(sync.RWMutex)
	currentRouterConfigs map[string]*RouterConfig
	currentRCLimiter     = &rcLimiter{}

	// 无频率限制
	routerRateUnlimited = &routerRateLimit{
		Limiter: ratelimit.NewUnlimited(),
	}
)

// SetRateLimiter 设置频率限制
func (rc *RouterConcurrency) SetRateLimiter(limit string) {
	rc.RateLimit.Store(limit)
	reg := regexp.MustCompile(`(\d+)/s`)
	rate := 0
	if reg.MatchString(limit) {
		result := reg.FindStringSubmatch(limit)
		if len(result) == 2 {
			rate, _ = strconv.Atoi(result[1])
		}
	}
	// 如果未设置限制，则使用无限制频率
	if rate <= 0 {
		syncAtomic.StorePointer(&rc.limit, unsafe.Pointer(routerRateUnlimited))
		return
	}
	rrl := &routerRateLimit{
		Limiter: ratelimit.New(rate),
	}
	syncAtomic.StorePointer(&rc.limit, unsafe.Pointer(rrl))
}

// Take 执行一次频率限制，此执行会根据当时频率延时
func (rc *RouterConcurrency) Take() {
	p := syncAtomic.LoadPointer(&rc.limit)
	if p == nil {
		return
	}
	limit := (*routerRateLimit)(p)
	limit.Limiter.Take()
}

// IncConcurrency 当前路由处理数+1
func (l *rcLimiter) IncConcurrency(key string) (current uint32, max uint32) {
	// 该map仅初始化一次，因此无需要考虑锁
	r, ok := l.m[key]
	if !ok {
		return
	}
	current = r.Current.Inc()
	max = r.Max.Load()
	// 如果设置为0或已超出最大并发限制，则直接返回
	if max == 0 || current > max {
		return
	}
	r.Take()
	return
}

// DecConcurrency 当前路由处理数-1
func (l *rcLimiter) DecConcurrency(key string) {
	r, ok := l.m[key]
	if !ok {
		return
	}
	r.Current.Dec()
}

// GetConcurrency 获取当前路由处理数
func (l *rcLimiter) GetConcurrency(key string) uint32 {
	r, ok := l.m[key]
	if !ok {
		return 0
	}
	return r.Current.Load()
}

// GetStats 获取统计
func (l *rcLimiter) GetStats() map[string]uint32 {
	result := make(map[string]uint32)
	for key, r := range l.m {
		result[key] = r.Current.Load()
	}
	return result
}

// 更新router config配置
func updateRouterMockConfigs(configs []string) {
	result := make(map[string]*RouterConfig)
	for _, item := range configs {
		v := &RouterConfig{}
		err := json.Unmarshal([]byte(item), v)
		if err != nil {
			log.Default().Error("router config is invalid",
				zap.Error(err),
			)
			AlarmError("router config is invalid:" + err.Error())
			continue
		}
		// 如果未配置Route或者method的则忽略
		if v.Route == "" || v.Method == "" {
			continue
		}
		result[v.Method+v.Route] = v
	}
	routerMutex.Lock()
	defer routerMutex.Unlock()
	currentRouterConfigs = result
}

// RouterGetConfig 获取路由配置
func RouterGetConfig(method, route string) *RouterConfig {
	routerMutex.RLock()
	defer routerMutex.RUnlock()
	return currentRouterConfigs[method+route]
}

// GetRouterMockConfig 获取路由mock配置
func GetRouterMockConfig() map[string]RouterConfig {
	routerMutex.RLock()
	defer routerMutex.RUnlock()
	result := make(map[string]RouterConfig)
	for key, value := range currentRouterConfigs {
		result[key] = *value
	}
	return result
}

// InitRouterConcurrencyLimiter 初始路由并发限制
func InitRouterConcurrencyLimiter(routers []elton.RouterInfo) {
	m := make(map[string]*RouterConcurrency)
	for _, item := range routers {
		m[item.Method+" "+item.Route] = &RouterConcurrency{}
	}
	currentRCLimiter.m = m
}

// GetRouterConcurrencyLimiter 获取路由并发限制器
func GetRouterConcurrencyLimiter() *rcLimiter {
	return currentRCLimiter
}

// ResetRouterConcurrency 重置路由并发数
func ResetRouterConcurrency(arr []string) {
	concurrencyConfigList := make([]*routerConcurrencyConfig, 0)
	for _, str := range arr {
		v := &routerConcurrencyConfig{}
		err := json.Unmarshal([]byte(str), v)
		if err != nil {
			log.Default().Error("router concurrency config is invalid",
				zap.Error(err),
			)
			AlarmError("router concurrency config is invalid:" + err.Error())
			continue
		}
		concurrencyConfigList = append(concurrencyConfigList, v)
	}
	for key, r := range currentRCLimiter.m {
		keys := strings.Split(key, " ")
		if len(keys) != 2 {
			continue
		}
		found := false
		for _, item := range concurrencyConfigList {
			if item.Method == keys[0] && item.Route == keys[1] {
				found = true
				// 设置并发请求量
				r.Max.Store(item.Max)
				// 获取rate limit配置，如果有调整则需要重新设置
				prevLimitDesc := r.RateLimit.Load()
				if prevLimitDesc != item.RateLimit {
					r.SetRateLimiter(item.RateLimit)
				}
			}
		}
		// 如果未配置，则设置为限制0（无限制）
		if !found {
			r.Max.Store(0)
		}
	}
}

// GetRouterConcurrency 获取路由并发限制数
func GetRouterConcurrency() map[string]uint32 {
	result := make(map[string]uint32)
	for key, r := range currentRCLimiter.m {
		v := r.Max.Load()
		if v != 0 {
			result[key] = v
		}
	}
	return result
}
