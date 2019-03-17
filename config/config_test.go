package config

import (
	"strings"
	"testing"
	"time"
)

func TestGetListen(t *testing.T) {
	if GetListen() != defaultListen {
		t.Fatalf("get listen fail")
	}
}

func TestGetENV(t *testing.T) {
	if GetENV() != "test" {
		t.Fatalf("get env fail")
	}
}

func TestGet(t *testing.T) {
	randomKey := "xx_xx_xx"
	if GetIntDefault("requestLimit", 0) != 1024 {
		t.Fatalf("get int fail")
	}

	if GetIntDefault(randomKey, 1) != 1 {
		t.Fatalf("get int default fail")
	}

	if GetString("app") != "cyber-tect" {
		t.Fatalf("get string fail")
	}

	if GetStringDefault(randomKey, "1") != "1" {
		t.Fatalf("get string default fail")
	}

	if GetDurationDefault(randomKey, time.Second) != time.Second {
		t.Fatalf("get time duration default fail")
	}

	if strings.Join(GetStringSlice("keys"), ",") != "cuttlefish,secret" {
		t.Fatalf("get string slice fail")
	}
}

func TestGetTrackKey(t *testing.T) {
	if GetTrackKey() != defaultTrackKey {
		t.Fatalf("get track key fail")
	}
}

func TestGetSessionConfig(t *testing.T) {
	scf := GetSessionConfig()
	if scf.TTL != defaultSessionTTL ||
		scf.Key != defaultSessionKey ||
		scf.CookiePath != defaultCookiePath {
		t.Fatalf("get session config fail")
	}
}
