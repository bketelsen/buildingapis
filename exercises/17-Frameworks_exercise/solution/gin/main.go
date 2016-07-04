package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

func main() {
	const (
		host = "localhost:8000"
	)

	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stderr, log15.TerminalFormat()))

	r := gin.Default()
	r.Use(requestIDMiddleware())
	r.Use(loggerMiddleware(logger))

	db := library.NewDB()

	r.POST("/courses", createCourse(db))
	r.GET("/courses/:id", showCourse(db))
	r.DELETE("/courses/:id", deleteCourse(db))

	r.POST("/registrations", createRegistration(db))
	r.GET("/registrations/:id", showRegistration(db))
	r.GET("/registrations", listRegistrations(db))

	logger.Info("ready", "host", host)
	r.Run(host)
}

// respondBadRequest writes a response with status 400 and the body built with the given format
// and values (a la fmt).
func respondBadRequest(c *gin.Context, format string, vals ...interface{}) {
	// NOTE: this would need to be sanitized or entirely removed before running in production.
	body := fmt.Sprintf(format, vals...)
	c.String(http.StatusBadRequest, body)
}

// respondInternal writes a response with status 500 and the body built with the given format and
// values (a la fmt).
func respondInternal(c *gin.Context, format string, vals ...interface{}) {
	// NOTE: this would need to be sanitized or entirely removed before running in production.
	body := fmt.Sprintf(format, vals...)
	c.String(http.StatusInternalServerError, body)
}

// respondNotFound writes a response with status 404.
func respondNotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}
