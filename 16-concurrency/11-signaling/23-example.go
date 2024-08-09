package main

import (
	"fmt"
	"time"
)

func main() {
	printNos()
}

func printNos() {
	timeoutCh := time.After(5 * time.Second)
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
