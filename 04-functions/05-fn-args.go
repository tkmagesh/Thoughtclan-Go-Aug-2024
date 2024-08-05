package main

import "fmt"

func main() {
	exec(f1) // to invoked f1
	exec(f2) // to invoked f2
	var f3 func()
	f3 = func() {
		fmt.Println("f3 invoked")
	}
	exec(f3)
	exec(func() {
		fmt.Println("anon fn invoked")
	})
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
