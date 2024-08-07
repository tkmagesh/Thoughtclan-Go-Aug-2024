package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	// Dummy
	Product
	Expiry string
}

func main() {
	var milk PerishableProduct = PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Nandini Milk",
			Cost: 48,
		},
		Expiry: "2 Days",
	}
	fmt.Println(milk)
	// fmt.Println(milk.Product.Name)
	fmt.Println(milk.Name)

	// Format 'milk' and print
	// Apply 10% discount to milk
	// Format 'milk' and print
}

func Format(p Product) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
