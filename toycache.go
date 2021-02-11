package toycache

type ToyCache interface {
	Put(string, interface{})
	Get(string) (interface{}, bool)
}

type toyCache struct {
	content map[string]interface{}
}

func New() ToyCache {
	return &toyCache{make(map[string]interface{})}
}

func (c *toyCache) Put(key string, value interface{}) {
	c.content[key] = value
}

func (c *toyCache) Get(key string) (interface{}, bool) {
	val, ok := c.content[key]
	return val, ok
}
