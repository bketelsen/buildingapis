package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// START OMIT
func main() {
	r := mux.NewRouter()
	canonical := handlers.CanonicalHost("http://www.gorillatoolkit.org", 302)
	r.HandleFunc("/route", YourHandler)

	log.Fatal(http.ListenAndServe(":7000", canonical(r)))
}

// END OMIT
