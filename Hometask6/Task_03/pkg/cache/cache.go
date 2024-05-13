package cache

import (
	"sync"
	"time"
)

type cacheElement struct {
	key       string
	obj       any
	usageDate time.Time
}

type Cache struct {
	data     map[string]*cacheElement
	capacity int
	mutex    *sync.RWMutex
}

func NewCache(cap int) Cache {
	c := Cache{
		data:     make(map[string]*cacheElement),
		capacity: cap,
		mutex:    new(sync.RWMutex),
	}
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			c.deleteOldElements()
		}
	}()
	return c
}

func (c *Cache) Set(key string, obj any) {
	if len(c.data) >= c.capacity {
		go c.deleteLatestUsedElement()
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = &cacheElement{
		key:       key,
		obj:       obj,
		usageDate: time.Now(),
	}
}

func (c *Cache) Get(key string) any {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.data[key] != nil { //double check for the case of deleting the element after the first check
		c.data[key].usageDate = time.Now()
		return c.data[key].obj
	}
	return nil
}

func (c *Cache) Remove(key string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.data[key] != nil {
		delete(c.data, key)
		return true
	}
	return false
}
