package routerconcurrency

import (
	"context"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	syncAtomic "sync/atomic"
	"unsafe"

	"github.com/vicanso/elton"
	"github.com/vicanso/cybertect/email"
	"github.com/vicanso/cybertect/log"
	"go.uber.org/atomic"
	"go.uber.org/ratelimit"
)

type (
	routerConcurrency struct {
		Route  string `json:"route"`
		Method string `json:"method"`
		Max    uint32 `json:"max"`
		// RateLimit 频率限制，如100/s
		RateLimit string `json:"rateLimit"`
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
	currentRCLimiter = &rcLimiter{}

	// 无频率限制
	routerRateUnlimited = &routerRateLimit{
		Limiter: ratelimit.NewUnlimited(),
	}
)

// SetRateLimiter 设置频率限制
func (rc *RouterConcurrency) setRateLimiter(limit string) {
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
func (l *rcLimiter) IncConcurrency(key string) (uint32, uint32) {
	// 该map仅初始化一次，因此无需要考虑锁
	r, ok := l.m[key]
	if !ok {
		return 0, 0
	}
	current := r.Current.Inc()
	max := r.Max.Load()
	// 如果设置为0或已超出最大并发限制，则直接返回
	if max == 0 || current > max {
		return current, max
	}
	r.Take()
	return current, max
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

// InitLimiter 初始路由并发限制
func InitLimiter(routers []elton.RouterInfo) {
	m := make(map[string]*RouterConcurrency)
	for _, item := range routers {
		m[item.Method+" "+item.Route] = &RouterConcurrency{}
	}
	currentRCLimiter.m = m
}

// GetLimiter 获取路由并发限制器
func GetLimiter() *rcLimiter {
	return currentRCLimiter
}

// Update 更新路由并发数
func Update(arr []string) {
	concurrencyConfigList := make([]*routerConcurrency, 0)
	for _, str := range arr {
		v := &routerConcurrency{}
		err := json.Unmarshal([]byte(str), v)
		if err != nil {
			log.Error(context.Background()).
				Err(err).
				Msg("router concurrency config is invalid")
			email.AlarmError("router concurrency config is invalid:" + err.Error())
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
					r.setRateLimiter(item.RateLimit)
				}
			}
		}
		// 如果未配置，则设置为限制0（无限制）
		if !found {
			r.Max.Store(0)
		}
	}
}

// List 获取路由并发限制数
func List() map[string]uint32 {
	result := make(map[string]uint32)
	for key, r := range currentRCLimiter.m {
		v := r.Max.Load()
		if v != 0 {
			result[key] = v
		}
	}
	return result
}
