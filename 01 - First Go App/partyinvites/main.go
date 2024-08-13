package main

import (
	"fmt"
	"html/template"
)

type Rsvp struct {	// Структура ответ
	Name, Email, Phone string
	WillAttend bool
}

var responses = make([]*Rsvp, 0, 10)  // срез указателей на экземпляры структуры Rsvp для хранения
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name + ".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main() {
	loadTemplates()
}