package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		time.Sleep(time.Second * 2)
		log.Println("hhahhaha 过一会儿再显示反正即使messageHandler return了， main进程还在，go routine不会被吃掉的")
	}()
	log.Println("我会出现在hahahha那一堆之前")
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func main() {
	mux := http.NewServeMux()
	// Use the shortcut method ServeMux.HandleFunc
	mux.HandleFunc("/welcome", messageHandler)
	log.Println("Listening...")
	http.ListenAndServe(":9998", mux)
}
