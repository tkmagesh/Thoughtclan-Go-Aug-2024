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

// factory function to hide the complexity of constructing a PerishableProduct object
func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		var milk PerishableProduct = PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Nandini Milk",
				Cost: 48,
			},
			Expiry: "2 Days",
		}
	*/

	// using the "factory" to create an instance
	milk := NewPerishableProduct(100, "Nandini Milk", 48, "2 Days")

	fmt.Println(milk)
	// fmt.Println(milk.Product.Name)
	fmt.Println(milk.Name)

	// Format 'milk' and print
	// fmt.Println(FormatPerishableProduct(milk))
	fmt.Println(milk.Format())

	// Apply 10% discount to milk
	// ApplyDiscount(&(milk.Product), 10)
	milk.ApplyDiscount(10)

	// FormatPerishableProduct 'milk' and print
	// fmt.Println(FormatPerishableProduct(milk))
	fmt.Println(milk.Format())
}

func (p *Product) Format() string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

// overriding Product.Format()
func (pp *PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}
