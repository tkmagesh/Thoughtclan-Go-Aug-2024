package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {

	// type inference
	var product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	product.WhoAmI()

	// fmt.Println(Format(product))
	fmt.Println(product.Format())

	// ApplyDiscount(&product, 10)
	// (&product).ApplyDiscount(10)
	// Even though the ApplyDiscount has a pointer receiver, it can be invoked with a value
	product.ApplyDiscount(10)

	// fmt.Println(Format(product))
	fmt.Println(product.Format())

	// Pointer
	pencilPtr := &Product{Id: 101, Name: "Pencil", Cost: 5}
	fmt.Println(pencilPtr.Format())
	pencilPtr.ApplyDiscount(20)
	fmt.Println(pencilPtr.Format())
}

/*
1. Write a Format() function which returns a printable (id = 100, name = "Pen", cost = 10) string of the given product
2. Write a ApplyDiscount() function which updates the given product cost with a discount
*/

func (Product) WhoAmI() {
	fmt.Println("I am a product!")
}

func (p Product) Format() string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
