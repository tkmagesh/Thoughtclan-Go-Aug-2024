package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var fn func()
	fn = getFn()
	fn()
}

func getFn() func() {
	if no := rand.Intn(20); no%2 == 0 {
		return f1
	}
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
