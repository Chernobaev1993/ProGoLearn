package main

import (
	"fmt"
	. "packages/fmt"
	"packages/store"
	"packages/store/cart"
	// "github.com/fatih/color"
)

func main() {
	prod := store.NewProduct("Kayak", "Watersports", 104.3)

	cart := cart.Cart{
		CustomerName: "Alex",
		Products: []store.Product{*prod},
	}
	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))

	fmt.Printf("Name: %s, cat: %s, price: %.2f\n", prod.Name, prod.Category, prod.GetPrice())
	prod.SetPrice(15.3)
	fmt.Printf("Name: %s, cat: %s, price: %.2f\n", prod.Name, prod.Category, prod.GetPrice())
	fmt.Println("Price:", ToCurrency(prod.GetPrice()))

	// color.Green("Name: " + cart.CustomerName)
	// color.Cyan("Total: " + ToCurrency(cart.GetTotal()))
}