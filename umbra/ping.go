package umbra

import "time"

type (
	// Ping ping check
	Ping struct {
		IP   string
		Type string
		// ping time consuming
		timeConsuming time.Duration
	}
)

// Check ping check
func (p *Ping) Check() (healthy bool, err error) {
	network := "icmp"
	if p.Type == "" {
		network = "ip4:" + network
	} else {
		network = p.Type + ":" + network
	}
	startedAt := time.Now()
	healthy, err = portCheck(network, p.IP, 0)
	p.timeConsuming = time.Since(startedAt)
	return
}

// GetDescription get the ping description
func (p *Ping) GetDescription() (description map[string]interface{}) {
	description = make(map[string]interface{})
	description["type"] = TypePing
	description["ip"] = p.IP
	description["timeConsuming"] = p.timeConsuming.String()
	return
}
