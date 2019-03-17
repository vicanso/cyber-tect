package helper

import (
	"net/http"
	"net/url"
	"time"

	"github.com/vicanso/cod"
	"github.com/vicanso/dusk"
	"github.com/vicanso/cyber-tect/log"
	"github.com/vicanso/cyber-tect/util"
	"github.com/vicanso/hes"
	"go.uber.org/zap"

	jsoniter "github.com/json-iterator/go"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	logger = log.Default()
	// DefaultHTTPClient default http client
	DefaultHTTPClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:           100,
			IdleConnTimeout:        90 * time.Second,
			TLSHandshakeTimeout:    5 * time.Second,
			ExpectContinueTimeout:  1 * time.Second,
			MaxResponseHeaderBytes: 5 * 1024,
		},
	}
)

const (
	errCategoryHTTPRequest = "http-request"
	contextID              = "cid"

	xForwardedForHeader = "X-Forwarded-For"
)

// httpConvertResponse convert http response
func httpConvertResponse(resp *http.Response, d *dusk.Dusk) (newResp *http.Response, newErr error) {
	statusCode := resp.StatusCode
	if statusCode < 400 {
		return
	}
	// 对于状态码大于400的，转化为 hes.Error
	he := &hes.Error{
		StatusCode: statusCode,
		Category:   json.Get(d.Body, "category").ToString(),
		Message:    json.Get(d.Body, "message").ToString(),
	}
	if he.Category != "" {
		he.Category = errCategoryHTTPRequest + "-" + he.Category
	} else {
		he.Category = errCategoryHTTPRequest
	}
	if he.Message == "" {
		he.Message = "unknown error"
	}
	return nil, he
}

// httpDoneEvent http请求完成的触发，用于统计、日志等输出
func httpDoneEvent(d *dusk.Dusk) error {
	req := d.Request
	resp := d.Response
	err := d.Err
	uri := req.URL.RequestURI()
	ht := d.GetHTTPTrace()
	use := ""
	if ht != nil {
		use = ht.Stats().Total.String()
	}
	statusCode := 0
	if err != nil {
		he, ok := err.(*hes.Error)
		if ok {
			statusCode = he.StatusCode
		}
	}
	if resp != nil {
		statusCode = resp.StatusCode
	}
	cid := ""
	cidValue := d.GetValue(contextID)
	if cidValue != nil {
		cid = cidValue.(string)
	}

	// TODO 是否将POST参数也记录（有可能会有敏感信息）
	// TODO 是否将响应数据输出（有可能敏感信息以及数据量较大），或者写入缓存数据库，保存较短时间方便排查
	if resp == nil || err != nil {
		logger.Error("http request fail",
			zap.String("method", req.Method),
			zap.String("host", req.Host),
			zap.String("uri", uri),
			zap.String("cid", cid),
			zap.Int("status", statusCode),
			zap.String("use", use),
			zap.Error(err),
		)
		return nil
	}
	logger.Info("http request done",
		zap.String("method", req.Method),
		zap.String("host", req.Host),
		zap.String("uri", uri),
		zap.String("cid", cid),
		zap.Int("status", statusCode),
		zap.String("use", use),
	)
	return nil
}

// httpErrorConvert convert http error
func httpErrorConvert(err error, d *dusk.Dusk) error {
	he, ok := err.(*hes.Error)
	resp := d.Response
	req := d.Request
	if !ok {
		he = hes.NewWithError(err)
		statusCode := http.StatusInternalServerError
		if resp != nil {
			statusCode = resp.StatusCode
		}
		if ue, ok := err.(*url.Error); ok {
			// 请求超时中断
			if ue.Timeout() {
				statusCode = http.StatusRequestTimeout
			}
		}
		he.StatusCode = statusCode
		he.Category = errCategoryHTTPRequest
	}
	// 仅在测试中输出请求 url至 hes中（避免将重要信息输出）
	if !util.IsProduction() {
		extra := he.Extra
		if extra == nil {
			extra = make(map[string]interface{})
		}
		url := req.URL
		extra["uri"] = url.RequestURI()
		extra["host"] = url.Host
		extra["method"] = req.Method
		he.Extra = extra
	}
	return he
}

func initDusk(d *dusk.Dusk, c *cod.Context) {
	if c != nil && c.ID != "" {
		d.SetValue(contextID, c.ID)
		// 设置x-forwarded-for
		v := c.GetRequestHeader(xForwardedForHeader)
		if v == "" {
			v = c.RealIP()
		}
		d.Set(xForwardedForHeader, v)
	}
	d.SetClient(DefaultHTTPClient)
	d.EnableTrace()
	d.OnResponseSuccess(httpConvertResponse)
	d.OnError(httpErrorConvert)
	d.OnDone(httpDoneEvent)
}

// NewRequestWithContext new request with context
func NewRequestWithContext(method, url string, c *cod.Context) (d *dusk.Dusk) {
	switch method {
	case http.MethodGet:
		d = dusk.Get(url)
	case http.MethodPost:
		d = dusk.Post(url)
	case http.MethodPatch:
		d = dusk.Patch(url)
	case http.MethodDelete:
		d = dusk.Delete(url)
	}
	if d != nil {
		initDusk(d, c)
	}
	return d
}

// GetWithContext get request with context
func GetWithContext(url string, c *cod.Context) *dusk.Dusk {
	return NewRequestWithContext(http.MethodGet, url, c)
}

// PostWithContext post request with context
func PostWithContext(url string, c *cod.Context) *dusk.Dusk {
	return NewRequestWithContext(http.MethodPost, url, c)
}

// PutWithContext put request with context
func PutWithContext(url string, c *cod.Context) *dusk.Dusk {
	return NewRequestWithContext(http.MethodPut, url, c)
}

// PatchWithContext patch request with context
func PatchWithContext(url string, c *cod.Context) *dusk.Dusk {
	return NewRequestWithContext(http.MethodPatch, url, c)
}

// DeleteWithContext delete request with context
func DeleteWithContext(url string, c *cod.Context) *dusk.Dusk {
	return NewRequestWithContext(http.MethodDelete, url, c)
}
