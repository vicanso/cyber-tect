package global

import "testing"

func TestChangeApplicationStatus(t *testing.T) {
	if IsApplicationRunning() {
		defer StartApplication()
	}
	StartApplication()
	if !IsApplicationRunning() {
		t.Fatalf("application should be running")
	}
	PauseApplication()
	if IsApplicationRunning() {
		t.Fatalf("application should be paused")
	}
}
