package global

import (
	"testing"

	"github.com/vicanso/cyber-tect/util"
)

func TestConnectingCount(t *testing.T) {
	var count uint32 = 100
	SaveConnectingCount(count)
	if GetConnectingCount() != count {
		t.Fatalf("set/get connecting count fail")
	}
}

func TestSyncMap(t *testing.T) {
	key := util.RandomString(8)
	value := 1
	Store(key, value)
	v, ok := Load(key)
	if !ok || v.(int) != value {
		t.Fatalf("load value fail")
	}
	_, loaded := LoadOrStore(key, 2)
	if !loaded {
		t.Fatalf("load or store fail")
	}
}

func TestLruCache(t *testing.T) {
	key := util.RandomString(8)
	value := 1
	Add(key, value)
	v, found := Get(key)
	if !found || v.(int) != value {
		t.Fatalf("lru cache add/get fail")
	}
	Remove(key)
	_, found = Get(key)
	if found {
		t.Fatalf("lru cache remove fail")
	}
}
