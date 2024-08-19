package main

import (
	"fmt"
)

type calcFunc func(float64) float64

func main() {
	products := map[string]float64{
		"Kayak": 275,
		"Lifejacket": 48.95,
	}
	myFunc1(products)
	fmt.Println("----------------------------------")
	myFunc2(products)
	fmt.Println("----------------------------------")
	myFunc3(products)
	fmt.Println("----------------------------------")

}

func myFunc3(prod map[string]float64) {
    for product, price := range prod {
        printPrice(product, price, selectCalculator(price))
	}	
}

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

func myFunc2(prod map[string]float64) {
	for product, price := range prod {
        if (price > 100) {
            printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		} 
	}
}

func myFunc1(prod map[string]float64) {
	for product, price := range prod {
		var calcFunction func(float64) float64
		fmt.Println("Функция ничего не присвоено?", calcFunction == nil)
		if price > 100 {
			calcFunction = calcWithTax
		} else {
			calcFunction = calcWithoutTax
		}
		fmt.Println("Функция ничего не присвоено?", calcFunction == nil)
		totalPrice := calcFunction(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(prod string, price float64, calc calcFunc) {
	fmt.Println("Product:", prod, ". Price:", calc(price))
}

