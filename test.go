package main

import "fmt"

func main() {
    Digitize(12345)
}

func Digitize(n int) []int {
  mySlice := []int{}
  for n > 0 {
    x := n % 10
    n = n / 10
    mySlice = append(mySlice, x)
  }
  fmt.Println(mySlice)
  return mySlice
}