package main

import (
	"log"
	"net/http"
	"time"
)

func simplelogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s", r.URL)
		next.ServeHTTP(w, r) // CALLING next HERE
	})
}

// START OMIT
func durationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// code executed before 'next' handler is here
		now := time.Now()
		next.ServeHTTP(w, r) // next is called HERE
		// code executed after 'next' handler is here
		dur := time.Since(now)
		log.Printf("Request took: %v", dur)
	})
}

// END OMIT
