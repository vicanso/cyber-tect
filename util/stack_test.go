package util

import (
	"strings"
	"testing"
)

func TestGetStack(t *testing.T) {
	if !strings.Contains(GetStack(0, 3)[0], "util.GetStack:") {
		t.Fatalf("get stack fail")
	}
}
