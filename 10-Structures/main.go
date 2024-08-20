package main

import (
	"fmt"
	"sort"
	"strings"
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

	fmt.Println(DNAtoRNA("DCAUURWEWUFEETTTTHHJTRLTLTJ"))

	// arrays
	ar1 := []int{11, 2, 33, 4, 5} 
	ar2 := []int{4, 72, 8, 9, 10}
	fmt.Println(myArrayFunc(ar1, ar2))
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

func DNAtoRNA (dna string) string {
	str := []string{}
	for _, val := range dna {
		temp := string(val)
		if string(val) == "T" {
			temp = "U"
		}
		str = append(str, temp)
	}
	str2 := strings.Join(str, "")
	return str2
}

func myArrayFunc(arr1, arr2 []int) []int {
	
	arr1 = append(arr1, arr2...)
	sort.Ints(arr1)
	myArr := make([]int, 0, len(arr1))
	myArr = append(myArr, arr1[0])
	for i := 1; i < len(arr1); i++ {
		if arr1[i] <= arr1[i - 1] {
			continue
		}
		myArr = append(myArr, arr1[i])
	}
 	return myArr
}
