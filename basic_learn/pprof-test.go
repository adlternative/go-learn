package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

var c = make(chan int)

func f(i int) {
	g(i + 1)
}

func g(i int) {
	h(i + 1)
}

func h(i int) {
	c <- 1
	f(i + 1)
}

func main() {
	for i := 0; i < 100; i++ {
		go f(0x10 * i)
	}
	log.Println(http.ListenAndServe(":8080", nil))
}
