package config

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]cacheItem
	ttl  time.Duration
}

type cacheItem struct {
	value     interface{}
	timestamp time.Time
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		data: make(map[string]cacheItem),
		ttl:  ttl,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheItem{
		value:     value,
		timestamp: time.Now(),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[key]
	if !exists || time.Since(item.timestamp) > c.ttl {
		return nil, false
	}

	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

func (c *Cache) CleanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, item := range c.data {
		if time.Since(item.timestamp) > c.ttl {
			delete(c.data, key)
		}
	}
}
