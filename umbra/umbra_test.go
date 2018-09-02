package umbra

import "testing"

func TestPortCheck(t *testing.T) {
	healthy, _, err := portCheck("tcp", baiduIPAddr, 443)
	if err != nil {
		t.Fatalf("port check fail, %v", err)
	}
	if !healthy {
		t.Fatalf("port check is unhealthy")
	}
}
