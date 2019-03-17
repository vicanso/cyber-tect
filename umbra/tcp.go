package umbra

import "time"

type (
	// TCP tcp port check
	TCP struct {
		IP   string
		Port int
		// tcp connect time consuming
		timeConsuming time.Duration
	}
	// TCPCheckResult tcp check result
	TCPCheckResult struct {
		IP            string        `json:"ip,omitempty"`
		Port          int           `json:"port,omitempty"`
		Err           error         `json:"err,omitempty"`
		Message       string        `json:"message,omitempty"`
		Healthy       bool          `json:"healthy,omitempty"`
		TimeConsuming time.Duration `json:"timeConsuming,omitempty"`
	}
)

// Check check the tcp port is healthy
func (t *TCP) Check() (healthy bool, err error) {
	startedAt := time.Now()
	healthy, err = portCheck("tcp", t.IP, t.Port)
	t.timeConsuming = time.Since(startedAt)
	return
}

// GetCheckResult get tcp check result
func (t *TCP) GetCheckResult() (result TCPCheckResult) {
	healthy, err := t.Check()
	result = TCPCheckResult{
		IP:            t.IP,
		Port:          t.Port,
		Healthy:       healthy,
		Err:           err,
		TimeConsuming: t.timeConsuming,
	}
	if err != nil {
		result.Message = err.Error()
	}
	return
}
