package main

import (
	"fmt"
	. "packages/fmt"
	"packages/store"
)

func main() {
	prod := store.NewProduct("Kayak", "Watersports", 104.3)
	fmt.Printf("Name: %s, cat: %s, price: %.2f\n", prod.Name, prod.Category, prod.GetPrice())
	prod.SetPrice(15.3)
	fmt.Printf("Name: %s, cat: %s, price: %.2f\n", prod.Name, prod.Category, prod.GetPrice())
	fmt.Println("Price:", ToCurrency(prod.GetPrice()))
}