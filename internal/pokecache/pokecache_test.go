package pokecache

import (
	"testing"
	"time"
)

func TestAddGetToCache(t *testing.T) {

	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key",
			inputVal: []byte("value"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("%s not found", c.inputKey)
			continue
		}
		if string(actual) != string(c.inputVal) {
			t.Errorf("%s does not match %s", c.inputVal, actual)
		}
	}

}

func TestReap(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	keyOne := "keyOne"
	cache.Add(keyOne, []byte("valueOne"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("keyOne should have been reaped")
	}
}

func TestReapFail(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	keyOne := "keyOne"
	cache.Add(keyOne, []byte("valueOne"))

	time.Sleep(time.Millisecond)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("keyOne shouldn't have been reaped")
	}
}
