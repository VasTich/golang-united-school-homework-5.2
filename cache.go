package cache

import "time"

type Cache struct {
	dict map[string]string
}

func NewCache() Cache {
	return Cache{dict: make(map[string]string)}
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.dict[key]
	return value, ok
}

func (c Cache) Put(key, value string) {
	c.dict[key] = value
}

func (c Cache) Keys() []string {
	var keys []string
	for k, _ := range c.dict {
		keys = append(keys, k)
	}
	
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
}
