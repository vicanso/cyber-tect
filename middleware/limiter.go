package middleware

import (
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/config"
	"github.com/vicanso/cyber-tect/global"
	"github.com/vicanso/hes"
)

var (
	errTooManyRequest = &hes.Error{
		StatusCode: http.StatusTooManyRequests,
		Message:    "too many request",
		Category:   errLimitCategory,
	}
	errTooFrequently = &hes.Error{
		StatusCode: http.StatusBadRequest,
		Message:    "request to frequently",
		Category:   errLimitCategory,
	}
)

const (
	defaultRequestLimit      = 2048
	concurrentLimitKeyPrefix = "mid-concurrent-limit"
	ipLimitKeyPrefix         = "mid-ip-limit"
	errLimitCategory         = "request-limit"
)

// NewLimiter create a limit middleware
func NewLimiter() cod.Handler {
	maxRequestLimit := uint32(config.GetIntDefault("requestLimit", defaultRequestLimit))
	var connectingCount uint32
	errTooManyRequest.Message += ("(" + strconv.Itoa(int(maxRequestLimit)) + ")")
	return func(c *cod.Context) (err error) {
		// 处理请求数+1/-1
		defer atomic.AddUint32(&connectingCount, ^uint32(0))
		v := atomic.AddUint32(&connectingCount, 1)
		global.SaveConnectingCount(v)
		if v >= maxRequestLimit {
			err = errTooManyRequest
			return
		}
		return c.Next()
	}
}
