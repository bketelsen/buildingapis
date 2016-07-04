package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

// reqIDKey is the context key used by the RequestID middleware to store the request ID value.
const reqIDKey = "reqIDKey"

// RequestIDHeader is the name of the header used to transmit the request ID.
const RequestIDHeader = "X-Request-Id"

// Counter used to create new request ids.
var reqID int64

// Common prefix to all newly created request ids for this process.
var reqPrefix string

// Initialize common prefix on process startup.
func init() {
	// algorithm taken from
	// https://github.com/zenazn/goji/blob/master/web/middleware/request_id.go#L44-L50
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}
	reqPrefix = string(b64[0:10])
}

func requestIDMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Request.Header.Get(RequestIDHeader)
		if id == "" {
			id = fmt.Sprintf("%s-%d", reqPrefix, atomic.AddInt64(&reqID, 1))
		}
		c.Set(reqIDKey, id)
		c.Next()
	}
}

func loggerMiddleware(logger log15.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		reqID := c.Value(reqIDKey)
		logger.Info("started", "req_id", reqID, c.Request.Method, c.Request.URL.String())
		defer func(now time.Time) {
			logger.Info("completed", "req_id", reqID, "time", time.Since(now).String())
		}(time.Now())
		c.Next()
	}
}
