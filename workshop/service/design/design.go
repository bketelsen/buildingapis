package design

import (
	. "github.com/bketelsen/buildingapis/workshop/service/design/public"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Define a few constants so we can easily switch between development and production values.
const (
	// Hostname is the service hostname and port in the form "hostname:port".
	// This is mainly used by the documentation, Swagger and client tool code.
	Hostname = "localhost:8080"

	// HTTPScheme is the scheme used by the API. As with Hostname this is mainly
	// used used by the documentation, Swagger and client tool code.
	HTTPScheme = "http"
)

// This block defines the global properties of the service API.
var _ = API("GoWorkshop", func() {

	// General metadata about the service
	Title("The Universal Workshop Service")
	Description("GoWorkshop is a simple example service that exposes a REST API")
	Version("1.0")
	Contact(func() {
		Name("The GoWorkshop developers")
		Email("gw@goa.design")
	})
	License(func() {
		Name("The MIT License (MIT)")
		URL("https://github.com/gophercon/buildingapis/blob/master/LICENSE")
	})

	// Endpoint definition
	Host(Hostname)
	Scheme(HTTPScheme)
	BasePath("/api")

	// Encoding and decoding
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("course", func() {

	Description("The course resource exposes the endpoints used to manage workshop courses")
	BasePath("/courses")
	Security(JWTAuth)

	// Common response to all actions
	Response(Unauthorized, ErrorMedia)

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

	// Retrieve course
	Action("show", func() {
		Description("Retrieve a course by ID")
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

	// List courses
	Action("list", func() {
		Description("List all courses")
		Routing(GET("/"))
		Params(func() {
			Param("name", String, "Filter by name")
		})
		Response(OK, CollectionOf(CourseMedia))
	})

	// Patch a course
	Action("patch", func() {
		Description("Update a course using patch semantic")
		Routing(PATCH("/:id"))
		Params(func() {
			Param("id", Integer, "The course ID", func() {
				Minimum(1)
			})
		})
		Payload(CoursePatchPayload)
		Response(NoContent)
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
		Response(NotFound)
	})

})

// This block defines the "registration" resource endpoints.
var _ = Resource("registration", func() {

	Description("The registration resource exposes the endpoints used to manage workshop registrations")
	BasePath("/registrations")
	Security(JWTAuth)

	// Common response to all actions
	Response(Unauthorized, ErrorMedia)

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

	// Retrieve registration
	Action("show", func() {
		Description("Retrieve a registration by ID")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "The registration ID", func() {
				Minimum(1)
			})
			Param("view", String, "The view used to render the registration", func() {
				Enum("default", "extended")
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
		Params(func() {
			Param("course_id", Integer, "Filter by course", func() {
				Minimum(1)
			})
		})
		Response(OK, CollectionOf(RegistrationMedia))
	})

	// Patch a registration
	Action("patch", func() {
		Description("Update a registration using patch semantic")
		Routing(PATCH("/:id"))
		Params(func() {
			Param("id", Integer, "The registration ID", func() {
				Minimum(1)
			})
		})
		Payload(RegistrationPatchPayload)
		Response(NoContent)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})
})

// This block defines the login endpoint. This endpoint is secured via basic auth.
var _ = Resource("session", func() {

	Description("The session resource exposes the login action")
	Security(BasicAuth)

	// Login
	Action("login", func() {
		Description(`Login uses basic auth and on successful auth returns a JWT in the response "Authorization" header`)
		Routing(GET("/token"))
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "JWT that can be used to make requests to the registration and course endpoints")
			})
		})
		Response(Unauthorized)
	})
})

// This block defines the "swagger" resource which serves static files (swagger definitions)
var _ = Resource("public", func() {

	Description("The public resource groups endpoints that serve static content")

	// Swagger JSON
	Files("/swagger.json", "swagger/swagger.json", func() {
		Description("API Swagger spec in JSON format")
	})

	// Swagger YAML
	Files("/swagger.yaml", "swagger/swagger.yaml", func() {
		Description("API Swagger spec in YAML format")
	})

	// Swagger UI
	Files("/swagger/*file", "public/", func() {
		Description("Swagger UI")
	})
})

// BasicAuth is the security scheme used to secure the JWT token retrieval endpoint.
var BasicAuth = BasicAuthSecurity("BasicAuth", func() {
	Description("User email and password authentication")
})

// JWTAuth is the security scheme used to secure all endpoints - except the JWT token retrieval
// endpoint (you'd have a bit of a chicken and egg issue otherwise).
var JWTAuth = JWTSecurity("JWTAuth", func() {
	Header("Authorization")
	TokenURL(HTTPScheme + "://" + Hostname + "/token")
})
