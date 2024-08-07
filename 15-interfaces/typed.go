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

type AreaFinder interface{ Area() float64 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// ver 3.0
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface {Perimeter() float64 }
}

/*
func PrintShapeStats2(x interface {
	Area() float64
	Perimeter() float64
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface {Perimeter() float64 }
}
*/

// ver 4.0
// fmt.Stringer interface implementation
func (c Circle) String() string {
	return fmt.Sprintf("[Circle] Radius : %v", c.Radius)
}

// fmt.Stringer interface implementation
func (r Rectangle) String() string {
	return fmt.Sprintf("[Rectangle] Length : %v, Breadth : %v", r.Length, r.Breadth)
}

func main() {
	c := Circle{12}
	fmt.Println(c)
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{10, 12}
	fmt.Println(r)
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)

}
