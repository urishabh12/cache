package hashMapStorage

import (
	e "github.com/urishabh12/cache/errors"
)

type HashMapStorage struct {
	hashMap map[interface{}]interface{}
	size    int
	maxSize int
}

func NewHashMapStorage(size int) *HashMapStorage {
	return &HashMapStorage{
		hashMap: make(map[interface{}]interface{}),
		size:    0,
		maxSize: size,
	}
}

func (h *HashMapStorage) Set(key interface{}, value interface{}) error {
	if h.size >= h.maxSize {
		return &e.StorageFullError{}
	}

	h.hashMap[key] = value
	h.size++
	return nil
}

func (h *HashMapStorage) Get(key interface{}) (interface{}, error) {
	resp, ok := h.hashMap[key]
	if !ok {
		return nil, &e.NotFoundError{}
	}

	return resp, nil
}

func (h *HashMapStorage) Remove(key interface{}) error {
	delete(h.hashMap, key)
	h.size--
	return nil
}
