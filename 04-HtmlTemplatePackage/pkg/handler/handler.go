package handler

import (
	"htmlTemplate/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")
	// renderTemplate(w, "home")
}
