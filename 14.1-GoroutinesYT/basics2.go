package main

import (
	"fmt"
)

type Resume struct {
	Language 		  string
	YearsOfExperience int8
	YearsOfEducation  int8
}

type Vacancy struct {
	Language		  string
	Company			  string
	YearsOfExperience int8
	YearsOfEducation  int8
}

func main() {
	fmt.Println("Rus")
}