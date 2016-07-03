//************************************************************************//
// API "GoWorkshop": Application Media Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/bketelsen/buildingapis/exercises/20-goa/solution/design
// --out=$(GOPATH)/src/github.com/bketelsen/buildingapis/exercises/20-goa/solution
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// CourseMedia media type.
//
// Identifier: application/vnd.goworkshop.course+json
type CourseMedia struct {
	// Course description
	Description *string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	// Course end date/time
	EndTime time.Time `json:"end_time" xml:"end_time" form:"end_time"`
	// Course href
	Href string `json:"href" xml:"href" form:"href"`
	// Course identifier
	ID int `json:"id" xml:"id" form:"id"`
	// Course location
	Location string `json:"location" xml:"location" form:"location"`
	// Course name
	Name string `json:"name" xml:"name" form:"name"`
	// Course start date/time
	StartTime time.Time `json:"start_time" xml:"start_time" form:"start_time"`
}

// Validate validates the CourseMedia media type instance.
func (mt *CourseMedia) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if len(mt.Location) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.location`, mt.Location, len(mt.Location), 2, true))
	}
	if len(mt.Name) < 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 3, true))
	}
	return
}

// CourseMediaLink media type.
//
// Identifier: application/vnd.goworkshop.course+json
type CourseMediaLink struct {
	// Course href
	Href string `json:"href" xml:"href" form:"href"`
	// Course identifier
	ID int `json:"id" xml:"id" form:"id"`
}

// Validate validates the CourseMediaLink media type instance.
func (mt *CourseMediaLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}

// CourseMediaCollection media type is a collection of CourseMedia.
//
// Identifier: application/vnd.goworkshop.course+json; type=collection
type CourseMediaCollection []*CourseMedia

// Validate validates the CourseMediaCollection media type instance.
func (mt CourseMediaCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.Location == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "location"))
		}

		if len(e.Location) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].location`, e.Location, len(e.Location), 2, true))
		}
		if len(e.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 3, true))
		}
	}
	return
}

// CourseMediaLinkCollection media type is a collection of CourseMediaLink.
//
// Identifier: application/vnd.goworkshop.course+json; type=collection
type CourseMediaLinkCollection []*CourseMediaLink

// Validate validates the CourseMediaLinkCollection media type instance.
func (mt CourseMediaLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}

	}
	return
}

// RegistrationMedia media type.
//
// Identifier: application/vnd.goworkshop.registration+json
type RegistrationMedia struct {
	// Attendee address
	Address *Address `json:"address" xml:"address" form:"address"`
	// Attendee first name
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	// Registration href
	Href string `json:"href" xml:"href" form:"href"`
	// Registration identifier
	ID int `json:"id" xml:"id" form:"id"`
	// Attendee last name
	LastName string `json:"last_name" xml:"last_name" form:"last_name"`
	// Links to related resources
	Links *RegistrationMediaLinks `json:"links,omitempty" xml:"links,omitempty" form:"links,omitempty"`
}

// Validate validates the RegistrationMedia media type instance.
func (mt *RegistrationMedia) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if mt.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if mt.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	if mt.Address != nil {
		if err2 := mt.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if len(mt.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, mt.FirstName, len(mt.FirstName), 2, true))
	}
	if len(mt.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, mt.LastName, len(mt.LastName), 2, true))
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// RegistrationMediaExtended media type.
//
// Identifier: application/vnd.goworkshop.registration+json
type RegistrationMediaExtended struct {
	// Attendee address
	Address *Address `json:"address" xml:"address" form:"address"`
	// Course being taught
	Course *CourseMedia `json:"course" xml:"course" form:"course"`
	// Attendee first name
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	// Registration href
	Href string `json:"href" xml:"href" form:"href"`
	// Registration identifier
	ID int `json:"id" xml:"id" form:"id"`
	// Attendee last name
	LastName string `json:"last_name" xml:"last_name" form:"last_name"`
	// Links to related resources
	Links *RegistrationMediaLinks `json:"links,omitempty" xml:"links,omitempty" form:"links,omitempty"`
}

// Validate validates the RegistrationMediaExtended media type instance.
func (mt *RegistrationMediaExtended) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Course == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "course"))
	}
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if mt.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if mt.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	if mt.Address != nil {
		if err2 := mt.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Course != nil {
		if err2 := mt.Course.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if len(mt.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, mt.FirstName, len(mt.FirstName), 2, true))
	}
	if len(mt.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, mt.LastName, len(mt.LastName), 2, true))
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// RegistrationMediaLinks contains links to related resources of RegistrationMedia.
type RegistrationMediaLinks struct {
	Course *CourseMediaLink `json:"course,omitempty" xml:"course,omitempty" form:"course,omitempty"`
}

// Validate validates the RegistrationMediaLinks type instance.
func (ut *RegistrationMediaLinks) Validate() (err error) {
	if ut.Course != nil {
		if err2 := ut.Course.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// RegistrationMediaCollection media type is a collection of RegistrationMedia.
//
// Identifier: application/vnd.goworkshop.registration+json; type=collection
type RegistrationMediaCollection []*RegistrationMedia

// Validate validates the RegistrationMediaCollection media type instance.
func (mt RegistrationMediaCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.FirstName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "first_name"))
		}
		if e.LastName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "last_name"))
		}
		if e.Address == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "address"))
		}

		if e.Address != nil {
			if err2 := e.Address.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if len(e.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].first_name`, e.FirstName, len(e.FirstName), 2, true))
		}
		if len(e.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].last_name`, e.LastName, len(e.LastName), 2, true))
		}
		if e.Links != nil {
			if err2 := e.Links.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// RegistrationMediaExtendedCollection media type is a collection of RegistrationMediaExtended.
//
// Identifier: application/vnd.goworkshop.registration+json; type=collection
type RegistrationMediaExtendedCollection []*RegistrationMediaExtended

// Validate validates the RegistrationMediaExtendedCollection media type instance.
func (mt RegistrationMediaExtendedCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Course == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "course"))
		}
		if e.FirstName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "first_name"))
		}
		if e.LastName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "last_name"))
		}
		if e.Address == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "address"))
		}

		if e.Address != nil {
			if err2 := e.Address.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if e.Course != nil {
			if err2 := e.Course.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if len(e.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].first_name`, e.FirstName, len(e.FirstName), 2, true))
		}
		if len(e.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].last_name`, e.LastName, len(e.LastName), 2, true))
		}
		if e.Links != nil {
			if err2 := e.Links.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// RegistrationMediaLinksArray contains links to related resources of RegistrationMediaCollection.
type RegistrationMediaLinksArray []*RegistrationMediaLinks

// Validate validates the RegistrationMediaLinksArray type instance.
func (ut RegistrationMediaLinksArray) Validate() (err error) {
	for _, e := range ut {
		if e.Course != nil {
			if err2 := e.Course.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
