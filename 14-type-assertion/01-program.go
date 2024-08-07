package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "Deserunt sit proident mollit amet ut sit sunt cupidatat nostrud."
	x = 99.99
	x = true
	x = struct{}{}
	fmt.Println(x)

	// x = 200
	x = "Ipsum dolore eiusmod ea magna dolor eiusmod labore."
	// y := x.(int) * 2 // successful compilation but error during runtime (if x is not an int)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// x = 100
	// x = "Deserunt sit proident mollit amet ut sit sunt cupidatat nostrud."
	// x = 99.99
	// x = true
	x = struct{}{}
	/* type switch */
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is a zero byte struct")
	default:
		fmt.Println("x is of unknown type")
	}

}
