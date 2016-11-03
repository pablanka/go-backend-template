package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pablanka/go-backend-template/router"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(*addr, r))
}
