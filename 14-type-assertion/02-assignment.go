package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, "20"))
	fmt.Println(sum(10, "20", "abc"))
	fmt.Println(sum(10, 20, 30, 40, "50"))
}

func sum(nos ...interface{}) int {
	var result int
	for _, no := range nos {
		switch val := no.(type) {
		case int:
			result += val
		case string:
			if n, err := strconv.Atoi(val); err == nil {
				result += n
			}
		}
	}
	return result
}
