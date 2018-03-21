package main

import (
	"log"
	"net/http"

	"github.com/ajstarks/svgo"
)

func main() {
	/*
		width := 500
		height := 500
		canvas := svg.New(os.Stdout)
		canvas.Start(width, height)
		canvas.Circle(width/2, height/2, 100)
		canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
		canvas.End()
	*/
	http.Handle("/circle", http.HandlerFunc(circle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	width := 500
	height := 500
	s := svg.New(w)
	s.Start(width, height)
	s.Circle(width/2, height/2, 125, "fill:none;stroke:black")
	s.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:red")

	s.End()
}
