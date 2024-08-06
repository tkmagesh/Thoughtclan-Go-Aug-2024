package calculator

import "fmt"

/*
var opCount int

func OpCount() int {
	return opCount
}
*/

var opCount map[string]int

func init() {
	fmt.Println("calculator[calc.go] initialized - 1")
	opCount = make(map[string]int)
}

func init() {
	fmt.Println("calculator[calc.go] initialized - 2")
}

func OpCount() map[string]int {
	return opCount
}
