package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bketelsen/buildingapis/exercises/library"
)

func main() {

	db := library.NewDB()
	cs := &CourseServer{
		DB: db,
	}
	// START OMIT
	http.Handle(courseBase, durationMiddleware(simplelogMiddleware(cs)))
	// END OMIT
	http.HandleFunc(registrationBase, registrations)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func simplelogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s", r.URL)
		next.ServeHTTP(w, r)
	})
}

func durationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		dur := time.Since(now)
		log.Printf("Request took: %v", dur)
	})
}
