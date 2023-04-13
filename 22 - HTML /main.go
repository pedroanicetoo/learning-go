package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := user{
		"Peter",
		"peter@parker.com",
	}

	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Listen on Port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
