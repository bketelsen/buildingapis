//go:generate goagen bootstrap -d github.com/gophercon/buildingapis/workshop/18-goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/gophercon/buildingapis/workshop/18-goa/app"
	"github.com/inconshreveable/log15"
)

func main() {
	// Create a new data base and load fixtures
	db := NewDB()

	// Instantiate logger
	logger := log15.New()
	logger.SetHandler(log15.StdoutHandler)

	// Create service
	service := NewService(db)
	service.WithLogger(goalog15.New(logger))

	// Mount "course" controller
	c := NewCourseController(service, db)
	app.MountCourseController(service, c)
	// Mount "public" controller
	c2 := NewPublicController(service)
	app.MountPublicController(service, c2)
	// Mount "registration" controller
	c3 := NewRegistrationController(service, db)
	app.MountRegistrationController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

// NewService instantiates and configures a new GoWorkshop service.
func NewService(db *MemDB) *goa.Service {
	// Create service
	service := goa.New("GoWorkshop")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middleware
	app.UseBasicAuthMiddleware(service, basicAuth(db))
	app.UseJWTAuthMiddleware(service, jwtAuth())

	return service
}
