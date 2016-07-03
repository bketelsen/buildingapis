package main

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	Name     string
	Category string
}

func serveInfo(w http.ResponseWriter, r *http.Request) {
	info := Info{Name: "Bobs Uncle", Category: "Relatives"}
	js, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
} // implicit return 200-ok
