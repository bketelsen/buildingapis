package public

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CourseMedia is the media type used to render courses.
var CourseMedia = MediaType("application/vnd.goworkshop.course+json", func() {
	Description("CourseMedia is the media type used to render courses")
	TypeName("CourseMedia")
	Reference(CoursePayload)

	Attributes(func() {
		Attribute("id", Integer, "Course identifier")
		Attribute("href", String, "Course href")
		Attribute("name")
		Attribute("description")
		Attribute("start_time")
		Attribute("end_time")
		Attribute("location")
		Required("id", "href", "name", "start_time", "end_time", "location")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("description")
		Attribute("start_time")
		Attribute("end_time")
		Attribute("location")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

// CoursePayload is the type used to create courses.
var CoursePayload = Type("CoursePatchPayload", func() {
	Description("CoursePayload is the type used to create courses")
	Attribute("name", String, "Course name", func() {
		MinLength(3)
	})
	Attribute("description", String, "Course description")
	Attribute("start_time", DateTime, "Course start date/time")
	Attribute("end_time", DateTime, "Course end date/time")
	Attribute("location", String, "Course location", func() {
		MinLength(2)
	})
	Required("name", "start_time", "end_time", "location")
})
