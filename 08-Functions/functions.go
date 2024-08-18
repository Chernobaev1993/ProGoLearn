package main

import (
	"fmt"
)

func main() {
	printPrice()
	printPrice2("Kayak", 100, 0.2)
	myFunc1(10, "Hi", "How", "Are you?")
	mySlice := []string{"hello", "my"}
	myFunc1(10, mySlice...)
	printSuppliers("Suckin")
	a, b := 10, 20
	fmt.Println("Befor func:", a, b)
	funcPointer1(&a, &b)
	fmt.Println("After func:", a, b)
	fmt.Println("-------------------------------------")

}

func printPrice() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax", kayakTax)
	fmt.Println("-------------------------------------")
}

func printPrice2(prod string, price, taxRate float64) {
	taxAmmount := price * taxRate
	fmt.Printf("%s price is %.2f and tax is %.2f\n", prod, price, taxAmmount)
	fmt.Println("-------------------------------------")
}

func myFunc1(x int, supp ...string) {
	for _, v := range supp {
		fmt.Printf("%s and %d\n", v, x)
	}
	fmt.Println("-------------------------------------")
}

func printSuppliers(product string, suppliers ...string ) {
    for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:",supplier) 
	}
	fmt.Println(suppliers)
	fmt.Println("-------------------------------------")
}

func funcPointer1(x, y *int) {
	fmt.Println("Before swap:", *x, *y)
	*x, *y = *y, *x
	fmt.Println("After swap:", *x, *y)
}
