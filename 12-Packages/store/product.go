// Package store provides types and methods 
// commonly required for online sales
package store

var standartTax = newTaxRate(0.25, 20)

// Product describes an item for sale
type Product struct {
	Name, Category string
	price float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) GetPrice() float64 {
	return standartTax.calcTax(p)
}

func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}