/*
Execute the logic for checking if a given number is a prime number "concurrently"
Ensure that there are only N (ex 10) goroutines started (workers)
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	start, end := getRange()
	outputCh := generatePrimes(10, start, end)
	printPrimes(outputCh)
}

func printPrimes(ch <-chan int) {
	for primeNo := range ch {
		fmt.Printf("prime no : %d\n", primeNo)
	}
}

func generatePrimes(workerCount int, start, end int) <-chan int {
	outputCh := make(chan int)
	inputCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for range workerCount {
			wg.Add(1)
			go checkPrime(wg, inputCh, outputCh)
		}
		for no := start; no <= end; no++ {
			inputCh <- no
		}
		close(inputCh)
		wg.Wait()
		close(outputCh)
	}()
	return outputCh
}

func checkPrime(wg *sync.WaitGroup, inputCh <-chan int, outputCh chan<- int) {
	defer wg.Done()
	fmt.Println("Starting checkPrime...")
	for no := range inputCh {
		if isPrime(no) {
			outputCh <- no
		}
	}
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
