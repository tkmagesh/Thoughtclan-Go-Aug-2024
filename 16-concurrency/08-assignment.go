/*
Execute the logic for checking if a given number is a prime number "concurrently"
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	start, end := getRange()
	ch := generatePrimes(start, end)
	printPrimes(ch)
}

func printPrimes(ch <-chan int) {
	for primeNo := range ch {
		fmt.Printf("prime no : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go checkPrime(wg, no, ch)
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

func checkPrime(wg *sync.WaitGroup, no int, ch chan int) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	ch <- no
}

func getRange() (start, end int) {
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	return
}
