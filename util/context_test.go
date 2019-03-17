package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/config"
)

func TestGetTrackID(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/me", nil)
	req.AddCookie(&http.Cookie{
		Name:  config.GetTrackKey(),
		Value: "abcd",
	})
	c := cod.NewContext(nil, req)
	if GetTrackID(c) != "abcd" {
		t.Fatalf("get track id fail")
	}
}
