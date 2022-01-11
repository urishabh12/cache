package lru

import (
	"fmt"
	"testing"
)

func Test_CorrectEvictionSequence(t *testing.T) {
	l := NewLRUEviction()
	l.KeyUsed(1)
	l.KeyUsed(2)
	l.KeyUsed(3)
	l.KeyUsed(1)
	l.KeyUsed(4)

	key, err := l.EvictKey()
	handleErr(t, err)
	assertEqualInt(t, key.(int), 2)
	key, err = l.EvictKey()
	handleErr(t, err)
	assertEqualInt(t, key.(int), 3)
	key, err = l.EvictKey()
	handleErr(t, err)
	assertEqualInt(t, key.(int), 1)
	key, err = l.EvictKey()
	handleErr(t, err)
	assertEqualInt(t, key.(int), 4)
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertEqualInt(t *testing.T, a int, b int) {
	if a != b {
		t.Fatalf(fmt.Sprintf("Values not equal %d != %d", a, b))
	}
}
