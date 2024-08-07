package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// ver 1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// ver 4.0
	/*
		add := getLogOperation(add)
		subtract := getLogOperation(subtract)

		add(100, 200)
		subtract(100, 200)
	*/

	/* ver 5.0 */
	logAdd := getLogOperation(add)
	profiledLogAdd := getProfileOperation(logAdd)
	profiledLogAdd(100, 200)

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	// ver 4.0
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/
}

/* final version */
type OperationFn func(int, int)

func getProfileOperation(op OperationFn) OperationFn {
	return func(i1, i2 int) {
		start := time.Now()
		op(i1, i2)
		elapsed := time.Since(start)
		fmt.Println("elapsed :", elapsed)
	}
}

func getLogOperation(op OperationFn) OperationFn {
	return func(x, y int) {
		log.Println("Invocation started")
		op(x, y)
		log.Println("Invocation completed")
	}
}

/* ver 5.0 */
/*
func getProfileOperation(op func(int, int)) func(int, int) {
	return func(i1, i2 int) {
		start := time.Now()
		op(i1, i2)
		elapsed := time.Since(start)
		fmt.Println("elapsed :", elapsed)
	}
}

*/
/* ver 4.0 */
/*
func getLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Invocation started")
		op(x, y)
		log.Println("Invocation completed")
	}
}
*/

// ver 3.0
func logOperation(op func(int, int), x, y int) {
	log.Println("Invocation started")
	op(x, y)
	log.Println("Invocation completed")
}

// ver 2.0
func logAdd(x, y int) {
	log.Println("Invocation started")
	add(x, y)
	log.Println("Invocation completed")
}

func logSubtract(x, y int) {
	log.Println("Invocation started")
	subtract(x, y)
	log.Println("Invocation completed")
}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
