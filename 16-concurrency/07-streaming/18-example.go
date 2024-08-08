package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		break
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
