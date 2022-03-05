package cache

import (
	"log"
	"sync"
)

type inMemoryCache struct {
	cache map[string][]byte
	mutex sync.RWMutex
	Stat
}

// set memory cache
func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if tmp, exist := c.cache[k]; exist {
		c.del(k, tmp)
	}

	c.cache[k] = v
	c.add(k, v)
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.cache[k], nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if value, exist := c.cache[k]; exist {
		delete(c.cache, k)
		c.del(k, value)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}

func New() Cache {
	c := newInMemoryCache()
	if c == nil {
		panic("unknow cache")
	}
	log.Println("cache serve is ready")
	return c
}
