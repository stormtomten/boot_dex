package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mux      *sync.Mutex
	interval time.Duration
}
type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries:  make(map[string]cacheEntry),
		mux:      &sync.Mutex{},
		interval: interval,
	}
	go newCache.reapLoop()
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.entries[key] = cacheEntry{val: val, createdAt: time.Now()}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, exists := c.entries[key]
	return entry.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}
