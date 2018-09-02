package umbra

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"
)

type (
	// HTTP http protocol check
	HTTP struct {
		IP          string
		URL         string
		Transport   *http.Transport
		EnableTrace bool
	}
	// HTTPTimeLine http timing
	HTTPTimeLine struct {
		Start                time.Time
		DNSStart             time.Time
		DNSDone              time.Time
		ConnectStart         time.Time
		ConnectDone          time.Time
		GotConnect           time.Time
		GotFirstResponseByte time.Time
		TLSHandshakeStart    time.Time
		TLSHandshakeDone     time.Time
		Done                 time.Time
	}
	// HTTPTimeLineStats http timeline stats
	HTTPTimeLineStats struct {
		DNSLookup        string `json:"dnsLookup,omitempty"`
		TCPConnection    string `json:"tcpConnection,omitempty"`
		TLSHandshake     string `json:"tlsHandshake,omitempty"`
		ServerProcessing string `json:"serverProcessing,omitempty"`
		ContentTransfer  string `json:"contentTransfer,omitempty"`
		Total            string `json:"total,omitempty"`
	}
)

// Stats get the stats of time line
func (tl *HTTPTimeLine) Stats() (stats *HTTPTimeLineStats) {
	stats = &HTTPTimeLineStats{}
	if !tl.DNSStart.IsZero() {
		stats.DNSLookup = tl.DNSDone.Sub(tl.DNSStart).String()
	}
	stats.TCPConnection = tl.ConnectDone.Sub(tl.ConnectStart).String()
	if !tl.TLSHandshakeStart.IsZero() {
		stats.TLSHandshake = tl.TLSHandshakeDone.Sub(tl.TLSHandshakeStart).String()
	}
	stats.ServerProcessing = tl.GotFirstResponseByte.Sub(tl.GotConnect).String()
	if tl.Done.IsZero() {
		tl.Done = time.Now()
	}
	stats.ContentTransfer = tl.Done.Sub(tl.GotFirstResponseByte).String()
	stats.Total = tl.Done.Sub(tl.Start).String()
	return
}

// NewClientTrace http client trace
func NewClientTrace() (trace *httptrace.ClientTrace, tl *HTTPTimeLine) {
	tl = &HTTPTimeLine{
		Start: time.Now(),
	}
	trace = &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			tl.DNSStart = time.Now()
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			tl.DNSDone = time.Now()
		},
		ConnectStart: func(_, _ string) {
			tl.ConnectStart = time.Now()
		},
		ConnectDone: func(net, addr string, err error) {
			tl.ConnectDone = time.Now()
		},
		GotConn: func(_ httptrace.GotConnInfo) {
			tl.GotConnect = time.Now()
		},
		GotFirstResponseByte: func() {
			tl.GotFirstResponseByte = time.Now()
		},
		TLSHandshakeStart: func() {
			tl.TLSHandshakeStart = time.Now()
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			tl.TLSHandshakeDone = time.Now()
		},
	}
	return
}

// Check check the http is healthy
func (h *HTTP) Check() (healthy bool, extra map[string]interface{}, err error) {
	if h.Transport == nil {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
		}
		// 如果指定了IP，则将检测url中的host解析至此IP
		if h.IP != "" {
			info, err := url.Parse(h.URL)
			if err != nil {
				return false, nil, err
			}
			hostname := info.Hostname()
			dialer := &net.Dialer{}
			transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
				// 直接将host替换为ip地址，则无需DNS查询
				addr = strings.Replace(addr, hostname, h.IP, 1)
				return dialer.DialContext(ctx, network, addr)
			}
		}
		h.Transport = transport
	}

	client := &http.Client{
		Transport: h.Transport,
		Timeout:   defaultTimeout,
	}
	req, err := http.NewRequest(http.MethodGet, h.URL, nil)
	if err != nil {
		return
	}
	var tl *HTTPTimeLine
	var trace *httptrace.ClientTrace
	if h.EnableTrace {
		trace, tl = NewClientTrace()
		req = req.WithContext(httptrace.WithClientTrace(context.Background(), trace))
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	extra = make(map[string]interface{})
	if tl != nil {
		extra["stats"] = tl.Stats()
	}
	statusCode := resp.StatusCode
	if statusCode >= http.StatusOK && statusCode < http.StatusBadRequest {
		healthy = true
	}
	extra["status"] = statusCode
	if !healthy {
		extra["body"] = string(body)
	}
	return
}

// GetDescription get the description of http checker
func (h *HTTP) GetDescription() (description map[string]interface{}) {
	description = make(map[string]interface{})
	description["type"] = TypeHTTP
	description["url"] = h.URL
	if h.IP != "" {
		description["ip"] = h.IP
	}
	return
}
