package main

import (
	"io/ioutil"
	"testing"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

func newService() *goa.Service {
	return goa.New("test")
}

// authorizedContext creates a JWT token using the private key found in keys/jwt.key and creates
// a context initialized with it.
func authorizedContext(t *testing.T) context.Context {
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	token.Claims = jwtgo.StandardClaims{
		Issuer:    "Issuer",
		Audience:  "Audience",
		ExpiresAt: in10m,
		Id:        uuid.NewV4().String(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: 2,
		Subject:   "subject",
	}
	b, err := ioutil.ReadFile("./keys/jwt.key")
	if err != nil {
		t.Fatalf("failed to load JWT private key: %s", err)
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		t.Fatalf("failed to parse JWT private key: %s", err)
	}
	signed, err := token.SignedString(privKey)
	if err != nil {
		t.Fatalf("failed to sign token: %s", err)
	}
	pubkeys, err := loadJWTPublicKeys()
	if err != nil {
		t.Fatalf("failed to load JWT public key: %s", err)
	}
	jt, err := jwtgo.Parse(signed, func(*jwtgo.Token) (interface{}, error) { return pubkeys[0], nil })
	if err != nil {
		t.Fatalf("failed to create token: %s", err)
	}
	return jwt.WithJWT(context.Background(), jt)
}
