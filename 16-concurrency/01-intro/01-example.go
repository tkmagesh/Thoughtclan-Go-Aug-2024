package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule execution of f1 through the builtin scheduler
	f2()

	// DO NOT use the below techniques (poor man's synchronization techniques)
	// time.Sleep(1 * time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
