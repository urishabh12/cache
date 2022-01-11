package hashMapStorage

import "errors"

type hashMapStorage struct {
	hashMap map[interface{}]interface{}
}

func NewHashMapStorage() *hashMapStorage {
	return &hashMapStorage{
		hashMap: make(map[interface{}]interface{}),
	}
}

func (h *hashMapStorage) Set(key interface{}, value interface{}) error {
	h.hashMap[key] = value
	return nil
}

func (h *hashMapStorage) Get(key interface{}) (interface{}, error) {
	resp, ok := h.hashMap[key]
	if !ok {
		return nil, errors.New("key dosen't exists")
	}

	return resp, nil
}

func (h *hashMapStorage) Remove(key interface{}) error {
	delete(h.hashMap, key)
	return nil
}
