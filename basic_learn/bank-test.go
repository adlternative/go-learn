package main

import (
	_ "basic_learn/bank"
	_ "basic_learn/bank2"
	"basic_learn/bank3"
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			bank3.Deposit(3)
		}()
	}
	time.Sleep(1*time.Second)
	fmt.Println(bank3.Balance())
}
