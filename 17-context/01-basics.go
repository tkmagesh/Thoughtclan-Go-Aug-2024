package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()
	printNos(cancelCtx)

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
