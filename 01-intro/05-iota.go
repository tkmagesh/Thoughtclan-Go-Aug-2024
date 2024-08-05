package main

import "fmt"

func main() {
	/*
		const red int = 0
		const green int = 1
		const blue int = 2
	*/
	/*
		const (
			red   int = iota
			green int = iota
			blue  int = iota
		)
	*/

	/* 	const (
	   		red int = iota
	   		green
	   		blue
	   	)
	*/
	/* const (
		red = iota
		green
		blue
	)
	*/

	/*
		const (
			red = iota
			green
			_
			blue
		)
	*/

	const (
		red = (iota * 2) + 1
		green
		blue
	)

	fmt.Printf("red = %d, green = %d and blue = %d\n", red, green, blue)

	const (
		x = 1 << iota
		w
		r
	)
	fmt.Printf("%b, %b, %b\n", r, w, x)
}
