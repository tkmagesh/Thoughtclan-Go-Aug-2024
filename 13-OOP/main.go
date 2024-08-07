package main

import (
	"fmt"
	"inventory-app/cmd"
	"inventory-app/internal/models"
)

func main() {
	inventory := models.NewInventory()
	/*
		inventory.AddProduct("Pen", 10, 10)
		inventory.AddProduct("Pencil", 5, 20)
		inventory.AddProduct("Marker", 50, 2)
		inventory.AddProduct("Notepad", 20, 30)
	*/
	commands := cmd.NewCommands(inventory)
	var userChoice int
LOOP:
	for {
		fmt.Println("1. Add new product")
		fmt.Println("2. Get products count")
		fmt.Println("3. Get product by index")
		fmt.Println("4. List all the products")
		fmt.Println("5. Get inventory value")
		fmt.Println("6. Exit")
		fmt.Scanln(&userChoice)
		switch userChoice {
		case 1:
			if err := commands.AddProduct(); err != nil {
				fmt.Println(err)
			}
			continue LOOP
		case 2:
			commands.QueryCount()
			continue LOOP
		case 3:
			if err := commands.QueryByIndex(); err != nil {
				fmt.Println(err)
			}
			continue LOOP
		case 4:
			commands.QueryProducts()
			continue LOOP
		case 5:
			commands.QueryInventoryValue()
			continue LOOP
		case 6:
			break LOOP
		}
	}
}
