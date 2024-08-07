package main

import (
	"fmt"
	"math"
)

// ver 1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ver 2.0
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

/*
func PrintArea(x interface{}) {
	switch obj := x.(type) {
	case Circle:
		fmt.Println("Area :", obj.Area())
	case Rectangle:
		fmt.Println("Area :", obj.Area())
	default:
		fmt.Println("argument does not support Area() method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch obj := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", obj.Area())
	default:
		fmt.Println("argument does not support Area() method")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	r := Rectangle{10, 12}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)

	PrintArea(100)
}
