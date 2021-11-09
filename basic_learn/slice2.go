package main

import "fmt"

func main() {
	s := []int(nil)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	s = []int{}
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	s = nil
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	s = s[:0]
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	s = s[len(s):]
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
