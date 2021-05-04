package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	// tick := time.Tick(1 * time.Second)
	tick := time.NewTicker(1 * time.Second)
	for countdown := 100; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick.C:
			fmt.Println("!")
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		// default:
		// 	//
		// 	fmt.Println("?")
		}
	}
	fmt.Println("launch!")
	tick.Stop() // cause the ticker's goroutine to terminate

}
