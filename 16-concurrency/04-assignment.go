/*
Execute the logic for checking if a given number is a prime number "concurrently"
*/
package main

import (
	"fmt"
	"sync"
)

var primes []int
var mutex sync.Mutex

func main() {
	start, end := getRange()
	generatePrimes(start, end)
	printPrimes(primes)
}

func printPrimes(primes []int) {
	for _, primeNo := range primes {
		fmt.Printf("prime no : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) {
	// var primes []int
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkPrime(wg, no)
	}
	wg.Wait()
}

func checkPrime(wg *sync.WaitGroup, no int) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primes = append(primes, no)
	}
	mutex.Unlock()
}

func getRange() (start, end int) {
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	return
}
