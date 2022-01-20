package main

import (
	"log"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t := template.New("home")
	t, _ = template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}

}
