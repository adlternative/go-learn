package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func (cache Ggg) Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

type Ggg struct {
	sync.Mutex
	mapping map[string]string
}

func main() {
	cache := Ggg{
		mapping: make(map[string]string),
	}
	cache.mapping["asd"] = "qqq"
	fmt.Println(cache.Lookup("asd"))
	fmt.Println(cache.Lookup("asd"))
	// Lookup("asdasd")
}
