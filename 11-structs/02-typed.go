package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
}

func main() {
	/* id = 100, name = "Pen", cost = 10 */

	/*
		var product Product
		product.id = 100
		product.name = "Pen"
		product.cost = 10
	*/

	// var product Product = Product{ id : 100, name : "Pen", cost : 10}
	/*
		var product Product = Product{
			id:   100,
			name: "Pen",
			cost: 10,
		}
	*/

	// type inference
	var product = Product{
		id:   100,
		name: "Pen",
		cost: 10,
	}
	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	// fmt.Printf("%+v\n", product)
	fmt.Println(Format(product)) //=> should print "id = 100, name = "Pen", cost = 10"
	ApplyDiscount(&product, 10)
	fmt.Println(Format(product)) //=> should print "id = 100, name = "Pen", cost = 9"
}

/*
1. Write a Format() function which returns a printable (id = 100, name = "Pen", cost = 10) string of the given product
2. Write a ApplyDiscount() function which updates the given product cost with a discount
*/

func Format(p Product) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}
