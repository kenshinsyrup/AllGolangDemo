package main

import "html/template"
import "net/http"

var testTemplate *template.Template

type User struct {
	// Admin bool
	ID            int
	Email         string
	HasPermission func(string) bool
}

type ViewData struct {
	User User
}

// func (u User) HasPermission(feature string) bool {
// 	if feature == "feature-a" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func main() {
	var err error
	testTemplate, err = template.New("hello6.html").Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			return false
		},
		"ifIE": func() template.HTML {
			return template.HTML("<!--[if IE]>")
		},
		"endif": func() template.HTML {
			return template.HTML("<![endif]-->")
		},
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	}).ParseFiles("hello6.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	user := User{
		ID:    1,
		Email: "j@c.io",
	}
	vd := ViewData{user}

	err := testTemplate.Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			if user.ID == 1 && feature == "feature-a" {
				return true
			}
			return false
		},
	}).Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
