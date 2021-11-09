package main

import "fmt"

func main() {
	s := make([]int, 0)
	s = append(s, 1)
	s = s[:0]
	if s == nil {
		fmt.Println("nil")
	}
}
