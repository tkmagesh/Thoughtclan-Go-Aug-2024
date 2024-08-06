/*
refactor the below into maintainable functions
*/
package main

import (
	"errors"
	"fmt"
)

var ErrPrimeRange error = errors.New("invalid range for finding primes")

func main() {
	for {
		start, end, err := getRange()
		if err == ErrPrimeRange {
			fmt.Println("Invalid range. Try again!")
			continue
		}
		if err != nil {
			fmt.Println("err :", err)
			break
		}
		primes := generatePrimes(start, end)
		printPrimes(primes)
		break
	}
}

func printPrimes(primes []int) {
	for _, primeNo := range primes {
		fmt.Printf("prime no : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) []int {
	// var primes []int
	var primes []int = make([]int, 0, (end-start)/2)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
			fmt.Printf("len(primes) : %d, cap(primes) : %d\n", len(primes), cap(primes))
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

func getRange() (start, end int, err error) {
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	if start > end {
		err = ErrPrimeRange
		return
	}
	return
}
