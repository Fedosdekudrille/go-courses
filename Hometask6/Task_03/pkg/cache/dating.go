package cache

import (
	"time"
)

var delElements []string

func (c *Cache) deleteLatestUsedElement() {
	var lastElement *cacheElement
	for { // if element was deleted by others, restart iteration
		c.mutex.RLock()
		for _, el := range c.data {
			if lastElement == nil {
				lastElement = el
			} else if lastElement.usageDate.Before(el.usageDate) {
				lastElement = el
			}
		}
		c.mutex.RUnlock()
		c.mutex.Lock()
		if lastElement != nil && c.data[lastElement.key] != nil {
			delete(c.data, lastElement.key)
			c.mutex.Unlock()
			return
		}
		c.mutex.Unlock()
	}
}

func (c *Cache) deleteOldElements() {
	c.mutex.RLock()
	for _, el := range c.data {
		if time.Since(el.usageDate) > time.Second*10 && c.data[el.key] != nil {
			delElements = append(delElements, el.key)
		}
	}
	c.mutex.RUnlock()
	c.mutex.Lock()
	for _, key := range delElements {
		delete(c.data, key)
	}
	c.mutex.Unlock()
}
