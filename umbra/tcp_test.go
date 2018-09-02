package umbra

import (
	"testing"
)

func TestTCP(t *testing.T) {
	tcp := &TCP{
		IP:   baiduIPAddr,
		Port: 443,
	}

	desc := tcp.GetDescription()
	if desc["ip"] != tcp.IP || desc["port"] != tcp.Port || desc["type"] != TypeTCP {
		t.Fatalf("get description fail")
	}

	healthy, extra, err := tcp.Check()
	if err != nil {
		t.Fatalf("tcp check fail, %v", err)
	}
	if !healthy {
		t.Fatalf("tcp check fail, it should be healthy")
	}
	if extra["stats"] == nil {
		t.Fatalf("tcp check extra info fail")
	}
}
