package middleware

import (
	"net/http"
	"time"

	"github.com/vicanso/cod"
	"github.com/vicanso/hes"
)

var (
	errQueryNotAllow = &hes.Error{
		StatusCode: http.StatusBadRequest,
		Message:    "query is not allowed",
		Category:   "common-validate",
	}
)

// NoQuery no query middleware
func NoQuery(c *cod.Context) (err error) {
	if c.Request.URL.RawQuery != "" {
		err = errQueryNotAllow
		return
	}
	return c.Next()
}

// WaitFor at least wait for duration
func WaitFor(d time.Duration) cod.Handler {
	ns := d.Nanoseconds()
	return func(c *cod.Context) (err error) {
		start := time.Now()
		err = c.Next()
		use := time.Now().UnixNano() - start.UnixNano()
		// 无论成功还是失败都wait for
		if use < ns {
			time.Sleep(time.Duration(ns-use) * time.Nanosecond)
		}
		return
	}
}
