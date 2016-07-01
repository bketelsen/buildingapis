package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/gophercon/buildingapis/workshop/18-goa/app"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const (
	// FailAuthMessage is the message returned when basic authentication fails
	FailAuthMessage = "Email not found or password does not match"

	// key used to store basic auth user in context
	userKey int = iota
)

// basicAuth is the basic auth middleware.
func basicAuth(db *MemDB) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve basic auth email and password
			email, pass, ok := req.BasicAuth()
			if !ok {
				return goa.ErrUnauthorized("No auth")
			}

			// Lookup user in database
			i, err := db.Get("users", "email", email)
			if err != nil {
				return goa.ErrUnauthorized(FailAuthMessage)
			}
			user := i.(*UserModel)

			// Compare passwords
			if bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.HashedPassword)) != nil {
				return goa.ErrUnauthorized(FailAuthMessage)
			}

			// Auth succeeded, proceed
			ctx = context.WithValue(ctx, userKey, user)
			return h(ctx, rw, req)
		}
	}
}

// contextUser extracts the user set by basic auth from the context.
func contextUser(ctx context.Context) *UserModel {
	if i := ctx.Value(userKey); i != nil {
		return i.(*UserModel)
	}
	return nil
}

// jwtAuth is the JWT auth middleware.
func jwtAuth() goa.Middleware {
	keys, err := loadJWTPublicKeys()
	if err != nil {
		panic(err)
	}
	return jwt.New(keys, nil, app.NewJWTAuthSecurity())
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func loadJWTPublicKeys() ([]*rsa.PublicKey, error) {
	keyFiles, err := filepath.Glob("./keys/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]*rsa.PublicKey, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}
