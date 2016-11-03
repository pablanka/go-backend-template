package templates

import (
	"html/template"
	"net/http"
)

// ParseNExecute parses and executes a given template.
func ParseNExecute(w http.ResponseWriter, data interface{}, templateURLS ...string) error {
	t, error := template.ParseFiles(templateURLS...)
	if error != nil {
		return error
	}
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return t.Execute(w, data)
}
