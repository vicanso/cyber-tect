package util

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	utcTimeStringLen := 20
	now := time.Now()
	if Now().Unix()-now.Unix() > 1 {
		t.Fatalf("get time now fail")
	}

	if len(NowString()) < utcTimeStringLen {
		t.Fatalf("get time now string fail")
	}

	if UTCNow().Unix()-now.Unix() > 1 {
		t.Fatalf("get utc now fail")
	}

	if len(UTCNowString()) != utcTimeStringLen {
		t.Fatalf("get utc now string fail")
	}

	_, err := ParseTime(NowString())
	if err != nil {
		t.Fatalf("parse time fail, %v", err)
	}

	if len(FormatTime(now)) < utcTimeStringLen {
		t.Fatalf("format time fail")
	}
}
