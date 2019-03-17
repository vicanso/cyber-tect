package middleware

import (
	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/router"
	"github.com/vicanso/cyber-tect/util"
)

const (
	xResponseID = "X-Response-Id"
)

// NewEntry create an entry middleware
func NewEntry() cod.Handler {
	return func(c *cod.Context) (err error) {
		// 生成context id
		c.ID = util.GenUlid()
		c.SetHeader(xResponseID, c.ID)

		// 设置所有的请求响应默认都为no cache
		c.NoCache()
		router.AddRouteCount(c.Request.Method, c.Route)

		return c.Next()
	}
}
