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
	// PingCheckResult ping check result
	PingCheckResult struct {
		IP            string        `json:"ip,omitempty"`
		Type          string        `json:"type,omitempty"`
		Err           error         `json:"err,omitempty"`
		Message       string        `json:"message,omitempty"`
		Healthy       bool          `json:"healthy,omitempty"`
		TimeConsuming time.Duration `json:"timeConsuming,omitempty"`
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

// GetCheckResult get ping check result
func (p *Ping) GetCheckResult() (result PingCheckResult) {
	healthy, err := p.Check()
	result = PingCheckResult{
		IP:            p.IP,
		Type:          p.Type,
		Healthy:       healthy,
		Err:           err,
		TimeConsuming: p.timeConsuming,
	}
	if err != nil {
		result.Message = err.Error()
	}
	return
}
