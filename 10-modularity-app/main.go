package main

import (
	"fmt"

	/* "github.com/tkmagesh/thoughtclan-go-aug-2024/10-modularity-app/calculator" */
	"github.com/fatih/color"
	calc "github.com/tkmagesh/thoughtclan-go-aug-2024/10-modularity-app/calculator" //package alias
	"github.com/tkmagesh/thoughtclan-go-aug-2024/10-modularity-app/calculator/utils"
)

func main() {
	// fmt.Printf("%q app started\n", appName)
	color.Red("%q app started\n", appName)
	run()
	/*
		calculator.Add(100, 200)
		calculator.Subtract(100, 200)
		fmt.Println("OpCount = ", calculator.OpCount())
	*/

	calc.Add(100, 200)
	calc.Subtract(100, 200)
	fmt.Println("OpCount = ", calc.OpCount())

	fmt.Println("IsEven(21) : ", utils.IsEven(21))
	fmt.Println("IsOdd(21) : ", utils.IsOdd(21))
}
