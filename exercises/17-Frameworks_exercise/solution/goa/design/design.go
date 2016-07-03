package design

import (
	. "github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa/design/public"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This block defines the global properties of the service API.
var _ = API("GoWorkshop", func() {

	// General metadata about the service
	Title("The Universal Workshop Service")
	Description("GoWorkshop is a simple example service that exposes a REST API")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("course", func() {

	Description("The course resource exposes the endpoints used to manage workshop courses")
	BasePath("/courses")

	// Create course
	Action("create", func() {
		Description("Create a new course")
		Routing(POST("/"))
		Payload(CoursePayload)
		Response(Created, func() {
			Headers(func() {
				Header("Location", String, "Newly created course href", func() {
					Pattern("/registrations/[0-9]+")
				})
			})
			Media(CourseMedia)
		})
		Response(BadRequest, ErrorMedia)
	})

	// Show a course
	Action("show", func() {
		Description("Show a courses")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "The course ID", func() {
				Minimum(1)
			})
		})
		Response(OK, CourseMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	// Delete a course
	Action("delete", func() {
		Description("Delete a course")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "The course ID", func() {
				Minimum(1)
			})
		})
		Response(NoContent)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})
})

// This block defines the "registration" resource endpoints.
var _ = Resource("registration", func() {

	Description("The registration resource exposes the endpoints used to manage workshop registrations")
	BasePath("/registrations")

	// Create registration
	Action("create", func() {
		Description("Create a new registration")
		Routing(POST("/"))
		Payload(RegistrationPayload)
		Response(Created, func() {
			Headers(func() {
				Header("Location", String, "Newly created registration href", func() {
					Pattern("/registrations/[0-9]+")
				})
			})
			Media(RegistrationMedia)
		})
		Response(BadRequest, ErrorMedia)
	})

	// Show a registration
	Action("show", func() {
		Description("Show a registration")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "The registration ID", func() {
				Minimum(1)
			})
		})
		Response(OK, RegistrationMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	// List registrations
	Action("list", func() {
		Description("List all registrations")
		Routing(GET("/"))
		Response(OK, CollectionOf(RegistrationMedia))
	})
})
