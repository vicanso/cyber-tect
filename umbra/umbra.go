package umbra

import (
	"net"
	"strconv"
	"time"
)

const (
	defaultTimeout = 10 * time.Second
)

// portCheck the port check
func portCheck(network, ip string, port int) (healthy bool, err error) {
	addr := ip
	if port != 0 {
		addr = net.JoinHostPort(ip, strconv.Itoa(port))
	}
	conn, err := net.DialTimeout(network, addr, defaultTimeout)
	if err != nil {
		return
	}
	defer conn.Close()
	healthy = true
	return
}
