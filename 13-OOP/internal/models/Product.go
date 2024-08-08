package models

import "fmt"

type Product struct {
	Id    int
	Name  string
	Cost  float64
	Units int
}

func (p Product) GetValue() float64 {
	return p.Cost * float64(p.Units)
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d, Value = %0.2f", p.Id, p.Name, p.Cost, p.Units, p.GetValue())
}

func NewProduct(id int, name string, cost float64, units int) *Product {
	return &Product{
		Id:    id,
		Name:  name,
		Cost:  cost,
		Units: units,
	}
}
