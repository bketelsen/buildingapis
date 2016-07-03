//************************************************************************//
// User Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa/design
// --out=$(GOPATH)/src/github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

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

// DecodeError decodes the Error instance encoded in resp body.
func (c *Client) DecodeError(resp *http.Response) (*goa.Error, error) {
	var decoded goa.Error
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CourseMedia is the media type used to render courses
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

// DecodeCourseMedia decodes the CourseMedia instance encoded in resp body.
func (c *Client) DecodeCourseMedia(resp *http.Response) (*CourseMedia, error) {
	var decoded CourseMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RegistrationMedia is the media type used to render registrations
type RegistrationMedia struct {
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
}

// DecodeRegistrationMedia decodes the RegistrationMedia instance encoded in resp body.
func (c *Client) DecodeRegistrationMedia(resp *http.Response) (*RegistrationMedia, error) {
	var decoded RegistrationMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RegistrationMediaCollection media type is a collection of RegistrationMedia.
type RegistrationMediaCollection []*RegistrationMedia

// DecodeRegistrationMediaCollection decodes the RegistrationMediaCollection instance encoded in resp body.
func (c *Client) DecodeRegistrationMediaCollection(resp *http.Response) (RegistrationMediaCollection, error) {
	var decoded RegistrationMediaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
