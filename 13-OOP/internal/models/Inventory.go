package models

type Inventory struct {
	products []*Product
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make([]*Product, 0),
	}
}

func (inventory *Inventory) AddProduct(name string, cost float64, units int) *Product {
	newProduct := NewProduct(len(inventory.products)+1, name, cost, units)
	inventory.products = append(inventory.products, newProduct)
	return newProduct
}

func (inventory *Inventory) GetProducts() []*Product {
	return inventory.products
}

func (inventory *Inventory) GetByIndex(idx int) *Product {
	return inventory.products[idx]
}

func (inventory *Inventory) GetValue() float64 {
	var totalValue float64
	for _, p := range inventory.products {
		totalValue += p.GetValue()
	}
	return totalValue
}

func (inventory *Inventory) GetCount() int {
	return len(inventory.products)
}
