/*
Execute the logic for checking if a given number is a prime number "concurrently"
Ensure that there are only N (ex 10) goroutines started (workers)
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var start int
	fmt.Println("Enter the start:")
	fmt.Scanln(&start)
	rootCtx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()
	outputCh := generatePrimes(timeoutCtx, 10, start)
	printPrimes(outputCh)
}

func printPrimes(ch <-chan int) {
	for primeNo := range ch {
		fmt.Printf("prime no : %d\n", primeNo)
	}
}

func generatePrimes(ctx context.Context, workerCount int, start int) <-chan int {
	outputCh := make(chan int)
	inputCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for range workerCount {
			wg.Add(1)
			go checkPrime(ctx, wg, inputCh, outputCh)
		}
		go func() {
			for no := start; ; no++ {
				inputCh <- no
			}
		}()
		wg.Wait()
		close(outputCh)
	}()
	return outputCh
}

func checkPrime(ctx context.Context, wg *sync.WaitGroup, inputCh <-chan int, outputCh chan<- int) {
	defer wg.Done()
	fmt.Println("Starting checkPrime...")
LOOP:
	for {
		select {
		case no := <-inputCh:
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				outputCh <- no
			}
		case <-ctx.Done():
			break LOOP
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
