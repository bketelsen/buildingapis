package main

import (
	"fmt"
	"julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func main() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
