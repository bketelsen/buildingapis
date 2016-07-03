package public

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// RegistrationMedia is the media type used to render registrations.
var RegistrationMedia = MediaType("application/vnd.goworkshop.registration+json", func() {
	Description("RegistrationMedia is the media type used to render registrations")
	TypeName("RegistrationMedia")
	Reference(RegistrationPayload)

	Attributes(func() {
		Attribute("id", Integer, "Registration identifier")
		Attribute("href", String, "Registration href")
		Attribute("course", CourseMedia, "Course being taught")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("address")
		Link("course")
		Required("id", "href", "course", "first_name", "last_name", "address")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("address")
		Attribute("links")
	})

	View("extended", func() {
		Attribute("id")
		Attribute("href")
		Attribute("course")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("address")
		Attribute("links")
	})
})

// RegistrationPayload is the type used to create registrations.
var RegistrationPayload = Type("RegistrationPayload", func() {
	Description("RegistrationPayload is the type used to create registrations")
	Reference(RegistrationPatchPayload)
	Attribute("course_href", String, "The href to the course resource that describes the course being taught", func() {
		Pattern("/courses/[0-9]+")
	})
	Attribute("first_name")
	Attribute("last_name")
	Attribute("address")
	Required("course_href", "first_name", "last_name", "address")
})

// RegistrationPatchPayload is the type used to patch registrations.
var RegistrationPatchPayload = Type("RegistrationPatchPayload", func() {
	Description("RegistrationPatchPayload is the type used to patch registrations")
	Attribute("first_name", String, "Attendee first name", func() {
		MinLength(2)
	})
	Attribute("last_name", String, "Attendee last name", func() {
		MinLength(2)
	})
	Attribute("address", Address, "Attendee address")
})

// Address is the type used to represent street addresses.
var Address = Type("Address", func() {
	Description("Address is a street address")
	Attribute("number", Integer, "Street number", func() {
		Minimum(1)
	})
	Attribute("street", String, "Street name", func() {
		MinLength(1)
	})
	Attribute("city", String, "City", func() {
		MinLength(1)
	})
	Attribute("state", String, "US State Code", func() {
		MinLength(2)
		MaxLength(2)
	})
	Attribute("zip", Integer, "US Zip code", func() {
		Minimum(10000)
		Maximum(99999)
	})
	Required("number", "street", "city", "state", "zip")
})
