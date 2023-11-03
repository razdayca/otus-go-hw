package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c lruCache) Set(key Key, value interface{}) bool {
	val, ok := c.items[key]
	if ok {
		c.queue.Remove(val)
		c.queue.PushFront(value)
		c.items[key].Value = value
		return true
	}

	if c.capacity == c.queue.Len() {
		for k, v := range c.items {
			if v.Value == c.queue.Back().Value {
				delete(c.items, k)
			}
		}
		c.queue.Remove(c.queue.Back())
	}

	c.items[key] = &ListItem{
		Value: value,
	}

	c.queue.PushFront(value)

	return false
}

func (c lruCache) Get(key Key) (interface{}, bool) {
	val, ok := c.items[key]
	if ok {
		c.queue.Remove(val)
		c.queue.PushFront(val)
		return val.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = map[Key]*ListItem{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
