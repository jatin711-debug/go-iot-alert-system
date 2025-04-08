package cache

import (
	"container/list"
	"sync"
	"time"
)

type entry struct {
	key       string
	value     interface{}
	expiresAt time.Time
}

type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	ll       *list.List
	mu       sync.Mutex
	ttl      time.Duration
}

func NewLRUCache(capacity int, ttl time.Duration) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		ll:       list.New(),
		ttl:      ttl,
	}
	go lru.cleanupLoop()
	return lru
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.cache[key]; ok {
		ent := ele.Value.(*entry)
		if time.Now().After(ent.expiresAt) {
			c.ll.Remove(ele)
			delete(c.cache, key)
			return nil, false
		}
		c.ll.MoveToFront(ele)
		return ent.value, true
	}
	return nil, false
}

func (c *LRUCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.cache[key]; ok {
		ent := ele.Value.(*entry)
		ent.value = value
		ent.expiresAt = time.Now().Add(c.ttl)
		c.ll.MoveToFront(ele)
		return
	}

	ent := &entry{
		key:       key,
		value:     value,
		expiresAt: time.Now().Add(c.ttl),
	}
	ele := c.ll.PushFront(ent)
	c.cache[key] = ele

	if c.ll.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.ll.Remove(ele)
		delete(c.cache, key)
	}
}

func (c *LRUCache) removeOldest() {
	ele := c.ll.Back()
	if ele != nil {
		ent := ele.Value.(*entry)
		delete(c.cache, ent.key)
		c.ll.Remove(ele)
	}
}

// background cleanup of expired keys
func (c *LRUCache) cleanupLoop() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.cleanup()
	}
}

func (c *LRUCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for e := c.ll.Back(); e != nil; {
		prev := e.Prev()
		ent := e.Value.(*entry)
		if time.Now().After(ent.expiresAt) {
			c.ll.Remove(e)
			delete(c.cache, ent.key)
		}
		e = prev
	}
}
