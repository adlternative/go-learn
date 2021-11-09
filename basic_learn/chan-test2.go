package main

import (
	"fmt"
	"time"
)

func main() {
	// in := make(chan int)
	// out := make(chan int)
	var sendOut chan int
	v := 1
	sendOut = nil
	for {
		select {
		case sendOut <- v:
			fmt.Printf("send")
		default:
			fmt.Printf("no!")
			time.Sleep(1 * time.Second)
		}
	}

}
