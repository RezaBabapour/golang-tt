package render

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t := template.New("home")
	t, _ = template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}

}
