package main

import (
	"allgolangdemo/templatego/src/templateapp4/views"
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
	contact.Render(w, nil)
}
