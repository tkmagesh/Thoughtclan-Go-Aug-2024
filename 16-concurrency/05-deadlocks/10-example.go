package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(10) // increment the counter by 1
	go f1()    // schedule execution of f1 through the builtin scheduler
	f2()

	wg.Wait() // block until the counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	randTime := rand.Intn(20)
	fmt.Printf("f1 will take %d seconds\n", randTime)
	time.Sleep(time.Duration(randTime) * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
