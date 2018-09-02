package umbra

type (
	// Ping ping check
	Ping struct {
		IP   string
		Type string
	}
)

// Check ping check
func (p *Ping) Check() (healthy bool, extra map[string]interface{}, err error) {
	network := "icmp"
	if p.Type == "" {
		network = "ip4:" + network
	} else {
		network = p.Type + ":" + network
	}
	return portCheck(network, p.IP, 0)
}

// GetDescription get the ping description
func (p *Ping) GetDescription() (description map[string]interface{}) {
	description = make(map[string]interface{})
	description["type"] = TypePing
	description["ip"] = p.IP
	if p.Type != "" {
		description["type"] = p.Type
	}
	return
}
