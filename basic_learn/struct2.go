package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// var w Wheel
	// w.X = 8
	// w.Y = 8
	// w.Radius = 5
	// w.Spokes = 20
	w := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	fmt.Printf("%#v\n", w)
	fmt.Println(w)
}
