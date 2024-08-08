package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1) // increment the counter by 1
	go func() {
		defer wg.Done()
		f1()
	}()
	f2()

	wg.Wait() // block until the counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	randTime := rand.Intn(20)
	fmt.Printf("f1 will take %d seconds\n", randTime)
	time.Sleep(time.Duration(randTime) * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
