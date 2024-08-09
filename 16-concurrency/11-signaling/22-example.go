package main

import (
	"fmt"
	"time"
)

func main() {
	printNos()
}

func printNos() {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(5 * time.Second)
		timeoutCh <- time.Now()
	}()
LOOP:
	for no := 1; ; no++ {
		select {
		case <-timeoutCh:
			fmt.Println("timeout occurred")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("no : ", no)
		}

	}
}
