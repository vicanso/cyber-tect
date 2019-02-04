package umbra

import (
	"strings"
)

type (
	// DNSServer dns server
	DNSServer struct {
		Name string
		Addr string
	}
)

var (
	defaultDNSServers = make([]*DNSServer, 0)
)

func init() {
	defaultServers := []string{
		"DNSPod || 119.29.29.29",
		"AliDNS || 223.5.5.5,223.6.6.6",
		"CNNIC SDNS || 1.2.4.8,210.2.4.8",
		"114 DNS || 114.114.114.114,114.114.115.115",
		"oneDNS || 112.124.47.27,114.215.126.16",
		"DNS 派 || 101.226.4.6,218.30.118.6,123.125.81.6,140.207.198.6",
		"GOOGLE || 8.8.8.8,8.8.4.4",
		"OPENDNS || 208.67.222.222,208.67.220.220",
		"QUAD9 || 9.9.9.9,149.112.112.112",
	}

	for _, str := range defaultServers {
		arr := strings.Split(str, " || ")
		ips := strings.SplitN(arr[1], ",", -1)
		for _, ip := range ips {
			defaultDNSServers = append(defaultDNSServers, &DNSServer{
				Name: arr[0],
				Addr: ip,
			})
		}
	}
}

// GetDefaultDNSServers get default dns servers
func GetDefaultDNSServers() []*DNSServer {
	return defaultDNSServers
}
