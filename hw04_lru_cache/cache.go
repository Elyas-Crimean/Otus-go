package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

func (c *lruCache) Clear() {
	c.m.Lock()
	defer c.m.Unlock()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	item, found := c.items[key]
	if found {
		c.queue.MoveToFront(item)
		return item.Value.(cachValue).v, true
	}
	return nil, false
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.m.Lock()
	defer c.m.Unlock()
	item, found := c.items[key]
	if found {
		item.Value = cachValue{key, value}
		c.queue.MoveToFront(item)
	} else {
		c.items[key] = c.queue.PushFront(cachValue{key, value})
		if c.queue.Len() > c.capacity {
			backKey := c.queue.Back().Value.(cachValue).key
			delete(c.items, backKey)
			c.queue.Remove(c.queue.Back())
		}
	}
	return found
}

type cachValue struct {
	key Key
	v   interface{}
}

type lruCache struct {
	m        sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
