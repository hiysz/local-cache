# LocalCache
轻量级并发安全本地缓存组件。

## Install
```bash
go get github.com/hiysz/localcache
```

## example
```go
package main

import (
	"fmt"
	"github.com/hiysz/localcache"
)

func main() {
	cache := localcache.New()
	err := cache.Set("keyName", 101)
	if err != nil {
		fmt.Println("set error:", err)
		return
	}
	v, err := cache.Get("keyName")
	fmt.Println(v, err)
}

```
