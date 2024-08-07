package main

import "fmt"

func main() {
	// fmt.Println(sum())
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, "20"))
	fmt.Println(sum(10, "20", "abc"))
	fmt.Println(sum(10, 20, 30, 40, "50"))
}

func sum(initial int, nos ...int) int {
	var result int = initial
	for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	}
	return result
}
