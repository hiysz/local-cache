package localcache

type ICache interface {
	Set(k string, v interface{}) error
	Get(k string) (interface{}, error)
	Del(k ...string) (bool, error)
	Exist(k string) (bool, error)
	FetchAll(callback func(interface{}))
	FlushAll() (bool, error)
}
