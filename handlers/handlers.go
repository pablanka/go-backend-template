package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pablanka/go-backend-template/models"
	"github.com/pablanka/go-backend-template/templates"
)

func parseNCloseBody(body io.ReadCloser) (interface{}, error) {
	decoder := json.NewDecoder(body)
	var i interface{}
	err := decoder.Decode(&i)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	return i, nil
}

// VIEWS HANDLERS-------------------------------------------------------------------------------------------------

// Index handles welcome page
func Index(w http.ResponseWriter, r *http.Request) {
	templates.ParseNExecute(w, map[string]string{"Name": "go-backend-template"}, "./templates/template.html")
}

/*// Webapp handles demo web app
func Webapp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./statics/webapp/angular-demo.html")
}*/

//Statics statics files entry point
func Statics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	http.ServeFile(w, r, "./statics/"+r.URL.Path)
}

// REST APIs ----------------------------------------------------------------------------------------------------

// GETUser handles get user API
func GETUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		Name: models.Name{
			FirstName: "Pablo",
			LastName:  "Acuña",
		},
		Age:       26,
		Job:       "Programer",
		Seniority: "Senior",
	}
	fmt.Fprintln(w, models.Stringnify(user))
}

// POSTUser handles post user API
func POSTUser(w http.ResponseWriter, r *http.Request) {
	user, _ := parseNCloseBody(r.Body)
	fmt.Fprintln(w, "Posted User: ", models.Stringnify(user))
}
