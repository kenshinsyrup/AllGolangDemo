package main

import (
	"allgolangdemo/templatego/src/templateapp4/views"
	"fmt"
	"net/http"
)

var index *views.View
var contact *views.View

func main() {
	index = views.NewView("bootstrap", "views/index.html")
	contact = views.NewView("bootstrap", "views/contact.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Render(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		contact.Render(w, nil)
	}
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
