package handlers

import (
	"net/http"

	"github.com/pablanka/go-backend-template/templates"
)

// Index handles welcome page
func Index(w http.ResponseWriter, r *http.Request) {
	templates.ParseNExecute(w, map[string]string{"Name": "go-backend-template"}, "./templates/template.html")
}

// Webapp handles demo web app
func Webapp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./statics/webapp/angular-demo.html")
}
