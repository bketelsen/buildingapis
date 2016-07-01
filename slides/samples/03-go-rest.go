package main

import "net/http"

func main() {

	m := http.NewServeMux()
	m.Handle("/api/courses", http.HandlerFunc(serveCourses))
	m.Handle("/api/registrations", http.HandlerFunc(serveRegistrations))
	m.Handle("/api/token", http.HandlerFunc(serveTokens))
}
