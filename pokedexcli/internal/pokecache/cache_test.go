package pokecache

import (
	"testing"
	"time"
)

func TestCacheReaping(t *testing.T) {
  const baseTime = time.Millisecond * 5
  const waitTime = baseTime + time.Millisecond * 5
  cache := NewCache(baseTime)
  cache.Add("key", []byte("val"))

  if _, exists := cache.Get("key"); !exists {
    t.Error("expected to find key")
    return
  }

  time.Sleep(waitTime)
  
  if _, exists := cache.Get("key"); exists {
    t.Error("cached key was not reaped")
    return
  }
}
