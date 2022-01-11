package storage

type Storage interface {
	Set(key interface{}, value interface{}) error
	Get(key interface{}) (interface{}, error)
	Remove(key interface{}) error
}
