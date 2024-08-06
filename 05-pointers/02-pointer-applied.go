package main

import "fmt"

func main() {
	var no int
	no = 100
	fmt.Println("[main] Before incrementing, no :", no)
	fmt.Println("[main] address of no :", &no)
	increment(&no)
	fmt.Println("[main] After incrementing, no :", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap( /*  */ )
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x *int) {
	fmt.Println("[increment] address of x :", x)
	(*x)++
}

func swap( /*  */ ) /* no return values */ {

}
