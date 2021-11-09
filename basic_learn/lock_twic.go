package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := sync.Mutex{}
	go func() {
		mu.Lock()
		fmt.Println("?")
		mu.Lock()
		fmt.Println("?")
	}()
	for {
		fmt.Println("hehe")
		time.Sleep(time.Second)
	}
}
