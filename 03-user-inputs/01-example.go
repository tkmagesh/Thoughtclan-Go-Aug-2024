package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("&name = ", &name)
		fmt.Println("Enter the user name :")
		fmt.Scanln(&name)
		fmt.Println("User name : ", name)
	*/

	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	fmt.Println("number = ", n)
}
