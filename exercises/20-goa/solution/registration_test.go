package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/bketelsen/buildingapis/exercises/20-goa/solution/app"
	"github.com/bketelsen/buildingapis/exercises/20-goa/solution/app/test"
	"github.com/bketelsen/buildingapis/exercises/library"
)

const (
	firstName  = "first"
	lastName   = "last"
	courseName = "building"
)

var (
	address = &app.Address{
		Number: 1,
		Street: "Main Street",
		City:   "Denver",
		State:  "CA",
		Zip:    80204,
	}

	descVar = "Building APIs in Go"
	start   = time.Now()
	end     = time.Now().Add(time.Duration(1) * time.Hour)

	coursePayload = &app.CreateCoursePayload{
		Name:        courseName,
		Description: &descVar,
		Location:    loc,
		StartTime:   start,
		EndTime:     end,
	}
)

func TestCreateRegistration(t *testing.T) {
	var (
		db         = library.NewEmptyDB()
		service    = NewService(db)
		ctrl       = NewRegistrationController(service, db)
		courseCtrl = NewCourseController(service, db)
	)

	_, cmt := test.CreateCourseCreated(t, context.Background(), service, courseCtrl, coursePayload)
	if cmt == nil {
		t.Errorf("create course: reponse media type is nil")
	}

	payload := &app.CreateRegistrationPayload{
		FirstName:  firstName,
		LastName:   lastName,
		CourseHref: cmt.Href,
		Address:    address,
	}
	rw, mt := test.CreateRegistrationCreated(t, context.Background(), service, ctrl, payload)
	if mt == nil {
		t.Errorf("create registration: reponse media type is nil")
	} else {
		if err := mt.Validate(); err != nil {
			t.Errorf("create registration: fail to validate response: %s", err)
		}
		if mt.FirstName != payload.FirstName {
			t.Errorf("create registration: invalid first name, expected %s, got %s", payload.FirstName, mt.FirstName)
		}
		if mt.LastName != payload.LastName {
			t.Errorf("create registration: invalid last name, expected %s, got %s", payload.LastName, mt.LastName)
		}
		if !reflect.DeepEqual(mt.Address, payload.Address) {
			t.Errorf("create registration: invalid address, expected %#+v, got %#+v", payload.Address, mt.Address)
		}
		if mt.Links == nil {
			t.Errorf("create registration: links is nil")
		} else {
			course := mt.Links.Course
			if app.CourseHref(course.ID) != cmt.Href {
				t.Errorf("create registration: invalid course href, expected %#+v, got %#+v", cmt.Href, app.CourseHref(course.ID))
			}
		}
		if mt.ID == 0 {
			t.Errorf("create registration: invalid ID 0")
		}
		if mt.Href != app.RegistrationHref(mt.ID) {
			t.Errorf("create registration: invalid href, expected %s, got %s", app.RegistrationHref(mt.ID), mt.Href)
		}
	}
	locH := rw.Header().Get("Location")
	if locH == "" {
		t.Errorf("create registration: missing Location header")
	} else {
		if locH != mt.Href {
			t.Errorf("create registration: invalid Location header, expected %s, got %s", mt.Href, locH)
		}
	}
}

func TestCreateRegistrationBadRequest(t *testing.T) {
	var (
		db         = library.NewEmptyDB()
		service    = NewService(db)
		ctrl       = NewRegistrationController(service, db)
		courseCtrl = NewCourseController(service, db)
	)

	_, cmt := test.CreateCourseCreated(t, context.Background(), service, courseCtrl, coursePayload)
	if cmt == nil {
		t.Fatalf("create course: reponse media type is nil")
	}
	courseHref := cmt.Href

	cases := map[string]struct {
		FirstName, LastName, CourseHref string
		Address                         *app.Address
		ErrorPattern                    string
	}{
		"no-firstname":    {"", lastName, courseHref, address, `attribute "first_name"`},
		"no-lastname":     {firstName, "", courseHref, address, `attribute "last_name"`},
		"no-course":       {firstName, lastName, "", address, `attribute "course_href"`},
		"no-address":      {firstName, lastName, courseHref, nil, `attribute "address"`},
		"short-firstname": {"a", lastName, courseHref, address, `must be greater or equal than 2`},
		"short-lastname":  {firstName, "a", courseHref, address, `must be greater or equal than 2`},
	}

	for k, tc := range cases {
		payload := &app.CreateRegistrationPayload{
			FirstName:  tc.FirstName,
			LastName:   tc.LastName,
			CourseHref: tc.CourseHref,
			Address:    tc.Address,
		}

		_, mt := test.CreateRegistrationBadRequest(t, context.Background(), service, ctrl, payload)

		if mt == nil {
			t.Errorf("%s: reponse media type is nil", k)
		} else {
			if !strings.Contains(mt.Detail, tc.ErrorPattern) {
				t.Errorf("%s: invalid error pattern, expected %s, got %s", k, tc.ErrorPattern, mt.Detail)
			}

		}
	}
}

func TestShowRegistrationOK(t *testing.T) {
	var (
		db         = library.NewEmptyDB()
		service    = NewService(db)
		ctrl       = NewRegistrationController(service, db)
		courseCtrl = NewCourseController(service, db)
	)

	_, comt := test.CreateCourseCreated(t, context.Background(), service, courseCtrl, coursePayload)
	if comt == nil {
		t.Fatalf("create course: reponse media type is nil")
	}
	courseHref := comt.Href
	payload := &app.CreateRegistrationPayload{
		FirstName:  firstName,
		LastName:   lastName,
		CourseHref: courseHref,
		Address:    address,
	}
	_, cmt := test.CreateRegistrationCreated(t, context.Background(), service, ctrl, payload)
	if cmt == nil {
		t.Fatalf("create registration: create reponse media type is nil")
	}

	_, mt := test.ShowRegistrationOK(t, context.Background(), service, ctrl, cmt.ID, nil)

	if diff := regDiff(mt, cmt); diff != "" {
		t.Errorf("show registration: response media type invalid: %s", diff)
	}
}

func TestShowRegistrationOKExtended(t *testing.T) {
	var (
		db         = library.NewEmptyDB()
		service    = NewService(db)
		ctrl       = NewRegistrationController(service, db)
		courseCtrl = NewCourseController(service, db)
	)

	_, comt := test.CreateCourseCreated(t, context.Background(), service, courseCtrl, coursePayload)
	if comt == nil {
		t.Fatalf("create course: reponse media type is nil")
	}
	courseHref := comt.Href
	payload := &app.CreateRegistrationPayload{
		FirstName:  firstName,
		LastName:   lastName,
		CourseHref: courseHref,
		Address:    address,
	}
	_, cmt := test.CreateRegistrationCreated(t, context.Background(), service, ctrl, payload)
	if cmt == nil {
		t.Fatalf("create registration: create reponse media type is nil")
	}
	expected := &app.RegistrationMediaExtended{
		Address:   address,
		Course:    comt,
		FirstName: firstName,
		Href:      cmt.Href,
		ID:        cmt.ID,
		LastName:  lastName,
		Links:     &app.RegistrationMediaLinks{Course: &app.CourseMediaLink{Href: comt.Href, ID: comt.ID}},
	}

	extended := "extended"
	_, mt := test.ShowRegistrationOKExtended(t, context.Background(), service, ctrl, cmt.ID, &extended)

	if diff := regDiffExt(mt, expected); diff != "" {
		t.Errorf("show registration extended: response media type invalid: %s", diff)
	}
}

func TestShowRegistrationNotFound(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewRegistrationController(service, db)
	)

	test.ShowRegistrationNotFound(t, context.Background(), service, ctrl, 100, nil)
}

func TestListRegistrationOK(t *testing.T) {
	var (
		db                      = library.NewEmptyDB()
		service                 = NewService(db)
		ctrl                    = NewRegistrationController(service, db)
		courseCtrl              = NewCourseController(service, db)
		courseHref, courseHref2 string
		courseID                int
	)
	{
		_, comt := test.CreateCourseCreated(t, context.Background(), service, courseCtrl, coursePayload)
		_, comt2 := test.CreateCourseCreated(t, context.Background(), service, courseCtrl, coursePayload)
		courseHref = comt.Href
		courseHref2 = comt2.Href
		elems := strings.Split(courseHref, "/")
		var err error
		courseID, err = strconv.Atoi(elems[len(elems)-1])
		if err != nil {
			t.Fatalf("list registration: invalid course ID %#v", courseID)
		}
	}
	payload := &app.CreateRegistrationPayload{
		FirstName:  firstName,
		LastName:   lastName,
		CourseHref: courseHref,
		Address:    address,
	}
	_, cmt := test.CreateRegistrationCreated(t, context.Background(), service, ctrl, payload)
	payload.CourseHref = courseHref2
	_, cmt2 := test.CreateRegistrationCreated(t, context.Background(), service, ctrl, payload)

	cases := map[string]struct {
		CourseID *int
		Expected []*app.RegistrationMedia
	}{
		"no-filter": {nil, []*app.RegistrationMedia{cmt, cmt2}},
		"filter":    {&courseID, []*app.RegistrationMedia{cmt}},
	}

	for k, tc := range cases {
		_, mt := test.ListRegistrationOK(t, context.Background(), service, ctrl, tc.CourseID)

		if mt == nil {
			t.Errorf("list registration %s: returned nil", k)
		} else {
			if len(mt) != len(tc.Expected) {
				t.Errorf("list registration %s: expected %d entry(ies), got %d", k, len(tc.Expected), len(mt))
			} else {
				for i, m := range tc.Expected {
					if diff := regDiff(mt[i], m); diff != "" {
						t.Errorf("list registration %s: response media type invalid: %s", k, diff)
					}
				}
			}
		}
	}
}

func regDiff(reg, expected *app.RegistrationMedia) (diff string) {
	if reg == nil && expected == nil {
		return ""
	}
	if reg == nil {
		return "reg is nil, expected isn't"
	}
	if expected == nil {
		return "reg is not nil, expected is"
	}
	if reg.Address == nil && expected.Address != nil {
		return "reg address is nil, expected isn't"
	}
	if reg.Address != nil && expected.Address == nil {
		return "reg address is not nil, expected is"
	}
	if reg.Address != nil {
		ra, oa := reg.Address, expected.Address
		if reg.Address.Number != expected.Address.Number {
			return fmt.Sprintf("reg address number is %d, expected is %d", ra.Number, oa.Number)
		}
		if reg.Address.Street != expected.Address.Street {
			return fmt.Sprintf("reg address street is %s, expected is %s", ra.Street, oa.Street)
		}
		if reg.Address.City != expected.Address.City {
			return fmt.Sprintf("reg address city is %s, expected is %s", ra.City, oa.City)
		}
		if reg.Address.State != expected.Address.State {
			return fmt.Sprintf("reg address state is %s, expected is %s", ra.State, oa.State)
		}
		if reg.Address.Zip != expected.Address.Zip {
			return fmt.Sprintf("reg address zip is %d, expected is %d", ra.Zip, oa.Zip)
		}
	}
	if reg.FirstName != expected.FirstName {
		return fmt.Sprintf("reg first name is %s, expected first name is %s", reg.FirstName, expected.FirstName)
	}
	if reg.LastName != expected.LastName {
		return fmt.Sprintf("reg last name is %s, expected last name is %s", reg.LastName, expected.LastName)
	}
	if reg.ID != expected.ID {
		return fmt.Sprintf("reg id is %d, expected id is %d", reg.ID, expected.ID)
	}
	if reg.Href != expected.Href {
		return fmt.Sprintf("reg href is %s, expected href is %s", reg.Href, expected.Href)
	}
	if reg.Links == nil && expected.Links != nil {
		return "reg links is nil, expected isn't"
	}
	if reg.Links != nil && expected.Links == nil {
		return "reg links is not nil, expected is"
	}
	if reg.Links != nil {
		if reg.Links.Course == nil && expected.Links != nil {
			return "reg links course is nil, expected isn't"
		}
		if reg.Links.Course != nil && expected.Links == nil {
			return "reg links course is not nil, expected is"
		}
		if reg.Links.Course != nil {
			if reg.Links.Course.ID != expected.Links.Course.ID {
				return fmt.Sprintf("reg links course id is %d, expected links course id is %d", reg.Links.Course.ID, expected.Links.Course.ID)
			}
			if reg.Links.Course.Href != expected.Links.Course.Href {
				return fmt.Sprintf("reg links course href is %s, expected links course href is %s", reg.Links.Course.Href, expected.Links.Course.Href)
			}
		}
	}
	return ""
}

func regDiffExt(reg, expected *app.RegistrationMediaExtended) (diff string) {
	r := &app.RegistrationMedia{
		Address:   reg.Address,
		FirstName: reg.FirstName,
		LastName:  reg.LastName,
		ID:        reg.ID,
		Href:      reg.Href,
		Links:     reg.Links,
	}
	e := &app.RegistrationMedia{
		Address:   expected.Address,
		FirstName: expected.FirstName,
		LastName:  expected.LastName,
		ID:        expected.ID,
		Href:      expected.Href,
		Links:     expected.Links,
	}
	if diff = regDiff(r, e); diff != "" {
		return diff
	}
	if reg.Course == nil && expected.Course != nil {
		return "reg links is nil, expected isn't"
	}
	if reg.Course != nil && expected.Course == nil {
		return "reg links is not nil, expected is"
	}
	if reg.Course != nil {
		co, ec := reg.Course, expected.Course
		if co.Description == nil && ec.Description != nil {
			return "reg course description is nil, expected isn't"
		}
		if co.Description != nil && ec.Description == nil {
			return "reg course description is not nil, ec is"
		}
		if co.Description != nil {
			if *co.Description != *ec.Description {
				return fmt.Sprintf("reg course description is %s, ec course description %s", *co.Description, *ec.Description)
			}
		}
		if co.EndTime != ec.EndTime {
			return fmt.Sprintf("reg course EndTime is %v, expected course EndTime %v", co.EndTime, ec.EndTime)
		}
		if co.Href != ec.Href {
			return fmt.Sprintf("reg course Href is %s, expected course Href %s", co.Href, ec.Href)
		}
		if co.ID != ec.ID {
			return fmt.Sprintf("reg course ID is %d, expected course ID %d", co.ID, ec.ID)
		}
		if co.Location != ec.Location {
			return fmt.Sprintf("reg course Location is %s, expected course Location %s", co.Location, ec.Location)
		}
		if co.Name != ec.Name {
			return fmt.Sprintf("reg course Name is %s, expected course Name %s", co.Name, ec.Name)
		}
		if co.StartTime != ec.StartTime {
			return fmt.Sprintf("reg course StartTime is %v, expected course StartTime %v", co.StartTime, ec.StartTime)
		}
	}
	return ""
}
