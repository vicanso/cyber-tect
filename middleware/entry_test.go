package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/vicanso/cod"
)

func TestNewEntry(t *testing.T) {
	fn := NewEntry()
	req := httptest.NewRequest("GET", "/users/me", nil)
	res := httptest.NewRecorder()
	c := cod.NewContext(res, req)
	c.Next = func() error {
		return nil
	}
	err := fn(c)
	if err != nil {
		t.Fatalf("new entry middleware fail, %v", err)
	}
	if c.GetHeader(xResponseID) != c.ID {
		t.Fatalf("set response id fail")
	}
}
