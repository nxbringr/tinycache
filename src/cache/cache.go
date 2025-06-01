package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	Value     interface{}
	ExpiresAT time.Time
}

type Cache struct {
	mu    sync.RWMutex
	store map[string]CacheEntry
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]CacheEntry),
	}
}

func (c *Cache) WriteEntry(key string, value interface{}) {
	c.mu.Lock()
	c.store[key] = CacheEntry{
		Value:     value,
		ExpiresAT: time.Now().Add(5 * time.Minute),
	}
	c.mu.Unlock()
}

func (c *Cache) ReadEntry(key string) (interface{}, bool) {
	c.mu.RLock()
	entry, ok := c.store[key]
	c.mu.RUnlock()

	if !ok || time.Now().After(entry.ExpiresAT) {
		return nil, false
	}
	return entry.Value, ok
}

func (c *Cache) DeleteEntry(key string) {
	c.mu.Lock()
	delete(c.store, key)
	c.mu.Unlock()

}

func (c *Cache) ReadAllEntries() map[string]CacheEntry {
	c.mu.RLock()

	result := make(map[string]CacheEntry)
	for k, v := range c.store {
		result[k] = v
	}
	return result
}
