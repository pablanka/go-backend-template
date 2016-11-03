package main

import (
	"log"
	"net/http"

	"github.org/pablanka/go-backend-template/router"
)

func main() {
	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
