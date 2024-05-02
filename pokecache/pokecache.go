package pokecache

import (
	"sync"
	"time"
)

type Cache_entry struct {
	Created_at time.Time
	Val        []byte
}

type Cache struct {
	Interval   time.Time
	Cache_data map[string]Cache_entry
	Mutex      *sync.Mutex
}

func New_cache(interval time.Duration) Cache {
    c := Cache{
		Interval:   time.Now().UTC(),
		Cache_data: make(map[string]Cache_entry),
        Mutex: &sync.Mutex{},
	}
    go c.Reap_loop(interval)
    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.Mutex.Lock()
    defer c.Mutex.Unlock()
    c.Cache_data[key] = Cache_entry{
        Created_at: time.Now(),
        Val: val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.Mutex.Lock()
    defer c.Mutex.Unlock()
    data, found := c.Cache_data[key]
    if !found {
        return nil, false
    }
    return data.Val, true
}

func (c *Cache) Reap_loop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.Mutex.Lock()
        for key, val := range c.Cache_data {
            if val.Created_at.Before(time.Now().UTC().Add(-interval)) {
                delete(c.Cache_data, key)

            }
        }
        c.Mutex.Unlock()
    }
}
