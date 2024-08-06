package calculator

import "fmt"

func init() {
	fmt.Println("calculator[subtract.go] initialized")
}

func Subtract(x, y int) {
	// opCount++
	opCount["subtract"]++
	fmt.Println("subtract result :", x-y)
}
