package main

import "fmt"

func main() {
	var i int8 = 100
	var f float64
	f = float64(i) //use the type name like a function for type conversion
	fmt.Println(f)
}
