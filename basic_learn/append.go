package main

import (
	"fmt"
)

func main() {
	var s []int
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 2)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, arr...)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	fmt.Println(s)
}
