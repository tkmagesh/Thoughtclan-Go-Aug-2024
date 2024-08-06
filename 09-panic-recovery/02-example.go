package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		if q, r, err := divideAdapter(100, divisor); err == ErrDivideByZero {
			fmt.Println("[main] do not attempt to divide by zero")
			continue
		} else if err != nil {
			fmt.Println("[main] err :", err)
			break
		} else {
			fmt.Printf("dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			break
		}
	}
}

func divideAdapter(x, y int) (q, r int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic occurred")
			return
		}
		fmt.Println("Done!")
	}()
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	q, r = divide(x, y)
	return
}

// 3rd party API
func divide(x, y int) (q, r int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	q, r = x/y, x%y
	return
}
