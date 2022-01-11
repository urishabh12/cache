package storage

type storage interface {
	Set(key interface{}, value interface{}) error
	Get(key interface{}) (interface{}, error)
	Remove(key interface{}) error
}
