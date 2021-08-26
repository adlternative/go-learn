package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}
	for _, v := range ints {
		fmt.Println(v)
		ints = ints[1:]
	}

}
