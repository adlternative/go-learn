package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}

	for i, v := range ints {
		fmt.Println(v, i)
		ints = ints[1:]
		fmt.Println(ints)
	}
}
