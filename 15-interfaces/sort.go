package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	Id    int
	Name  string
	Cost  float64
	Units int
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d", p.Id, p.Name, p.Cost, p.Units)
}

type Products []Product

func (products Products) String() string {
	builder := strings.Builder{}
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%s\n", p))
	}
	return builder.String()
}

// sort.Interface implementation
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

// Swap swaps the elements with indexes i and j.
func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

// Sort By Cost
type SortByCost struct {
	Products
}

// override Products.Less()
func (sb *SortByCost) Less(i, j int) bool {
	return sb.Products[i].Cost < sb.Products[j].Cost
}

func main() {
	products := Products{
		Product{3, "Pen", 10, 20},
		Product{1, "Pencil", 5, 100},
		Product{4, "Marker", 50, 10},
		Product{2, "Notepad", 15, 50},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Sort by id")
	sort.Sort(products)
	fmt.Println(products)

	fmt.Println("Sort by cost")
	sort.Sort(&SortByCost{products})
	fmt.Println(products)

	fmt.Println("Sort by units")
	sort.Slice(products, func(i, j int) bool {
		return products[i].Units < products[j].Units
	})
	fmt.Println(products)

}
