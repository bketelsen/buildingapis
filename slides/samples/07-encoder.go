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
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
