/*

Given the following string
"Cillum anim esse laborum aliqua nostrud Reprehenderit qui et occaecat ad adipisicing incididunt nulla Amet incididunt eiusmod sint fugiat dolor velit cillum cillum fugiat Mollit dolore aliqua reprehenderit eiusmodConsectetur excepteur cillum reprehenderit sint Lorem cupidatat sint id laboris anim cillum ipsum nostrud anim irure fugiat mollit laboris mollit Eu aliquip ex pariatur excepteur dolor laboris laboris Eiusmod id sunt eu excepteur laboris consectetur deserunt Veniam consectetur aliqua occaecat occaecat elit commodo ex magna deserunt nisi fugiat Lorem sunt ut ad sunt do ullamco consequat velit aute minim exercitation in fugiat"

find out the "size of words" that occurs the most
ex output:
4 letter words occurs the most with 12 occurances

Hint:
use strings.Split() functions
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Cillum anim esse laborum aliqua nostrud Reprehenderit qui et occaecat ad adipisicing incididunt nulla Amet incididunt eiusmod sint fugiat dolor velit cillum cillum fugiat Mollit dolore aliqua reprehenderit eiusmodConsectetur excepteur cillum reprehenderit sint Lorem cupidatat sint id laboris anim cillum ipsum nostrud anim irure fugiat mollit laboris mollit Eu aliquip ex pariatur excepteur dolor laboris laboris Eiusmod id sunt eu excepteur laboris consectetur deserunt Veniam consectetur aliqua occaecat occaecat elit commodo ex magna deserunt nisi fugiat Lorem sunt ut ad sunt do ullamco consequat velit aute minim exercitation in fugiat"
	words := strings.Split(str, " ")
	wordsCountBySize := getCountBySize(words)
	maxCount, maxSize := getMax(wordsCountBySize)
	fmt.Printf("%d letter occurs the most with %d occurrances\n", maxSize, maxCount)

}

func getMax(wordsCountBySize map[int]int) (maxCount, maxSize int) {
	for size, count := range wordsCountBySize {
		if count > maxCount {
			maxCount = count
			maxSize = size
		}
	}
	return
}

func getCountBySize(words []string) map[int]int {
	wordsCountBySize := make(map[int]int)
	for _, word := range words {
		size := len(word)
		wordsCountBySize[size]++
	}
	return wordsCountBySize
}
