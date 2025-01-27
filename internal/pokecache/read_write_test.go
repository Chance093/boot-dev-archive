package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestGetAddCache(t *testing.T) {
  cases := []struct {
    key string
    val []byte
  }{
    {
      key: "https://hello.com/wow",
      val: []byte("here is some data"),
    },
    {
      key: "https://wow.com/hello",
      val: []byte("here is some other data"),
    },
  }

  const interval = time.Second * 5

  for i, c := range cases {
    t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
      cache := NewCache(interval)
      cache.Add(c.key, c.val)
      val, exists := cache.Get(c.key)
      if !exists {
        t.Errorf("Key '%v' not found in cache", c.key)
        return
      }

      if string(val) != string(c.val) {
        t.Errorf("cached val does not match case val")
        return
      }
    })
  }
}
