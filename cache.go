package cache

import (
	er "github.com/urishabh12/cache/errors"
	e "github.com/urishabh12/cache/eviction_policy"
	s "github.com/urishabh12/cache/storage"
)

type Cache struct {
	storage        s.Storage
	evictionPolicy e.EvictionPolicy
}

func NewCache(st s.Storage, ep e.EvictionPolicy) Cache {
	return Cache{
		storage:        st,
		evictionPolicy: ep,
	}
}

func (c *Cache) Set(key interface{}, value interface{}) error {
	err := c.storage.Set(key, value)
	c.evictionPolicy.KeyUsed(key)
	if er.IsStorageFullError(err) {
		evictedKey, err := c.evictionPolicy.EvictKey()
		if err != nil {
			return err
		}

		c.storage.Remove(evictedKey)
		c.Set(key, value)
	}

	return err
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	value, err := c.storage.Get(key)
	if err != nil {
		return nil, err
	}
	c.evictionPolicy.KeyUsed(key)

	return value, nil
}
