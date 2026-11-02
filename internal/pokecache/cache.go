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
	entries map[string]cacheEntry
	m       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: map[string]cacheEntry{},
		m:       sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.m.Lock()
		for k, v := range c.entries {
			if time.Since(v.createdAt) > interval {
				delete(c.entries, k)
			}
		}
		c.m.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	defer c.m.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	entry, ok := c.entries[key]

	return entry.val, ok
}
