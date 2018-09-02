package umbra

import "testing"

func TestDNS(t *testing.T) {
	dns := DNS{
		Server:   "114.114.114.114:53",
		Hostname: "www.baidu.com",
	}
	desc := dns.GetDescription()
	if desc["type"] != TypeDNS ||
		desc["server"] != dns.Server ||
		desc["hostname"] != dns.Hostname {
		t.Fatalf("get description fail")
	}
	healthy, extra, err := dns.Check()
	if err != nil {
		t.Fatalf("dns check fail, %v", err)
	}
	if !healthy {
		t.Fatalf("dns check fail, it should be healthy")
	}

	if extra["stats"] == nil ||
		extra["answers"] == nil {
		t.Fatalf("tcp check extra infofail")
	}
}
