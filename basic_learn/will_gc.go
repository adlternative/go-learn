package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type test struct {
	arg int
	ch  chan int
}

func main() {
	var mu sync.Mutex
	var w sync.WaitGroup
	m := make(map[interface{}]interface{})

	w.Add(1)
	go func() {
		defer w.Done()
		t := &test{
			arg: 1,
			ch:  make(chan int),
		}
		mu.Lock()
		m["t1"] = t
		mu.Unlock()
		select {
		case val := <-t.ch:
			fmt.Println(val)
		}
	}()
	w.Add(1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		runtime.GC()
		for {
			mu.Lock()
			for _, v := range m {
				go func(v interface{}) {
					v.(*test).ch <- v.(*test).arg
				}(v)
			}
			mu.Unlock()

		}
	}()
	w.Wait()
}
