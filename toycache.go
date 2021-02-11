package toycache

import (
	"container/list"
)

type ToyCache interface {
	Put(string, interface{})
	Get(string) (interface{}, bool)
}

const unlimitedSize = -1

type toyCache struct {
	content      map[string]interface{}
	oldestWrites *list.List
	maxSize      int
}

type writeExpire struct {
	key string
}

func MaxSize(maxSize int) func(*toyCache) {
	return func(configured *toyCache) {
		configured.maxSize = maxSize
	}
}

func New(options ...func(subject *toyCache)) ToyCache {
	created := &toyCache{
		content:      make(map[string]interface{}),
		oldestWrites: list.New(),
		maxSize:      unlimitedSize,
	}

	for _, o := range options {
		o(created)
	}

	return created
}

func (c *toyCache) Put(key string, value interface{}) {
	if len(c.content) == c.maxSize {
		c.removeOldest()
	}

	c.oldestWrites.PushBack(writeExpire{key})
	c.content[key] = value
}

func (c *toyCache) Get(key string) (interface{}, bool) {
	val, ok := c.content[key]
	return val, ok
}

func (c *toyCache) removeOldest() {
	oldest := c.oldestWrites.Front()
	delete(c.content, oldest.Value.(writeExpire).key)
	c.oldestWrites.Remove(oldest)
}
