package main

import (
	"fmt"
)

func main() {
	kayak := Product {
		name: "Kayak",
		category: "Watersports",
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println(kayak.price)
	fmt.Println("-------------------------------")

	myArray := [1]StockLevel{
		{
			Product: Product{"Kayak", "Watersport", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.96},
			count: 10,
		},
	}
	fmt.Println(myArray[0].Product.name)

	mySLice := []StockLevel{
		{
			Product: Product{"Kayak", "water", 275},
			Alternate: Product{"Lifej", "water", 56},
			count: 11,
		},
	}
	fmt.Println(mySLice[0].Alternate.name)

	myMap := map[string]StockLevel{
		"Kayak": {
			Product: Product{"Kayak", "water", 275},
			Alternate: Product{"Lifej", "water", 69},
			count: 11,
		},
	}
	fmt.Println(myMap["Kayak"].Product.name)
	fmt.Println("-------------------------------")

	p1 := &Product{
		name: "Vasya",
		category: "Human",
		price: 101,
	}

	p2 := &p1
	p1.name = "Grisha"
	fmt.Println((*p2).name, p1.name)
	calcTax(p1)
	fmt.Println(p1.price)
	fmt.Println(*p1)
}

type Product struct {
	name, category string
	price float64
}

type StockLevel struct {
	Product
	Alternate Product
	count int }

func calcTax(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}
