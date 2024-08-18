package main

import (
	"fmt"
	"strconv"
)

func main() {
	ifFunc()
	forRange()
	metki()
}

func ifFunc() {
	priceString := "275f"
	if price, err := strconv.Atoi(priceString); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Price is", price)
	}

}

func forRange() {
	str := "каяк"
	for i, c := range str {
		fmt.Println(i, string(c))
	}
}

func metki() {
	counter := 0
    target: fmt.Println("Counter", counter)
    counter++
    if (counter < 5) {
		goto target 
	}			
}
