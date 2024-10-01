package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	tick *time.Ticker
	dur time.Duration
	mu *sync.Mutex
}

func NewCache(interval int) Cache {
	var entries = make(map[string]cacheEntry)
	duration := time.Duration(interval) * time.Second
	t := time.NewTicker(duration)
	c := Cache{
		cacheEntries: entries,
		tick: t,
		dur: duration,
		mu: &sync.Mutex{},
	}
	go c.reapLoop()
	return c
}

func (c Cache) Add(url string, b []byte) {
	t := time.Now()
	c.cacheEntries[url] = cacheEntry{created_at: t, val: b}
}

func (c Cache) Get(url string) ([]byte, bool) {
	cache, ok := c.cacheEntries[url]
	return cache.val, ok
}

func (c Cache) reapLoop() {
	for {
		<-c.tick.C
		c.mu.Lock()
		for url, ce := range c.cacheEntries {
			if time.Since(ce.created_at) > c.dur {
				delete(c.cacheEntries, url)
			}
		}
		c.mu.Unlock()
	}
}

type cacheEntry struct {
	created_at time.Time
	val []byte
}