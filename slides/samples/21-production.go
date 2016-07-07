package main

import "log"

type APIError struct {
	Code    int
	Message string
}

var ErrNotFound = APIError{Code: 404, Message: "Not Found"}

func (a APIError) Error() {
	return a.Message
}

func logError(err APIError, location string) {
	log.Println("Error: %s at %s", err, location)
	metrics.Increment(err.Code)
}
