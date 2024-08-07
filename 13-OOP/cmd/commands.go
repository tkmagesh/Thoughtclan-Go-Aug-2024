package cmd

import (
	"errors"
	"fmt"
	"inventory-app/internal/models"
	"strconv"
	"strings"
)

type Commands struct {
	inventory *models.Inventory
}

var ErrUserInput error = errors.New("user input error")

func NewCommands(inventory *models.Inventory) *Commands {
	return &Commands{
		inventory: inventory,
	}
}

func (commands *Commands) AddProduct() error {
	var userInput string
	fmt.Println("Enter the product info (name,cost,units) :")
	fmt.Scanln(&userInput)
	fields := strings.Split(userInput, ",")

	name := fields[0]
	if cost, err := strconv.ParseFloat(fields[1], 64); err == nil {
		if units, err := strconv.Atoi(fields[2]); err == nil {
			newProduct := commands.inventory.AddProduct(name, cost, units)
			fmt.Println("New product added :", newProduct.ToString())
			return nil
		}
	}
	return ErrUserInput
}

func (commands *Commands) QueryCount() {
	count := commands.inventory.GetCount()
	fmt.Printf("Products Count : %d\n", count)
}

func (commands *Commands) QueryByIndex() error {
	var idx int
	fmt.Println("Enter the index :")
	if _, err := fmt.Scanln(&idx); err != nil {
		return ErrUserInput
	}
	if product := commands.inventory.GetByIndex(idx); product != nil {
		fmt.Println(product.ToString())
		return nil
	}
	fmt.Println("Product not found!")
	return nil
}

func (commands *Commands) QueryProducts() {

	if products := commands.inventory.GetProducts(); len(products) != 0 {
		for _, p := range products {
			fmt.Println(p.ToString())
		}
		return
	}
	fmt.Println("No products found!")
}

func (commands *Commands) QueryInventoryValue() {
	invValue := commands.inventory.GetValue()
	fmt.Printf("Inventory value : %02.f\n", invValue)
}
