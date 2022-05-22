package cache

import "time"

type DeadlineValue struct {
	value    string
	expired  bool
	deadline time.Time
}

func NewNoDeadlineValue(val string) DeadlineValue {
	return DeadlineValue{value: val, expired: false, deadline: time.Now()}
}

func NewDeadLineValue(val string, expiredTime time.Time) DeadlineValue {
	return DeadlineValue{value: val, expired: true, deadline: expiredTime}
}

type Cache struct {
	dict  map[string]DeadlineValue
	timer *time.Timer
}

func NewCache() Cache {
	var cache Cache
	cache.dict = make(map[string]DeadlineValue)
	cache.timer = time.AfterFunc(time.Second, func() {
		expiryTime := time.Now()
		for k, v := range cache.dict {
			if v.expired && v.deadline.Before(expiryTime) {
				delete(cache.dict, k)
			}
		}
	})

	return cache
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.dict[key]
	return value.value, ok
}

func (c Cache) Put(key, value string) {
	c.dict[key] = NewNoDeadlineValue(value)
}

func (c Cache) Keys() []string {
	var keys []string
	for k, _ := range c.dict {
		keys = append(keys, k)
	}

	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.dict[key] = NewDeadLineValue(value, deadline)
}
