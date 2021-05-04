package main

import (
	"fmt"
	"log"
)

func main() {
	if err := f(3); err != nil {
		log.Println(err)
	}
}
func f(x int) (err error) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	return f(x - 1)

}
