package main

import (
	"fmt"
)

func f() *int {
	v := 1
	return &v
}

func main() {
	var x, y int

	ret := f()
	fmt.Println(ret, *ret)
	ret = f()
	fmt.Println(ret, *ret)
	fmt.Println(&x == &x, &x == &y, &x == nil, x) // "true false false"
}
