package umbra

type (
	// TCP tcp port check
	TCP struct {
		IP   string
		Port int
	}
)

// Check check the tcp port is healthy
func (t *TCP) Check() (healthy bool, extra map[string]interface{}, err error) {
	return portCheck("tcp", t.IP, t.Port)
}

// GetDescription get the description of tcp checker
func (t *TCP) GetDescription() (description map[string]interface{}) {
	description = make(map[string]interface{})
	description["type"] = TypeTCP
	description["ip"] = t.IP
	description["port"] = t.Port
	return
}
