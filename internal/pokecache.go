package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.RWMutex
	table    map[string]cacheEntry
	interval time.Duration
}

func (c *Cache) NewCache(interval time.Duration) {
	c.table = make(map[string]cacheEntry)
	c.interval = interval
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		table:    make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.table[key] = newEntry

	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.RLock()
	entry, ok := c.table[key]
	c.mu.RUnlock()

	if !ok {
		return []byte{}, false
	}

	return entry.val, true

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	defer ticker.Stop()

	for range ticker.C {
		for k, item := range c.table {

			c.mu.Lock()
			if time.Since(item.createdAt) > c.interval {
				delete(c.table, k)
			}
			c.mu.Unlock()
		}
	}

}
