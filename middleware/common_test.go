package middleware

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/vicanso/cod"
)

func TestNoQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/me?a=1", nil)
	c := cod.NewContext(nil, req)
	c.Next = func() error {
		return nil
	}
	if NoQuery(c) != errQueryNotAllow {
		t.Fatalf("no query middleware is fail")
	}

	req = httptest.NewRequest("GET", "/users/me", nil)
	c.Request = req
	if NoQuery(c) != nil {
		t.Fatalf("no query should pass")
	}
}

func TestWaitFor(t *testing.T) {
	start := time.Now()
	d := 10 * time.Millisecond
	fn := WaitFor(d)
	c := cod.NewContext(nil, nil)
	c.Next = func() error {
		return nil
	}
	err := fn(c)
	if err != nil {
		t.Fatalf("wait for middleware fail, %v", err)
	}
	use := time.Since(start)
	if use < d {
		t.Fatalf("wait for middleware fail")
	}
}
