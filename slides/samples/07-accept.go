package main

import (
	"net/http"
	"strings"
)

// detectAccept checks the headers of the request and returns the appropriate
// content type or "text/plain" as default
func detectAccept(r *http.Request) string {
	requested := r.Header.Get("Accept")
	accepts := strings.Split(requested, ",")
	switch accepts[0] {
	case "application/json":
		return "application/json"
	case "application/xml":
		return "application/xml"
	default:
		return "text/plain"
	}
}
