//go:generate goagen bootstrap -d github.com/bketelsen/buildingapis/exercises/21-goa/solution/design

package main

import (
	"log/syslog"
	"os"

	"github.com/armon/go-metrics"
	"github.com/bketelsen/buildingapis/exercises/21-goa/solution/app"
	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
)

func main() {
	// Create a new data base and load fixtures
	db := library.NewDB()

	// Create service
	service := goa.New("GoWorkshop")

	// Setup logger
	logger := log15.New()
	syshandler, err := log15.SyslogHandler(syslog.LOG_INFO|syslog.LOG_LOCAL0, "", log15.LogfmtFormat())
	if err != nil {
		panic(err)
	}
	logger.SetHandler(log15.MultiHandler(
		log15.StreamHandler(os.Stdout, log15.TerminalFormat()),
		syshandler,
	))
	service.WithLogger(goalog15.New(logger))

	// Setup metrics
	sink, err := metrics.NewStatsdSink("localhost:8125")
	if err != nil {
		panic(err.Error())
	}
	if err := goa.NewMetrics(metrics.DefaultConfig("GoWorkshop"), sink); err != nil {
		panic(err.Error())
	}

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
