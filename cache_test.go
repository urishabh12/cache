package cache

import (
	"fmt"
	"testing"

	l "github.com/urishabh12/cache/eviction_policy/lru"
	h "github.com/urishabh12/cache/storage/hashMapStorage"
)

func NewLRUHashMapCache(size int) Cache {
	return Cache{
		storage:        h.NewHashMapStorage(size),
		evictionPolicy: l.NewLRUEviction(),
	}
}

func Test_SetAndGet(t *testing.T) {
	c := NewLRUHashMapCache(10)
	c.Set(1, "One")
	c.Set(3, "Three")

	val, err := c.Get(1)
	handleErr(t, err)
	assertEqualString(t, val.(string), "One")
	val, err = c.Get(3)
	handleErr(t, err)
	assertEqualString(t, val.(string), "Three")
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertEqualString(t *testing.T, a string, b string) {
	if a != b {
		t.Fatalf(fmt.Sprintf("Values not equal %s != %s", a, b))
	}
}
