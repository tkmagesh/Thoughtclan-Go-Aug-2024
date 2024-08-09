package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)

	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()

	printNos(timeoutCtx)
	if err := timeoutCtx.Err(); err == context.DeadlineExceeded {
		fmt.Println("cancelled due to timeout")
	} else if err == context.Canceled {
		fmt.Println("programmatically cancelled")
	}

}

func printNos(ctx context.Context) {
LOOP:
	for no := 1; ; no++ {
		select {
		case <-ctx.Done():
			fmt.Println("stop signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("no : ", no)
		}

	}
}
