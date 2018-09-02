package umbra

import (
	"testing"
)

func TestPing(t *testing.T) {
	p := Ping{
		IP: baiduIPAddr,
	}
	desc := p.GetDescription()
	if desc["ip"] != p.IP ||
		desc["type"] != TypePing {
		t.Fatalf("get description fail")
	}
	healthy, extra, err := p.Check()
	if err != nil {
		t.Fatalf("ping check fail, %v", err)
	}
	if !healthy {
		t.Fatalf("ping check fail, it should be healthy")
	}
	if extra["stats"] == nil {
		t.Fatalf("ping check extra info fail")
	}

}
