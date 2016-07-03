// +build appengine // HL

package main

import (
	"net/http"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/inconshreveable/log15"
)

func init() {
	goa.Log.SetHandler(log15.StreamHandler(os.Stderr, log15.LogfmtFormat())) // HL

	service := goa.New("cellar")
	ac := NewAccountController(service)
	app.MountAccountController(service, ac)

	http.HandleFunc("/", service.HTTPHandler().ServeHTTP) // HL
}
