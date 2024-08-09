package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	printNos(stopCh)

}

func printNos(stopCh <-chan struct{}) {
LOOP:
	for no := 1; ; no++ {
		select {
		case <-stopCh:
			fmt.Println("stop signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("no : ", no)
		}

	}
}
