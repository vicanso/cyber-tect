package umbra

import (
	"testing"
)

func TestTCPCheck(t *testing.T) {
	tc := TCP{
		IP:   "8.8.8.8",
		Port: 53,
	}
	result := tc.GetCheckResult()
	if !result.Healthy {
		t.Fatalf("tcp check fail, %v", result.Err)
	}
}
