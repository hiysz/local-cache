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
