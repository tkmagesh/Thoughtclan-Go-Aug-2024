/*
refactor the below into maintainable functions
*/
package main

import "fmt"

func main() {
	start, end := getRange()
	primes := generatePrimes(start, end)
	printPrimes(primes)
}

func printPrimes(primes []int) {
	for _, primeNo := range primes {
		fmt.Printf("prime no : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func getRange() (start, end int) {
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	return
}
