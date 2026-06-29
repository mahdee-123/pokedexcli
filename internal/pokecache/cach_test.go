package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "pokemon"
	expected := []byte("pikachu")

	cache.Add(key, expected)

	val, ok := cache.Get(key)

	if !ok {
		t.Errorf("expected to find key %s in cache", key)
		return
	}

	if string(val) != string(expected) {
		t.Errorf("expected %s, got %s", expected, val)
	}
}

