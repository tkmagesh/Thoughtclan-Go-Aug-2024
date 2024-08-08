package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}
	flag.IntVar(&count, "count", 0, "Number of goroutines to execute")
	flag.Parse()
	fmt.Printf("Starting %d goroutines... Hit ENTER to start!!\n", count)
	fmt.Scanln()
	for idx := range count {
		wg.Add(1)
		go fn(wg, idx+1)
	}
	wg.Wait()
	fmt.Println("Done!")
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	randTime := rand.Intn(20)
	time.Sleep(time.Duration(randTime) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)

}
