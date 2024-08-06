package calculator

import "fmt"

func init() {
	fmt.Println("calculator[add.go] initialized")
}

func Add(x, y int) {
	// opCount++
	opCount["add"]++
	fmt.Println("add result :", x+y)
}
