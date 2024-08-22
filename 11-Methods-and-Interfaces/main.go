package main

import (
	"fmt"
)

type Expense interface {
    getName() string
    getCost(annual bool) float64
}

type Person struct {
	name, city string
}


func main() {
	var expense Expense = &Product{"Kayak", "Water", 275}

	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersport", 48.95},
		Service{"Boat cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
		switch value := item.(type) {
		case Product: fmt.Println(value.name, value.price)
		case *Product: fmt.Println(value.name, value.price)
		case Service: fmt.Println(value.description, value.monthlyFee)
		case Person: fmt.Println(value.city, value.name)
		case *Person: fmt.Println(value.city, value.name)
		case string, bool, int: println("Built0in type:", value)
		default: fmt.Println("Default:", value)
		}
	}
}


