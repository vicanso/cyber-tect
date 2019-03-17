package umbra

import (
	"testing"
)

func TestHTTPCheck(t *testing.T) {
	http := HTTP{
		URL: "https://www.baidu.com/",
	}
	result := http.GetCheckResult()
	if !result.Healthy {
		t.Fatalf("http check fail, %v", result.Err)
	}
}
