package hashMapStorage

import (
	"fmt"
	"testing"
)

func Test_SetAndGetValueUsingIntKey(t *testing.T) {
	s := NewHashMapStorage(10)
	key := 1
	value := "test"
	s.Set(key, value)
	v, err := s.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	v = v.(string)
	if v != value {
		t.Fatalf(fmt.Sprintf("values are not equal %s != %s", value, v))
	}
}

func Test_SetAndGetValueUsingStringKey(t *testing.T) {
	s := NewHashMapStorage(10)
	key := "1"
	value := "test"
	s.Set(key, value)
	v, err := s.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	v = v.(string)
	if v != value {
		t.Fatalf(fmt.Sprintf("values are not equal %s != %s", value, v))
	}
}

func Test_AccessInvalidKey(t *testing.T) {
	s := NewHashMapStorage(10)
	_, err := s.Get("a")

	if err == nil {
		t.Fatalf("should give error when key not found")
	}
}

func Test_RemoveKey(t *testing.T) {
	s := NewHashMapStorage(10)
	key := 1
	s.Set(key, "")
	s.Remove(key)

	_, err := s.Get("a")
	if err == nil {
		t.Fatalf("key present even after removing")
	}
}
