package util

import "testing"

func TestEnv(t *testing.T) {
	if IsDevelopment() ||
		IsProduction() ||
		!IsTest() {
		t.Fatalf("check env fail")
	}
}
