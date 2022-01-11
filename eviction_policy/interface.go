package eviction_policy

type EvictionPolicy interface {
	KeyUsed(key interface{}) error
	EvictKey() (interface{}, error)
}
