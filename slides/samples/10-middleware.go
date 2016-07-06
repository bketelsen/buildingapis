package main

import (
	"log"
	"net/http"
	"time"
)

// START OMIT
func simplelogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s", r.URL)
		next.ServeHTTP(w, r)
	})
}

// END OMIT

func durationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		dur := time.Since(now)
		log.Printf("Request took: %v", dur)
	})
}
