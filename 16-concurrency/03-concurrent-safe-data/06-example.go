package main

import "fmt"

var count int

func main() {
	for range 300 {
		increment()
	}
	fmt.Println("count :", count)
}

func increment() {
	count++
}
