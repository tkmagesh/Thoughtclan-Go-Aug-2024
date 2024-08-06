package main

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
func main() {
	var err error
	err = doSomething()
	if err != nil {
		fmt.Println("[main] doSomething(): err :", err)
		return
	}
	fmt.Println("[main] all good!")
}

func doSomething() error {
	if rand.Intn(20)%2 == 0 {
		fmt.Println("[doSomething] something done successfully")
		return nil
	}
	fmt.Println("[doSomething] error occurred")
	var err error
	err = errors.New("error occurred while doing something")
	return err
}
*/

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("[main] doSomething(): err :", err)
		return
	}
	fmt.Println("[main] all good!")
}

func doSomething() error {
	if rand.Intn(20)%2 == 0 {
		fmt.Println("[doSomething] something done successfully")
		return nil
	}
	fmt.Println("[doSomething] error occurred")
	return errors.New("error occurred while doing something")

}
