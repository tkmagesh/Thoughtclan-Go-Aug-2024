package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int

	//address from a value
	noPtr = &no

	// dereferencing - value from an address
	var x int
	x = *noPtr
	fmt.Println(x)

	// in other words
	fmt.Println(no == *(&no))
}
