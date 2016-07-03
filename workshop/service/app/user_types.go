//************************************************************************//
// API "GoWorkshop": Application User Types
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

// Address is a street address
type address struct {
	// City
	City *string `json:"city,omitempty" xml:"city,omitempty" form:"city,omitempty"`
	// Street number
	Number *int `json:"number,omitempty" xml:"number,omitempty" form:"number,omitempty"`
	// US State Code
	State *string `json:"state,omitempty" xml:"state,omitempty" form:"state,omitempty"`
	// Street name
	Street *string `json:"street,omitempty" xml:"street,omitempty" form:"street,omitempty"`
	// US Zip code
	Zip *int `json:"zip,omitempty" xml:"zip,omitempty" form:"zip,omitempty"`
}

// Validate validates the address type instance.
func (ut *address) Validate() (err error) {
	if ut.Number == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "number"))
	}
	if ut.Street == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "street"))
	}
	if ut.City == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "city"))
	}
	if ut.State == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "state"))
	}
	if ut.Zip == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "zip"))
	}

	if ut.City != nil {
		if len(*ut.City) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.city`, *ut.City, len(*ut.City), 1, true))
		}
	}
	if ut.Number != nil {
		if *ut.Number < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.number`, *ut.Number, 1, true))
		}
	}
	if ut.State != nil {
		if len(*ut.State) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.state`, *ut.State, len(*ut.State), 2, true))
		}
	}
	if ut.State != nil {
		if len(*ut.State) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.state`, *ut.State, len(*ut.State), 2, false))
		}
	}
	if ut.Street != nil {
		if len(*ut.Street) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.street`, *ut.Street, len(*ut.Street), 1, true))
		}
	}
	if ut.Zip != nil {
		if *ut.Zip < 10000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.zip`, *ut.Zip, 10000, true))
		}
	}
	if ut.Zip != nil {
		if *ut.Zip > 99999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.zip`, *ut.Zip, 99999, false))
		}
	}
	return
}

// Publicize creates Address from address
func (ut *address) Publicize() *Address {
	var pub Address
	if ut.City != nil {
		pub.City = *ut.City
	}
	if ut.Number != nil {
		pub.Number = *ut.Number
	}
	if ut.State != nil {
		pub.State = *ut.State
	}
	if ut.Street != nil {
		pub.Street = *ut.Street
	}
	if ut.Zip != nil {
		pub.Zip = *ut.Zip
	}
	return &pub
}

// Address is a street address
type Address struct {
	// City
	City string `json:"city" xml:"city" form:"city"`
	// Street number
	Number int `json:"number" xml:"number" form:"number"`
	// US State Code
	State string `json:"state" xml:"state" form:"state"`
	// Street name
	Street string `json:"street" xml:"street" form:"street"`
	// US Zip code
	Zip int `json:"zip" xml:"zip" form:"zip"`
}

// Validate validates the Address type instance.
func (ut *Address) Validate() (err error) {
	if ut.Street == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "street"))
	}
	if ut.City == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "city"))
	}
	if ut.State == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "state"))
	}

	if len(ut.City) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.city`, ut.City, len(ut.City), 1, true))
	}
	if ut.Number < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.number`, ut.Number, 1, true))
	}
	if len(ut.State) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.state`, ut.State, len(ut.State), 2, true))
	}
	if len(ut.State) > 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.state`, ut.State, len(ut.State), 2, false))
	}
	if len(ut.Street) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.street`, ut.Street, len(ut.Street), 1, true))
	}
	if ut.Zip < 10000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.zip`, ut.Zip, 10000, true))
	}
	if ut.Zip > 99999 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.zip`, ut.Zip, 99999, false))
	}
	return
}

// CoursePatchPayload is the type used to patch courses
type coursePatchPayload struct {
	// Course description
	Description *string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	// Course end date/time
	EndTime *time.Time `json:"end_time,omitempty" xml:"end_time,omitempty" form:"end_time,omitempty"`
	// Course location
	Location *string `json:"location,omitempty" xml:"location,omitempty" form:"location,omitempty"`
	// Course start date/time
	StartTime *time.Time `json:"start_time,omitempty" xml:"start_time,omitempty" form:"start_time,omitempty"`
}

// Validate validates the coursePatchPayload type instance.
func (ut *coursePatchPayload) Validate() (err error) {
	if ut.Location != nil {
		if len(*ut.Location) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.location`, *ut.Location, len(*ut.Location), 2, true))
		}
	}
	return
}

// Publicize creates CoursePatchPayload from coursePatchPayload
func (ut *coursePatchPayload) Publicize() *CoursePatchPayload {
	var pub CoursePatchPayload
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.EndTime != nil {
		pub.EndTime = ut.EndTime
	}
	if ut.Location != nil {
		pub.Location = ut.Location
	}
	if ut.StartTime != nil {
		pub.StartTime = ut.StartTime
	}
	return &pub
}

// CoursePatchPayload is the type used to patch courses
type CoursePatchPayload struct {
	// Course description
	Description *string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	// Course end date/time
	EndTime *time.Time `json:"end_time,omitempty" xml:"end_time,omitempty" form:"end_time,omitempty"`
	// Course location
	Location *string `json:"location,omitempty" xml:"location,omitempty" form:"location,omitempty"`
	// Course start date/time
	StartTime *time.Time `json:"start_time,omitempty" xml:"start_time,omitempty" form:"start_time,omitempty"`
}

// Validate validates the CoursePatchPayload type instance.
func (ut *CoursePatchPayload) Validate() (err error) {
	if ut.Location != nil {
		if len(*ut.Location) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.location`, *ut.Location, len(*ut.Location), 2, true))
		}
	}
	return
}

// CoursePayload is the type used to create courses
type coursePayload struct {
	// Course description
	Description *string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	// Course end date/time
	EndTime *time.Time `json:"end_time,omitempty" xml:"end_time,omitempty" form:"end_time,omitempty"`
	// Course location
	Location *string `json:"location,omitempty" xml:"location,omitempty" form:"location,omitempty"`
	// Course name
	Name *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	// Course start date/time
	StartTime *time.Time `json:"start_time,omitempty" xml:"start_time,omitempty" form:"start_time,omitempty"`
}

// Validate validates the coursePayload type instance.
func (ut *coursePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.StartTime == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "start_time"))
	}
	if ut.EndTime == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "end_time"))
	}
	if ut.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if ut.Location != nil {
		if len(*ut.Location) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.location`, *ut.Location, len(*ut.Location), 2, true))
		}
	}
	if ut.Name != nil {
		if len(*ut.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 3, true))
		}
	}
	return
}

// Publicize creates CoursePayload from coursePayload
func (ut *coursePayload) Publicize() *CoursePayload {
	var pub CoursePayload
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.EndTime != nil {
		pub.EndTime = *ut.EndTime
	}
	if ut.Location != nil {
		pub.Location = *ut.Location
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.StartTime != nil {
		pub.StartTime = *ut.StartTime
	}
	return &pub
}

// CoursePayload is the type used to create courses
type CoursePayload struct {
	// Course description
	Description *string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	// Course end date/time
	EndTime time.Time `json:"end_time" xml:"end_time" form:"end_time"`
	// Course location
	Location string `json:"location" xml:"location" form:"location"`
	// Course name
	Name string `json:"name" xml:"name" form:"name"`
	// Course start date/time
	StartTime time.Time `json:"start_time" xml:"start_time" form:"start_time"`
}

// Validate validates the CoursePayload type instance.
func (ut *CoursePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if len(ut.Location) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.location`, ut.Location, len(ut.Location), 2, true))
	}
	if len(ut.Name) < 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, len(ut.Name), 3, true))
	}
	return
}

// RegistrationPatchPayload is the type used to patch registrations
type registrationPatchPayload struct {
	// Attendee address
	Address *address `json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	// Attendee first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty" form:"first_name,omitempty"`
	// Attendee last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty" form:"last_name,omitempty"`
}

// Validate validates the registrationPatchPayload type instance.
func (ut *registrationPatchPayload) Validate() (err error) {
	if ut.Address != nil {
		if err2 := ut.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	return
}

// Publicize creates RegistrationPatchPayload from registrationPatchPayload
func (ut *registrationPatchPayload) Publicize() *RegistrationPatchPayload {
	var pub RegistrationPatchPayload
	if ut.Address != nil {
		pub.Address = ut.Address.Publicize()
	}
	if ut.FirstName != nil {
		pub.FirstName = ut.FirstName
	}
	if ut.LastName != nil {
		pub.LastName = ut.LastName
	}
	return &pub
}

// RegistrationPatchPayload is the type used to patch registrations
type RegistrationPatchPayload struct {
	// Attendee address
	Address *Address `json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	// Attendee first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty" form:"first_name,omitempty"`
	// Attendee last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty" form:"last_name,omitempty"`
}

// Validate validates the RegistrationPatchPayload type instance.
func (ut *RegistrationPatchPayload) Validate() (err error) {
	if ut.Address != nil {
		if err2 := ut.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	return
}

// RegistrationPayload is the type used to create registrations
type registrationPayload struct {
	// Attendee address
	Address *address `json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	// The href to the course resource that describes the course being taught
	CourseHref *string `json:"course_href,omitempty" xml:"course_href,omitempty" form:"course_href,omitempty"`
	// Attendee first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty" form:"first_name,omitempty"`
	// Attendee last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty" form:"last_name,omitempty"`
}

// Validate validates the registrationPayload type instance.
func (ut *registrationPayload) Validate() (err error) {
	if ut.CourseHref == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "course_href"))
	}
	if ut.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if ut.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if ut.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	if ut.Address != nil {
		if err2 := ut.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.CourseHref != nil {
		if ok := goa.ValidatePattern(`/courses/[0-9]+`, *ut.CourseHref); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.course_href`, *ut.CourseHref, `/courses/[0-9]+`))
		}
	}
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	return
}

// Publicize creates RegistrationPayload from registrationPayload
func (ut *registrationPayload) Publicize() *RegistrationPayload {
	var pub RegistrationPayload
	if ut.Address != nil {
		pub.Address = ut.Address.Publicize()
	}
	if ut.CourseHref != nil {
		pub.CourseHref = *ut.CourseHref
	}
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
	}
	if ut.LastName != nil {
		pub.LastName = *ut.LastName
	}
	return &pub
}

// RegistrationPayload is the type used to create registrations
type RegistrationPayload struct {
	// Attendee address
	Address *Address `json:"address" xml:"address" form:"address"`
	// The href to the course resource that describes the course being taught
	CourseHref string `json:"course_href" xml:"course_href" form:"course_href"`
	// Attendee first name
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	// Attendee last name
	LastName string `json:"last_name" xml:"last_name" form:"last_name"`
}

// Validate validates the RegistrationPayload type instance.
func (ut *RegistrationPayload) Validate() (err error) {
	if ut.CourseHref == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "course_href"))
	}
	if ut.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if ut.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if ut.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	if ut.Address != nil {
		if err2 := ut.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`/courses/[0-9]+`, ut.CourseHref); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.course_href`, ut.CourseHref, `/courses/[0-9]+`))
	}
	if len(ut.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, ut.FirstName, len(ut.FirstName), 2, true))
	}
	if len(ut.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, ut.LastName, len(ut.LastName), 2, true))
	}
	return
}
