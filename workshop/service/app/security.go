//************************************************************************//
// API "GoWorkshop": Application Security
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/bketelsen/buildingapis/workshop/service/design
// --out=$(GOPATH)/src/github.com/bketelsen/buildingapis/workshop/service
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

type (
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseJWTAuthMiddleware mounts the JWTAuth auth middleware onto the service.
func UseJWTAuthMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("JWTAuth"), middleware)
}

// NewJWTAuthSecurity creates a JWTAuth security definition.
func NewJWTAuthSecurity() *goa.JWTSecurity {
	def := goa.JWTSecurity{
		In:       goa.LocHeader,
		Name:     "Authorization",
		TokenURL: "http://localhost:8080/token",
	}
	return &def
}

// UseBasicAuthMiddleware mounts the BasicAuth auth middleware onto the service.
func UseBasicAuthMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("BasicAuth"), middleware)
}

// NewBasicAuthSecurity creates a BasicAuth security definition.
func NewBasicAuthSecurity() *goa.BasicAuthSecurity {
	def := goa.BasicAuthSecurity{}
	def.Description = "User email and password authentication"
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}
