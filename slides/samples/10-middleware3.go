package main

import "net/http"

// START OMIT
func secretSquirrelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vals, ok := r.Header["X-SECRET-SQUIRREL"] // map[string][]string
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var found bool
		for _, h := range vals { // iterate header contents
			if h == "Gophers" {
				found = true
			}
		}
		if !found {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r) // next is called HERE
	})
}

// END OMIT
