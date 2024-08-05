package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	greetFullName("Magesh", "Kuppan")
	msg := getGreet("Suresh")
	fmt.Println(msg)
	fmt.Println(add(100, 200))
	fmt.Println(divide(100, 7))
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d, remainder = %d\n", q, r)
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/

}

/* 0 parameters, 0 return values */
func sayHi() {
	fmt.Println("Hi!")
}

/* 1 parameter, 0 return values */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* 2 parameters, 0 return values */
/*
func greetFullName(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greetFullName(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

/* 1 paramete, 1 return value */
func getGreet(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

/* 2 parameters, 1 return values */
func add(x, y int) int {
	return x + y
}

/* 2 parameters, 2 return values */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// named result syntax
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
