package umbra

import (
	"testing"
)

const (
	baiduIPAddr = "14.215.177.38"
)

func TestHTTPCheck(t *testing.T) {
	h := &HTTP{
		URL:         "https://www.baidu.com/",
		IP:          baiduIPAddr,
		EnableTrace: true,
	}
	desc := h.GetDescription()
	if desc["url"] != h.URL || desc["ip"] != h.IP || desc["type"] != TypeHTTP {
		t.Fatalf("get description fail")
	}
	healthy, extra, err := h.Check()
	if err != nil {
		t.Fatalf("http check fail, %v", err)
	}
	if !healthy {
		t.Fatalf("http check fail, it should be healthy")
	}
	if extra["status"] != 200 ||
		extra["stats"] == nil {
		t.Fatalf("http check extra info fail")
	}
}
