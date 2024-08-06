package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// add new items to the slice
	nos = append(nos, 10)
	fmt.Println(nos)

	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	// append another slice
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d and nos2[0] = %d\n", nos[0], nos2[0])

	sort(nos)
	fmt.Println(nos)

	// slicing
	fmt.Println("nos[2:5] = ", nos[2:5])
	fmt.Println("nos[:5] = ", nos[:5])
	fmt.Println("nos[2:] = ", nos[2:])

	subset := nos[:5]
	subset[0] = 8888
	fmt.Println("subset :", subset)
	fmt.Println("nos :", nos)

	nosCopy := make([]int, len(nos))
	copy(nosCopy, nos)
	fmt.Println(nosCopy)

	nos[0] = -100
	fmt.Println("nos :", nos)
	fmt.Println("nosCopy :", nosCopy)

}

func sort(list []int) /* no return values */ {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
