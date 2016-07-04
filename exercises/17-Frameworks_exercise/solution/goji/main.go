package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/inconshreveable/log15"

	"goji.io"
	"goji.io/pat"
)

func main() {
	const (
		host = "localhost:8000"
	)

	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stderr, log15.TerminalFormat()))

	mux := goji.NewMux()
	mux.UseC(requestIDMiddleware())
	mux.UseC(loggerMiddleware(logger))

	db := library.NewDB()

	mux.HandleFuncC(pat.Post("/courses"), createCourse(db))
	mux.HandleFuncC(pat.Get("/courses/:id"), showCourse(db))
	mux.HandleFuncC(pat.Delete("/courses/:id"), deleteCourse(db))

	mux.HandleFuncC(pat.Post("/registrations"), createRegistration(db))
	mux.HandleFuncC(pat.Get("/registrations/:id"), showRegistration(db))
	mux.HandleFuncC(pat.Get("/registrations"), listRegistrations(db))

	logger.Info("ready", "host", host)
	http.ListenAndServe(host, mux)
}

// respondBadRequest writes a response with status 400 and the body built with the given format
// and values (a la fmt).
func respondBadRequest(w http.ResponseWriter, format string, vals ...interface{}) {
	// NOTE: this would need to be sanitized or entirely removed before running in production.
	body := fmt.Sprintf(format, vals...)
	w.WriteHeader(400)
	w.Write([]byte(body))
}

// respondInternal writes a response with status 500 and the body built with the given format and
// values (a la fmt).
func respondInternal(w http.ResponseWriter, format string, vals ...interface{}) {
	// NOTE: this would need to be sanitized or entirely removed before running in production.
	body := fmt.Sprintf(format, vals...)
	w.WriteHeader(500)
	w.Write([]byte(body))
}

// respondNotFound writes a response with status 404.
func respondNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
}
