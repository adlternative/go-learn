package main

import "fmt"

func f(ch chan bool) {
	c := <-ch
	fmt.Println(c)
}

func main() {
	ch := make(chan bool)
	for i := 0; i != 2; i++ {
		go f(ch)
	}
	ch <- true
	ch <- true

	for {
	}

}
