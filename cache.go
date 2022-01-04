package localcache

import (
	"sync"
)

type Cache struct {
	mu   *sync.RWMutex
	data map[string]interface{}
}

func New() *Cache {
	obj := Cache{
		data: make(map[string]interface{}),
		mu:   new(sync.RWMutex),
	}
	return &obj
}

// Set 设置单条数据
func (c *Cache) Set(k string, v interface{}) error {
	c.mu.Lock()
	c.data[k] = v
	c.mu.Unlock()
	return nil
}

// Get 根据key获取数据
func (c *Cache) Get(k string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if v, ok := c.data[k]; ok {
		return v, nil
	}
	return nil, nil
}

// Del 根据key清除对应数据
func (c *Cache) Del(k ...string) (bool, error) {
	c.mu.Lock()
	for i := 0; i < len(k); i++ {
		delete(c.data, k[i])
	}
	c.mu.Unlock()
	return true, nil
}

// Exist 检查key是否存在
func (c *Cache) Exist(k string) (bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[k]
	return ok, nil
}

// FetchAll 取出所有数据并遍历回调
func (c *Cache) FetchAll(callback func(interface{})) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, v := range c.data {
		callback(v)
	}
}

// FlushAll 清除当前所有数据
func (c *Cache) FlushAll() (bool, error) {
	c.mu.Lock()
	c.data = make(map[string]interface{})
	c.mu.Unlock()
	return true, nil
}
