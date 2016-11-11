package main

import (
	"html/template"
	"os"
)

type Test struct {
	HTML     string
	SafeHTML template.HTML
	Title    string
	Path     string
	Dog      Dog
	Map      map[string]string
}

type Dog struct {
	Name string
	Age  int
}

func main() {
	// 1
	// t, err := template.ParseFiles("hello.html")
	// if err != nil {
	// 	panic(err)
	// }

	// data := struct {
	// 	Name string
	// }{
	// 	Name: "John Smith",
	// }

	// err = t.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	// 2
	t, err := template.ParseFiles("context.html")
	if err != nil {
		panic(err)
	}

	data := Test{
		HTML:     "<h1>A header!</h1>",
		SafeHTML: template.HTML("<h1>A Safe header</h1>"),
		Title:    "Backslash! An in depth look at the \"\\\" character.",
		Path:     "/dashboard/settings",
		Dog:      Dog{"Fido", 6},
		Map: map[string]string{
			"key":       "value",
			"other_key": "other_value",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
