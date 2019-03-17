package umbra

import (
	"testing"
)

func TestDNSCheck(t *testing.T) {
	dnsServer := &DNSServer{
		Name: "google",
		Addr: "8.8.8.8",
	}
	servers := make([]*DNSServer, 0)
	servers = append(servers, dnsServer)

	results := CheckDNSServers("www.baidu.com", servers)

	result := results[0]
	if !result.Healthy ||
		len(result.IPAddr) == 0 ||
		result.Err != nil {
		t.Fatalf("dns check fail, %v", result.Err)
	}
}
