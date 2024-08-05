/*
anonymous functions
  - functions with no name
  - have to be immediately invoked
*/
package main

import "fmt"

func main() {
	/* 0 parameters, 0 return values */
	func() {
		fmt.Println("Hi!")
	}()

	/* 1 parameter, 0 return values */
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")
}
