package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}

}

func genNos(ch chan<- int) {
	count := rand.Intn(20)
	// count := 5
	for idx := range count {
		ch <- (idx + 1) * 10
	}
	close(ch)
}
