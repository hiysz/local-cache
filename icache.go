package localcache

type ICache interface {
	Put(k string, v interface{}) (bool, error)
	Get(k string) (interface{}, error)
	Del(k ...string) (bool, error)
	Exist(k string) (bool, error)
	FetchAll(callback func(interface{}))
	FlushAll() (bool, error)
}
