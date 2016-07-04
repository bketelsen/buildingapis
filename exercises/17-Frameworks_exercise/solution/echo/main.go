package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/inconshreveable/log15"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	const (
		host = "localhost:8000"
	)

	e := echo.New()

	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stderr, log15.TerminalFormat()))

	e.Use(requestIDMiddleware())
	e.Use(loggerMiddleware(logger))

	db := library.NewDB()

	e.POST("/courses", createCourse(db))
	e.GET("/courses/:id", showCourse(db))
	e.DELETE("/courses/:id", deleteCourse(db))

	e.POST("/registrations", createRegistration(db))
	e.GET("/registrations/:id", showRegistration(db))
	e.GET("/registrations", listRegistrations(db))

	logger.Info("ready", "host", host)
	e.Run(standard.New(host))
}

// respondBadRequest writes a response with status 400 and the body built with the given format
// and values (a la fmt).
func respondBadRequest(c echo.Context, format string, vals ...interface{}) error {
	// NOTE: this would need to be sanitized or entirely removed before running in production.
	body := fmt.Sprintf(format, vals...)
	return c.String(http.StatusBadRequest, body)
}

// respondInternal writes a response with status 500 and the body built with the given format and
// values (a la fmt).
func respondInternal(c echo.Context, format string, vals ...interface{}) error {
	// NOTE: this would need to be sanitized or entirely removed before running in production.
	body := fmt.Sprintf(format, vals...)
	return c.String(http.StatusInternalServerError, body)
}

// respondNotFound writes a response with status 404.
func respondNotFound(c echo.Context) error {
	return c.NoContent(http.StatusNotFound)
}
