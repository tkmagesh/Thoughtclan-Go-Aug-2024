/*
refactor the below into maintainable functions
*/
package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime no :", no)
	}
}
