package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	divisor := 0
	if q, r, err := divide(100, divisor); err == ErrDivideByZero {
		fmt.Println("[main] do not attempt to divide by zero")
	} else if err != nil {
		fmt.Println("[main] err :", err)
	} else {
		fmt.Printf("dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)

	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("divide by zero error")
	}
	return x / y, x % y, nil
}
*/

// using named result
func divide(x, y int) (q, r int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	q, r = x/y, x%y
	return
}
