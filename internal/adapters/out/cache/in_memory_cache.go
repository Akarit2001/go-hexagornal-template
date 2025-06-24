package cache

import (
	"encoding/json"
	"go-hex-temp/internal/ports/output"
	"sync"
	"time"
)

type inMemoryCache struct {
	mu    sync.RWMutex
	store map[string][]byte
}

// Del implements output.Cache.
func (c *inMemoryCache) Del(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
	return nil
}

func NewInMemoryCache() output.Cache {
	return &inMemoryCache{
		store: make(map[string][]byte),
	}
}

func (c *inMemoryCache) Save(key string, data any, _ time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = b
	return nil
}

func (c *inMemoryCache) Load(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.store[key]
	return val, ok
}
