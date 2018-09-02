package umbra

import (
	"net"
	"strconv"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
	// TypeHTTP http checker
	TypeHTTP = "HTTP"
	// TypePing  ping checker
	TypePing = "Ping"
	// TypeTCP tcp checker
	TypeTCP = "TCP"
	// TypeDNS dns checker
	TypeDNS = "DNS"
)

type (
	// Checker checker interface
	Checker interface {
		Check() (bool, map[string]interface{}, error)
		GetDescription() map[string]interface{}
	}
)

// portCheck the port check
func portCheck(network, ip string, port int) (healthy bool, extra map[string]interface{}, err error) {
	started := time.Now()
	addr := ip
	if port != 0 {
		addr = net.JoinHostPort(ip, strconv.Itoa(port))
	}
	conn, err := net.DialTimeout(network, addr, defaultTimeout)
	if err != nil {
		return
	}
	defer conn.Close()
	extra = make(map[string]interface{})
	extra["stats"] = map[string]string{
		"total": time.Since(started).String(),
	}
	healthy = true
	return
}
