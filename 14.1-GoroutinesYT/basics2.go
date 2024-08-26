package main

import (
	"fmt"
	"time"
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
	resume := &Resume{
		Language: "Go",
		YearsOfExperience: 2,
		YearsOfEducation: 5,
	}

	vacancies := []*Vacancy{
		{"Go", "Nike", 4, 6},
		{"Java", "Yandex", 1, 2},
		{"Go", "Sber", 2, 1},
	}

}

func applyForVacanciesFn(r *Resume, v []*Vacancy, applyFunc func(r *Resume, v []*Vacancy) (invite []*Vacancy)) {
	var invite []*Vacancy
	timeStart := time.Now()

	invite = applyFunc(r, v)
	fmt.Printf("ВАС ПРИГЛАСИЛИ ")
} 