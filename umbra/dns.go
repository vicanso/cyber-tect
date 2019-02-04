package umbra

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type (
	// DNS dns checker
	DNS struct {
		// Server dns server(ip:port)
		Server string
		// Hostname hostname
		Hostname string
		// ip addr
		IPAddr []net.IPAddr
		// resolve time consuming
		TimeConsuming time.Duration
	}
	// DNSCheckResult dns check result
	DNSCheckResult struct {
		Name          string
		Addr          string
		Hostname      string
		IPAddr        []string
		TimeConsuming time.Duration
		Healthy       bool
		Err           error
	}
)

// Check check the dns resolve
func (d *DNS) Check() (healthy bool, err error) {
	dnsServer := d.Server
	if !strings.Contains(dnsServer, ":") {
		dnsServer += ":53"
	}
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dia := net.Dialer{}
			return dia.DialContext(ctx, "udp", dnsServer)
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	startedAt := time.Now()
	ipAddr, err := r.LookupIPAddr(ctx, d.Hostname)
	d.TimeConsuming = time.Since(startedAt)
	if err != nil {
		return
	}
	d.IPAddr = ipAddr
	healthy = true
	return
}

// GetCheckResult get dns check result
func (d *DNS) GetCheckResult(serverName string) (result *DNSCheckResult) {
	healthy, err := d.Check()
	result = &DNSCheckResult{
		Hostname:      d.Hostname,
		Addr:          d.Server,
		Name:          serverName,
		Healthy:       healthy,
		Err:           err,
		TimeConsuming: d.TimeConsuming,
	}
	if len(d.IPAddr) != 0 {
		ipAddr := make([]string, len(d.IPAddr))
		for index, ip := range d.IPAddr {
			ipAddr[index] = ip.String()
		}
		result.IPAddr = ipAddr
	}
	return
}

// CheckDNSServers check all dns servers
func CheckDNSServers(hostname string, servers []*DNSServer) {
	chs := make(chan bool, 5)
	wg := new(sync.WaitGroup)
	for _, item := range servers {
		go func(server *DNSServer) {
			wg.Add(1)
			chs <- true
			dns := DNS{
				Hostname: hostname,
				Server:   server.Addr,
			}
			result := dns.GetCheckResult(server.Name)
			fmt.Println(result)
			<-chs
			wg.Done()
		}(item)
	}
	wg.Wait()
}
