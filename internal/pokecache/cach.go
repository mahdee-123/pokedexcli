package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache(interval time.Duration) Cache {
	c := Cache {
		cache: make(map[string]cacheEntry),
	}

	go (&c).reapLoop(interval)
	return c

}

// Create a cache.Add() method that adds a new entry to the cache. It should take a key (a string) and a val (a []byte).


func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()


	for range ticker.C {
		c.mu.Lock()
		for key,entry := range c.cache {
			if time.Since(entry.createdAt) > interval {	
				delete(c.cache , key)
				
			}
		}
		c.mu.Unlock()  
	}
} 

func (c *Cache) Add(key string , value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt : time.Now(), 
		val : value,
	} 
}

func (c *Cache) Get(key string) ([]byte , bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cache,ok := c.cache[key]; ok {
		return cache.val, true
	}
	return nil ,false
} 


