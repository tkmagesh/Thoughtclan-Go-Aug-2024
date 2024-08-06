package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		fmt.Println("[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("app panicked, err :", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Println(q, r)
}

// using named result
func divide(x, y int) (q, r int) {
	defer func() {
		fmt.Println("[divide] - deferred")
	}()
	fmt.Println("[divide] calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	q = x / y
	fmt.Println("[divide] calculating remainder")
	r = x % y
	return
}
