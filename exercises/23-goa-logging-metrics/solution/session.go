package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bketelsen/buildingapis/exercises/20-goa/solution/app"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/satori/go.uuid"
)

// SessionController implements the session resource.
type SessionController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewSessionController creates a session controller.
func NewSessionController(service *goa.Service) *SessionController {
	// Load JWT private jey
	b, err := ioutil.ReadFile("./keys/jwt.key")
	if err != nil {
		panic(err)
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(fmt.Sprintf("jwt: failed to load private key: %s", err))
	}
	return &SessionController{
		Controller: service.NewController("SessionController"),
		privateKey: privKey,
	}
}

// Login is called after basic auth has succeeded.
// It creates a JWT and returns it in the response Authorization header.
func (c *SessionController) Login(ctx *app.LoginSessionContext) error {
	// Generate JWT
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	token.Claims = jwtgo.StandardClaims{
		Issuer:    "GoWorkshop",           // who creates the token and signs it
		Audience:  contextUser(ctx).Email, // to whom the token is intended to be sent
		ExpiresAt: in10m,                  // time when the token will expire (10 minutes from now)
		Id:        uuid.NewV4().String(),  // a unique identifier for the token
		IssuedAt:  time.Now().Unix(),      // when the token was issued/created (now)
		NotBefore: 2,                      // time before which the token is not yet valid (2 minutes ago)
		Subject:   "user",                 // the subject/principal is whom the token is about
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)
	return nil
}
