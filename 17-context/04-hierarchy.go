package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")
	cancelCtx, cancel := context.WithCancel(valCtx)
	go func() {
		fmt.Println("Hit ENTER to stop....")
		fmt.Scanln()
		cancel()
	}()
	doneCh := fn(cancelCtx)
	<-doneCh
}

func fn(ctx context.Context) <-chan struct{} {
	doneCh := make(chan struct{})
	fmt.Println("[fn] root-key : ", ctx.Value("root-key"))
	fnValCtx := context.WithValue(ctx, "fn-key", "fn-val")
	go func() {
		fnTimeoutCtx, cancel := context.WithTimeout(fnValCtx, 1*time.Minute)
		defer cancel()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go f1(fnTimeoutCtx, wg)

		wg.Add(1)
		go f2(fnTimeoutCtx, wg)

		wg.Wait()
		close(doneCh)
	}()
	return doneCh
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[f1] root-key : ", ctx.Value("root-key"))
	fmt.Println("[f1] fn-key : ", ctx.Value("fn-key"))
LOOP:
	for i := 0; ; i += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[f1] cancellation signal received")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("[f1] : %d\n", i)
		}
	}
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[f2] root-key : ", ctx.Value("root-key"))
	fmt.Println("[f2] fn-key : ", ctx.Value("fn-key"))
LOOP:
	for i := 1; ; i += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[f2] cancellation signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("[f2] : %d\n", i)
		}

	}
}
