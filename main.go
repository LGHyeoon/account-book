package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles("resources/views/login.html"))
		tmp1.Execute(w, nil)
	}

	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
