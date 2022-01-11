package eviction_policy

type eviction_policy interface {
	KeyUsed(key interface{}) error
	EvictKey(key interface{}) (interface{}, error)
}
