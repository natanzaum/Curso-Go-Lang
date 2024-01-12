package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	u := usuario{"Tamara Pereira", "tamarindo@gmail.com"}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
