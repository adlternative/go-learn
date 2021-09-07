package main

import (
	"fmt"
	"time"
)

func dosomething() {
	fmt.Println("做事!")
	time.Sleep(3 * time.Second)
	fmt.Println("做事完了!")
}
func main() {
	t := time.NewTimer(1 * time.Second)
	fmt.Println("故意睡!")
	time.Sleep(3 * time.Second)
	fmt.Println("故意睡醒!")
	for {
		select {
		case <-t.C:
			fmt.Println("超时!")
			t.Reset(1 * time.Second)
			go dosomething()
		}
	}
}
