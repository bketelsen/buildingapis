//go:generate goagen bootstrap -d github.com/bketelsen/buildingapis/exercises/21-goa/solution/design

package main

import (
	"github.com/bketelsen/buildingapis/exercises/21-goa/solution/app"
	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create a new data base and load fixtures
	db := library.NewDB()

	// Create service
	service := goa.New("GoWorkshop")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "course" controller
	c := NewCourseController(service, db)
	app.MountCourseController(service, c)
	// Mount "registration" controller
	c2 := NewRegistrationController(service, db)
	app.MountRegistrationController(service, c2)
	// Mount "public" controller
	c3 := NewPublicController(service)
	app.MountPublicController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
