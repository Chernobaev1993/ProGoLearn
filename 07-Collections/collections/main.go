package main

import (
	"fmt"
	// "math/rand"
	"sort"
	// "strings"
)

func main() {
	arrays1()
	array_pointer()
	slice1()
	append1()
	test1()
	test2()
	test3()
	capTest1()
	sliceFromSlice()
	copyTest()
	sortTest()
	getBaseArray()
	map1()
	stringsAreRune()
	test4()
	unic()
}

func arrays1() {
	var names [3]string
	names[0] = "sanek"
	names[1] = "ksu"
	names[2] = "fedya"
	fmt.Println(names)

	names2 := [3]string{"sanek", "ksu", "fedya"}
	fmt.Println(names2)
	fmt.Println("--------------------------")
}

func array_pointer() {
	names := [...]string{"sanek", "ksu", "fedya"}
	otherArray := names
	names[0] = "igor"
	fmt.Println("names:", names)
	fmt.Println("otherArray:", otherArray)

	otherArray2 := &names
	names[0] = "SANYA"
	fmt.Println("names:", names)
	fmt.Println("otherArray2:", *otherArray2)

	pointer1 := &names[1]
	names[1] = "KSENIA"
	*pointer1 = "123"
	fmt.Println(*pointer1, names)
	fmt.Println("--------------------------")
}

func slice1() {
	names := make([]string, 3)
	names[0] = "sanya"
	names[1] = "ksu"
	names[2] = "fedya"
	fmt.Println(names)

	names2 := []string{"sanya", "ksu", "fedya"}
	fmt.Println(names2)
	fmt.Println("--------------------------")
}

func append1() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	names = append(names, "Hat", "Gloves")
	fmt.Println(names)

	names2 := []string{"sanya", "ksu", "fedya"}
	fmt.Println("names2: ", names2, len(names), cap(names))
	appNames := append(names, "larisa", "katya")
	fmt.Println("appNames: ", appNames, len(appNames), cap(appNames))

	fmt.Println("--------------------------")
}

func test1() {
	names := [3]string{"ksu", "ale", "fed"}
	s1 := names[:]
	s2 := append(s1, "lar", "ant")
	s1[0] = "xxx"
	fmt.Println(s1, s2, names)
	fmt.Println("--------------------------")
}

func test2() {
	names := make([]string, 3, 6)
	names[0] = "san"
	names[1] = "ksu"
	names[2] = "fed"
	appNames := append(names, "ant", "lar")
	fmt.Println(names, appNames)

	appNames[0] = "XXX"
	fmt.Println(names, appNames)
	fmt.Println("--------------------------")
}

func test3() {
	names := [3]string{"ksu", "ale", "fed"}
	s1 := names[:]
	s1 = append(s1, "lar")
	s1[0] = "XXX"
	fmt.Println(names, s1)

	names2 := [3]string{"ksu", "san", "alex"}
	s2 := names2[:2]
	fmt.Println(names2)
	s2 = append(s2, "san")
	fmt.Println(names2, s2)
	fmt.Println("--------------------------")

	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println(someNames, allNames)
	someNames = append(someNames, "Gloves")
	fmt.Println(someNames, allNames)
	someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	fmt.Println("--------------------------")
}

func capTest1() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3:3]
	fmt.Println("someNames:", someNames, len(someNames), cap(someNames))
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	//someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	fmt.Println("--------------------------")
}

func sliceFromSlice() {
	names := [4]string{"ksu", "san", "fed", "ant"}
	fmt.Println(names)
	sl1 := names[1:]
	fmt.Println(sl1, len(sl1), cap(sl1))
	sl2 := sl1[1:3]
	fmt.Println(sl2, len(sl2), cap(sl2))

	sl2[0] = "ekat"
	fmt.Println(names, sl1, sl2)
	fmt.Println("--------------------------")
}

func copyTest() {
	names := [4]string{"ksu", "san", "fed", "ant"}
	allNames := names[:]
	someNames := make([]string, 2)
	copy(someNames, allNames)
	fmt.Println(someNames, len(someNames), cap(someNames))

	products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    allNames1 := products[1:] // [Lifejacket Paddle Hat] 3 3
    someNames1 := []string { "Boots", "Canoe"} // [Boots Canoe] 2 2
    copy(someNames1[1:], allNames1[2:3])   // [Canoe Hat]
	fmt.Println("--------------------------")
}

func sortTest() {
	products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    sort.Strings(products)
    for index, value := range products {
        fmt.Println("Index:", index, "Value:", value)
    }
	fmt.Println("--------------------------")
}

func getBaseArray() {
	names := make([]string, 4, 6)
	names[0] = "ksu"
	arrayPtr := (*[4]string)(names)
	fmt.Println(*arrayPtr)
	fmt.Println(names)
	fmt.Println("--------------------------")
}

func map1() {
	products := make(map[string]float64, 10)
    products["Kayak"] = 279
    products["Lifejacket"] = 48.95
    fmt.Println("Map size:", len(products))
    fmt.Println("Price:", products["Kayak"])
    fmt.Println("Price:", products["Hat"])
	fmt.Println("--------------------------")
}

func stringsAreRune() {
	// s := "Привет меня зовут Natasha"
	// fmt.Println(len(s))
	// sRune := []rune(s)
	// fmt.Println(sRune)
	// for i, value := range sRune {
	// 	fmt.Printf("Index: %d, value: %s\n", i, string(value))
	// }
	// fmt.Println("--------------------------")
	// for i, value := range s {
	// 	fmt.Printf("Index: %d, value: %s\n", i, string(value))
	// }
	// bs := []byte(s)
	// for i := range bs {
	// 	bs[i] += uint8(rand.Intn(2)) 
	// }
	// fmt.Printf("%s\n", bs)
	fmt.Println("--------------------------")
}

func test4() {
	var word string
    fmt.Scan(&word)
    wordR := []rune(word)
    reverse := []rune{}
    for i := len(wordR) - 1; i > 0; i-- {
        reverse = append(reverse, wordR[i])
    }
	fmt.Println(string(reverse), wordR)
    if string(wordR) == string(reverse) {
        fmt.Println("Палиндром")
    } else {
        fmt.Println("Нет")
    }
	fmt.Println("--------------------------")
}

func unic() {
	s := []rune("09AZaz")
	fmt.Println(s)
}