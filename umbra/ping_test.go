package umbra

import (
	"fmt"
	"testing"
)

func TestPingCheck(t *testing.T) {
	p := Ping{
		IP: "8.8.8.8",
	}
	result := p.GetCheckResult()
	fmt.Println(result)
}
