package cmd

import "inventory-app/internal/models"

type IInventory interface {
	AddProduct(name string, cost float64, units int) *models.Product
	GetProducts() []*models.Product
	GetByIndex(idx int) *models.Product
	GetValue() float64
	GetCount() int
}
