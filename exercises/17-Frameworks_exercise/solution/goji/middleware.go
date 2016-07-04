package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/inconshreveable/log15"

	"goji.io"
	"golang.org/x/net/context"
)

// middlewareKey is the private type used for iddlewares to store values in the context.
// It is private to avoid possible collisions with keys used by other packages.
type middlewareKey int

// reqIDKey is the context key used by the RequestID middleware to store the request ID value.
const reqIDKey middlewareKey = 1

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

func requestIDMiddleware() func(goji.Handler) goji.Handler {
	return func(h goji.Handler) goji.Handler {
		m := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get(RequestIDHeader)
			if id == "" {
				id = fmt.Sprintf("%s-%d", reqPrefix, atomic.AddInt64(&reqID, 1))
			}
			ctx = context.WithValue(ctx, reqIDKey, id)
			h.ServeHTTPC(ctx, w, r)
		}
		return goji.HandlerFunc(m)
	}
}

func loggerMiddleware(logger log15.Logger) func(goji.Handler) goji.Handler {
	return func(h goji.Handler) goji.Handler {
		m := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			reqID := ctx.Value(reqIDKey)
			logger.Info("started", "req_id", reqID, r.Method, r.URL.String())
			defer func(now time.Time) {
				logger.Info("completed", "req_id", reqID, "time", time.Since(now).String())
			}(time.Now())
			h.ServeHTTPC(ctx, w, r)
		}
		return goji.HandlerFunc(m)
	}
}
