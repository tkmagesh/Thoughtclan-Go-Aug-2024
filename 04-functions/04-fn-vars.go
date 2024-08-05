package main

import "fmt"

func main() {
	/* 0 parameters, 0 return values */
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()
	sayHi = func() {
		fmt.Println("Hello!")
	}
	sayHi()

	/* 1 parameter, 0 return values */
	// var greet func(uName string)
	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	/* 2 parameters, 1 return values */
	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))
}
