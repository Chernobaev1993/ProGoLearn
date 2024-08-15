package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	arithmetic()
	wrap()
	cel_div()
	incr_decr()
	concat()
	compare()
	compare_pointers()
	logic()
	preo()
	preoFloat()
	parse()
	testStrconv()
}

func arithmetic() {
	price, tax := 275.00, 27.4
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax
	fmt.Printf("sum: %0.2f, dif: %0.2f, product: %0.2f, quotient: %0.2f\n", sum, difference, product, quotient)
	fmt.Println("-------------------------------------------")
}

func wrap() {
	var intVal = math.MaxInt64
	floatVal := math.MaxFloat64
	fmt.Printf("intVal: %d, intVal * 2: %d\n",intVal, intVal * 2)
	fmt.Printf("floatVal: %.3f, floatVal * 2: %f\n",floatVal, floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))
	fmt.Println("-------------------------------------------")
}

func cel_div() {
	posRes := 3 % 2
	negRes := -3 % 2
	absRes := math.Abs(float64(negRes))
	fmt.Printf("posRes: %d, negRes: %d, absRes: %.2f\n", posRes, negRes, absRes)
	fmt.Println("-------------------------------------------")
}

func incr_decr() {
	incr1 := 10.2
	incr1++
	fmt.Println(incr1)
	incr1 += 2
	fmt.Println(incr1)
	incr1 -= 2
	fmt.Println(incr1)
	incr1--
	fmt.Println(incr1)
	fmt.Println("-------------------------------------------")
}

func concat() {
	greeting := "Hello"
	language := "Go"
	combinedString := greeting + ", " + language
	fmt.Println(combinedString)
	fmt.Println("-------------------------------------------")
}

func compare() {
	first := 100
	const second = 200.00
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second
	fmt.Println(equal, notEqual, lessThan, lessThanOrEqual, greaterThan, greaterThanOrEqual)
	fmt.Println("-------------------------------------------")
}

func compare_pointers() {
	value1 := 100
	value2 := 100
	pointer1 := &value1
	pointer2 := &value1
	pointer3 := &value2
	fmt.Println(pointer1 == pointer2, pointer1 == pointer3, *pointer1 == *pointer3)
	fmt.Println("-------------------------------------------")
}

func logic() {
	maxKmh := 100
	passCap := 4
	airbags := true
	familyCar := passCap > 2 && airbags
	sportCar := maxKmh > 100 || passCap == 2
	canCategorize := !familyCar && !sportCar
	fmt.Printf("family: %v, sport: %v, categor: %v\n",familyCar, sportCar, canCategorize)
	fmt.Println("-------------------------------------------")
}

func preo() {
	kayak := 275
	socBall := 19.5
	total := float64(kayak) + socBall
	total2 := kayak + int(socBall)	// отброс точности
	total3 := int8(total2)			// переполнение
	fmt.Println(total, total2, total3)
	fmt.Println("-------------------------------------------")
}

func preoFloat() {
	myFloat1 := 27.3
	fmt.Println(
		math.Ceil(myFloat1), 
		math.Floor(myFloat1), 
		math.Round(myFloat1),
		math.RoundToEven(myFloat1),
	)
	fmt.Println("-------------------------------------------")
}

func parse() {
	myParse1, err := strconv.ParseBool("TRUE")
	// myParse1, err = strconv.ParseBool("0")
	fmt.Println(myParse1, err)

	myParse2, err := strconv.ParseBool("not true")
	if err == nil {
		fmt.Println(myParse2)
	} else {
		fmt.Println(err)
	}

	myInt1 := "100"
	fmt.Println(strconv.ParseInt(myInt1, 0, 8))
	fmt.Println(strconv.ParseInt(myInt1, 2, 8))
	myInt1 = "500"
	fmt.Println(strconv.ParseInt(myInt1, 10, 8))
	fmt.Println("-------------------------------------------")
}

func testStrconv() {
	myBool1 := strconv.FormatBool(true)
	myInt1 := strconv.FormatInt(100, 10)
	tempVal := 100
	myInt2 := strconv.FormatInt(int64(tempVal), 2) // представить в двоичном формате
	myUint1 := strconv.FormatUint(100, 10)
	fmt.Println(myBool1, reflect.TypeOf(myBool1))
	fmt.Println(myInt1, reflect.TypeOf(myInt1))
	fmt.Println(myInt2)
	fmt.Println(myUint1)


	fmt.Println("-------------------------------------------")
}