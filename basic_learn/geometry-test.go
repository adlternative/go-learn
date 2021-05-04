package main

import (
	"abc/geometry"
	"fmt"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{3, 4}
	fmt.Println(geometry.Distance(q, p))
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	p.ScaleBy(12)
	// p.ScaleBy2(12)
	fmt.Println(p)

}
