package umbra

import (
	"context"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"

	"github.com/vicanso/dusk"
)

type (
	// HTTP http protocol check
	HTTP struct {
		IP         string
		URL        string
		statusCode int
		ht         *dusk.HTTPTrace
		header     http.Header
	}
	// HTTPCheckResult http check result
	HTTPCheckResult struct {
		IP         string   `json:"ip,omitempty"`
		URL        string   `json:"url,omitempty"`
		StatusCode int      `json:"statusCode,omitempty"`
		Header     []string `json:"header,omitempty"`
		Healthy    bool     `json:"healthy,omitempty"`
		Err        error    `json:"err,omitempty"`
		Message    string   `json:"message,omitempty"`

		Host           string        `json:"host,omitempty"`
		Addrs          []string      `json:"addrs,omitempty"`
		Network        string        `json:"network,omitempty"`
		Addr           string        `json:"addr,omitempty"`
		Reused         bool          `json:"reused,omitempty"`
		WasIdle        bool          `json:"wasIdle,omitempty"`
		IdleTime       time.Duration `json:"idleTime,omitempty"`
		Protocol       string        `json:"protocol,omitempty"`
		TLSVersion     string        `json:"tlsVersion,omitempty"`
		TLSResume      bool          `json:"tlsResume,omitempty"`
		TLSCipherSuite string        `json:"tlsCipherSuite,omitempty"`

		TimeLineStats *dusk.HTTPTimelineStats `json:"timeLineStats,omitempty"`
	}
)

const userAgentChrome = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36"
const acceptEncodingChrome = "gzip, deflate, br"

// Check check the http is healthy
func (h *HTTP) Check() (healthy bool, err error) {
	transport := &http.Transport{
		DisableCompression: true,
	}
	// 如果指定了IP，则将检测url中的host解析至此IP
	if h.IP != "" {
		info, err := url.Parse(h.URL)
		if err != nil {
			return false, err
		}
		hostname := info.Hostname()
		dialer := &net.Dialer{}
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			// 直接将host替换为ip地址，则无需DNS查询
			addr = strings.Replace(addr, hostname, h.IP, 1)
			return dialer.DialContext(ctx, network, addr)
		}
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   defaultTimeout,
	}
	req, err := http.NewRequest(http.MethodGet, h.URL, nil)
	req.Header.Set("User-Agent", userAgentChrome)
	req.Header.Set("Accept-Encoding", acceptEncodingChrome)
	if err != nil {
		return
	}
	trace, ht := dusk.NewClientTrace()
	ctx := httptrace.WithClientTrace(context.Background(), trace)
	h.ht = ht
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	statusCode := resp.StatusCode
	if statusCode >= http.StatusOK && statusCode < http.StatusBadRequest {
		healthy = true
	}
	h.statusCode = statusCode
	h.header = resp.Header
	ht.Done = time.Now()

	return
}

// GetCheckResult get check result
func (h *HTTP) GetCheckResult() (result HTTPCheckResult) {
	healthy, err := h.Check()
	result = HTTPCheckResult{
		IP:         h.IP,
		URL:        h.URL,
		StatusCode: h.statusCode,
		Healthy:    healthy,
		Err:        err,
	}
	if err != nil {
		result.Message = err.Error()
	}
	header := make([]string, 0)
	for name, values := range h.header {
		for _, v := range values {
			header = append(header, name+": "+v)
		}
	}
	result.Header = header
	ht := h.ht
	if ht != nil {
		result.TimeLineStats = ht.Stats()
		result.Host = ht.Host
		result.Addrs = ht.Addrs
		result.Network = ht.Network
		result.Addr = ht.Addr
		result.Reused = ht.Reused
		result.WasIdle = ht.WasIdle
		result.IdleTime = ht.IdleTime
		result.Protocol = ht.Protocol
		result.TLSVersion = ht.TLSVersion
		result.TLSResume = ht.TLSResume
		result.TLSCipherSuite = ht.TLSCipherSuite
	}
	return
}
