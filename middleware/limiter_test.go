package middleware

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/vicanso/cod"
	concurrentLimiter "github.com/vicanso/cod-concurrent-limiter"
	"github.com/vicanso/hes"
)

func TestCreateConcurrentLimitLock(t *testing.T) {
	fn := createConcurrentLimitLock("test-create-concurrent-limit-", time.Second, true)
	c := cod.NewContext(nil, nil)
	key := "abcd"
	success, done, err := fn(key, c)
	if err != nil || !success {
		t.Fatalf("concurrent limit fail, %v", err)
	}
	success, _, err = fn(key, c)
	if success || err != nil {
		t.Fatalf("the second time should return error")
	}
	done()
	success, _, err = fn(key, c)
	if err != nil || !success {
		t.Fatalf("after call done function, concurrent limit should pass, err:%v", err)
	}
}

func TestNewLimiter(t *testing.T) {
	fn := NewLimiter()
	c := cod.NewContext(nil, nil)
	c.Next = func() error {
		return nil
	}
	err := fn(c)
	if err != nil {
		t.Fatalf("new limiter middleware fail, %v", err)
	}
}

func TestNewConcurrentLimit(t *testing.T) {
	fn := NewConcurrentLimit([]string{
		"q:type",
	}, time.Second, "test-limit-")
	req := httptest.NewRequest("GET", "/users/me?type=1", nil)
	c := cod.NewContext(nil, req)
	c.Next = func() error {
		return nil
	}
	err := fn(c)
	if err != nil {
		t.Fatalf("concurrent limit fail, %v", err)
	}
	err = fn(c)
	he, ok := err.(*hes.Error)
	if !ok || he.Category != concurrentLimiter.ErrCategory {
		t.Fatalf("should return too frequently error")
	}
}

func TestNewIPLimit(t *testing.T) {
	fn := NewIPLimit(1, time.Second, "test-ip-limit-")
	req := httptest.NewRequest("GET", "/users/me", nil)
	c := cod.NewContext(nil, req)
	c.Next = func() error {
		return nil
	}
	err := fn(c)
	if err != nil {
		t.Fatalf("ip limit middleware fail, %v", err)
	}

	err = fn(c)
	if err != errTooFrequently {
		t.Fatalf("should return too frequently error")
	}
}
