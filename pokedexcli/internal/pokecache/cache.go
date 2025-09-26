package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]CacheEntry
	mu       *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]CacheEntry),
    mu: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
    c.mu.Lock()

    for key, entry := range c.entries {
      elapsedTime := time.Now().Sub(entry.createdAt)
      if elapsedTime > interval {
        delete(c.entries, key)
      }
    }

    c.mu.Unlock()
	}
}
