package main

import (
	"fmt"
)

func main() {
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["pen"] = 5
		productRanks["pencil"] = 1
		productRanks["marker"] = 3
	*/

	// var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// var productRanks = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	productRanks := map[string]int{
		"pen":    5,
		"pencil": 1,
		"marker": 3,
	}
	fmt.Println(productRanks)

	fmt.Println("len(productRanks) : ", len(productRanks))

	// iterating a map
	fmt.Println("iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of a key
	// keyToCheck := "pen"
	keyToCheck := "notepad"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key - %q does not exist\n", keyToCheck)
	}

	// removing a key
	// keyToDelete := "pen"
	keyToDelete := "notepad"
	delete(productRanks, keyToDelete)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToDelete, productRanks)
}
