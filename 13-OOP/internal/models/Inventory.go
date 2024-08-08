package models

import "sort"

type Inventory struct {
	products *Products
}

func NewInventory() *Inventory {
	products := NewProducts()
	return &Inventory{
		products: products,
	}
}

func (inventory *Inventory) AddProduct(name string, cost float64, units int) *Product {
	newProduct := NewProduct(len(inventory.products.list)+1, name, cost, units)
	inventory.products.AddNew(newProduct)
	return newProduct
}

func (inventory *Inventory) GetProducts() []*Product {
	return inventory.products.list
}

func (inventory *Inventory) GetByIndex(idx int) *Product {
	return inventory.products.list[idx]
}

func (inventory *Inventory) GetValue() float64 {
	var totalValue float64
	for _, p := range inventory.products.list {
		totalValue += p.GetValue()
	}
	return totalValue
}

func (inventory *Inventory) GetCount() int {
	return inventory.products.GetCount()
}

func (inventory *Inventory) SortByCost() {
	sort.Sort(inventory.products)
}
