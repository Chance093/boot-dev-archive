package pokecache

import "time"

// Add an entry to the PokeCache
func (c *Cache) Add(key string, val []byte) {
  c.mu.Lock()
  defer c.mu.Unlock()

  c.entries[key] = CacheEntry{
    createdAt: time.Now(),
    val: val,
  }
}

// Get an entry from the PokeCache
// Returns a byte slice containing cached data and boolean based on if entry was found
func (c *Cache) Get(key string) ([]byte, bool) {
  c.mu.Lock()
  defer c.mu.Unlock()

  val, exists := c.entries[key]
  if !exists {
    return nil, false
  }

  return val.val, true
}
