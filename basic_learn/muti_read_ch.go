package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			v, ok := <-ch
			if !ok {
				fmt.Println("channel is closed")
				return
			}
			fmt.Print(v)
		}()
	}
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(1 * time.Second)
}
