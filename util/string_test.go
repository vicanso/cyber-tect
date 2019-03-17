package util

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	if len(RandomString(8)) != 8 {
		t.Fatalf("get random string fail")
	}
}

func TestGenUlid(t *testing.T) {
	if len(GenUlid()) != 26 {
		t.Fatalf("gen ulid fail")
	}
}

func TestSha256(t *testing.T) {
	if Sha256("abcd") != "iNQmb9TmM40TuEX88olXnSCciXgjuSF9o+Fhk28DFYk=" {
		t.Fatalf("sha 256 fail")
	}
}

func TestContainsString(t *testing.T) {
	arr := []string{
		"a",
		"b",
	}
	if !ContainsString(arr, "b") ||
		ContainsString(arr, "c") {
		t.Fatalf("contain string fail")
	}
}
