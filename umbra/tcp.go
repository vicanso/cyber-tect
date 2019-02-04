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
)

// Check check the tcp port is healthy
func (t *TCP) Check() (healthy bool, err error) {
	startedAt := time.Now()
	healthy, err = portCheck("tcp", t.IP, t.Port)
	t.timeConsuming = time.Since(startedAt)
	return
}

// GetDescription get the description of tcp checker
func (t *TCP) GetDescription() (description map[string]interface{}) {
	description = make(map[string]interface{})
	description["type"] = TypeTCP
	description["ip"] = t.IP
	description["port"] = t.Port
	description["timeConsuming"] = t.timeConsuming.String()
	return
}
