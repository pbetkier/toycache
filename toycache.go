package toycache

import (
	"container/list"
	"time"
)

type ToyCache interface {
	Put(string, interface{})
	Get(string) (interface{}, bool)
}

type Clock interface {
	Now() time.Time
}

const unlimitedSize = -1
const noWriteTTL = 0

type toyCache struct {
	content      map[string]interface{}
	oldestWrites *list.List
	clock        Clock
	maxSize      int
	writeTTL     time.Duration
}

type realClock struct{}

func (rc *realClock) Now() time.Time {
	return time.Now()
}

type writeExpire struct {
	key        string
	expiration time.Time
}

func MaxSize(maxSize int) func(*toyCache) {
	return func(configured *toyCache) {
		configured.maxSize = maxSize
	}
}

func WriteTTL(ttl time.Duration) func(*toyCache) {
	return func(configured *toyCache) {
		configured.writeTTL = ttl
	}
}

func WithClock(clock Clock) func(*toyCache) {
	return func(configured *toyCache) {
		configured.clock = clock
	}
}

func New(options ...func(subject *toyCache)) ToyCache {
	created := &toyCache{
		content:      make(map[string]interface{}),
		oldestWrites: list.New(),
		clock:        &realClock{},
		maxSize:      unlimitedSize,
		writeTTL:     noWriteTTL,
	}

	for _, o := range options {
		o(created)
	}

	return created
}

func (c *toyCache) Put(key string, value interface{}) {
	c.removeExpired()

	if len(c.content) == c.maxSize {
		c.removeOldest()
	}

	c.oldestWrites.PushBack(writeExpire{key, c.clock.Now()})
	c.content[key] = value
}

func (c *toyCache) Get(key string) (interface{}, bool) {
	c.removeExpired()

	val, ok := c.content[key]
	return val, ok
}

func (c *toyCache) removeOldest() {
	oldest := c.oldestWrites.Front()
	delete(c.content, oldest.Value.(writeExpire).key)
	c.oldestWrites.Remove(oldest)
}

func (c *toyCache) removeExpired() {
	if c.writeTTL == noWriteTTL {
		return
	}

	for e := c.oldestWrites.Front(); e != nil; e = c.oldestWrites.Front() {
		we := e.Value.(writeExpire)
		if we.expiration.Before(c.clock.Now()) {
			delete(c.content, we.key)
			c.oldestWrites.Remove(e)
		} else {
			return
		}
	}
}
