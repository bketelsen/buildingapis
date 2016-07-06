package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	n := negroni.New(Middleware1, Middleware2)
	// Or use a middleware with the Use() function
	n.Use(Middleware3)
	// router goes last
	n.UseHandler(router)
	http.ListenAndServe(":3001", n)
}

// END OMIT
