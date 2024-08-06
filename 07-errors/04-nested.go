package main

import (
	"errors"
	"fmt"
)

var ErrF1 error = errors.New("f1 error")
var ErrF2 error = errors.New("f2 error")

func main() {
	if e := f1(); e != nil {
		if errors.Is(e, ErrF1) {
			fmt.Println("error occurred in f1()")
		}
		if errors.Is(e, ErrF2) {
			fmt.Println("error occurred in f2()")
		}
		return
	}
	fmt.Println("All good!")
}

func f1() error {
	e := f2()
	if e == nil {
		return nil
	}
	// return e
	return fmt.Errorf("%w %w", e, ErrF1)
}

func f2() error {
	return ErrF2
}
