package main

import (
	"fmt"
	"io"
	"net/http"
)

func httpError(w io.Writer, err string, status int) {
	msg := fmt.Sprintf("Oops! There was an error:\n %s", err)
	http.Error(w, msg, status)
	return
}
