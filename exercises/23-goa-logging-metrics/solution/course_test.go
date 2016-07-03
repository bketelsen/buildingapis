package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/bketelsen/buildingapis/exercises/20-goa/solution/app"
	"github.com/bketelsen/buildingapis/exercises/20-goa/solution/app/test"
	"github.com/bketelsen/buildingapis/exercises/library"
)

const (
	loc  = "Denver"
	desc = "Building APIs"
)

// Variable so we can take the address.
var name = "building"

func TestCreateCourse(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewCourseController(service, db)
		start   = time.Now()
		end     = time.Now().Add(time.Duration(1) * time.Hour)
	)

	cases := map[string]struct {
		Name, Location, Description string
		StartTime, EndTime          time.Time
	}{
		"complete": {name, loc, desc, start, end},
		"no-desc":  {name, loc, "", start, end},
	}

	for k, tc := range cases {
		payload := &app.CreateCoursePayload{
			EndTime:   tc.EndTime,
			Location:  tc.Location,
			Name:      tc.Name,
			StartTime: tc.StartTime,
		}
		if tc.Description != "" {
			payload.Description = &tc.Description
		}

		rw, mt := test.CreateCourseCreated(t, context.Background(), service, ctrl, payload)

		if mt == nil {
			t.Errorf("%s: reponse media type is nil", k)
		} else {
			if err := mt.Validate(); err != nil {
				t.Errorf("%s: fail to validate response: %s", k, err)
			}
			if mt.Name != payload.Name {
				t.Errorf("%s: invalid name, expected %s, got %s", k, payload.Name, mt.Name)
			}
			if mt.Description == nil && payload.Description != nil ||
				payload.Description == nil && mt.Description != nil ||
				mt.Description != nil && (*mt.Description != *payload.Description) {

				t.Errorf("%s: invalid description, expected %s, got %s", k, payload.Description, mt.Description)
			}
			if mt.Location != payload.Location {
				t.Errorf("%s: invalid location, expected %s, got %s", k, payload.Location, mt.Location)
			}
			if mt.StartTime != payload.StartTime {
				t.Errorf("%s: invalid start time, expected %s, got %s", k, payload.StartTime, mt.StartTime)
			}
			if mt.EndTime != payload.EndTime {
				t.Errorf("%s: invalid end time, expected %s, got %s", k, payload.EndTime, mt.EndTime)
			}
			if mt.ID == 0 {
				t.Errorf("%s: invalid ID 0", k)
			}
			if mt.Href != app.CourseHref(mt.ID) {
				t.Errorf("%s: invalid href, expected %s, got %s", k, app.CourseHref(mt.ID), mt.Href)
			}
		}
		locH := rw.Header().Get("Location")
		if locH == "" {
			t.Errorf("%s: missing Location header", k)
		} else {
			if locH != mt.Href {
				t.Errorf("%s: invalid Location header, expected %s, got %s", k, mt.Href, locH)
			}
		}
	}
}

func TestCreateCourseBadRequest(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewCourseController(service, db)
		start   = time.Now()
		end     = time.Now().Add(time.Duration(1) * time.Hour)
	)

	cases := map[string]struct {
		Name, Location, Description string
		StartTime, EndTime          time.Time
		ErrorPattern                string
	}{
		"no-name":    {"", loc, desc, start, end, `attribute "name"`},
		"no-loc":     {name, "", desc, start, end, `attribute "location"`},
		"short-name": {"a", loc, desc, start, end, `must be greater or equal than 3`},
	}

	for k, tc := range cases {
		payload := &app.CreateCoursePayload{
			Name:      tc.Name,
			EndTime:   tc.EndTime,
			Location:  tc.Location,
			StartTime: tc.StartTime,
		}
		if tc.Description != "" {
			payload.Description = &tc.Description
		}

		_, mt := test.CreateCourseBadRequest(t, context.Background(), service, ctrl, payload)

		if mt == nil {
			t.Errorf("%s: reponse media type is nil", k)
		} else {
			if mt.Status != 400 {
				t.Errorf("%s: invalid status, expected 400, got %s", k, mt.Status)
			}
			if !strings.Contains(mt.Detail, tc.ErrorPattern) {
				t.Errorf("%s: invalid error pattern, expected %s, got %s", k, tc.ErrorPattern, mt.Detail)
			}

		}
	}
}

func TestShowCourseOK(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewCourseController(service, db)
		start   = time.Now()
		end     = time.Now().Add(time.Duration(1) * time.Hour)
	)

	cases := map[string]struct {
		Name, Location, Description string
		StartTime, EndTime          time.Time
	}{
		"complete": {name, loc, desc, start, end},
		"no-desc":  {name, loc, "", start, end},
	}

	for k, tc := range cases {
		payload := &app.CreateCoursePayload{
			Name:      tc.Name,
			Location:  tc.Location,
			StartTime: tc.StartTime,
			EndTime:   tc.EndTime,
		}
		if tc.Description != "" {
			payload.Description = &tc.Description
		}

		_, cmt := test.CreateCourseCreated(t, context.Background(), service, ctrl, payload)

		if cmt == nil {
			t.Fatalf("%s: create reponse media type is nil", k)
		}

		_, mt := test.ShowCourseOK(t, context.Background(), service, ctrl, cmt.ID)

		if !reflect.DeepEqual(mt, cmt) {
			t.Errorf("%s: show response media type invalid, expected %#+v, got %#+v", k, cmt, mt)
		}
	}
}

func TestShowCourseNotFound(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewCourseController(service, db)
	)

	test.ShowCourseNotFound(t, context.Background(), service, ctrl, 100)
}

func TestShowCourseDeleteNoContent(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewCourseController(service, db)
		start   = time.Now()
		end     = time.Now().Add(time.Duration(1) * time.Hour)
	)

	cases := map[string]struct {
		Name, Location, Description string
		StartTime, EndTime          time.Time
	}{
		"complete": {name, loc, desc, start, end},
		"no-desc":  {name, loc, "", start, end},
	}

	for k, tc := range cases {
		payload := &app.CreateCoursePayload{
			Name:      tc.Name,
			Location:  tc.Location,
			StartTime: tc.StartTime,
			EndTime:   tc.EndTime,
		}
		if tc.Description != "" {
			payload.Description = &tc.Description
		}

		_, cmt := test.CreateCourseCreated(t, context.Background(), service, ctrl, payload)

		if cmt == nil {
			t.Fatalf("%s: create reponse media type is nil", k)
		}

		test.ShowCourseOK(t, context.Background(), service, ctrl, cmt.ID)
		test.DeleteCourseNoContent(t, context.Background(), service, ctrl, cmt.ID)
		test.ShowCourseNotFound(t, context.Background(), service, ctrl, cmt.ID)
	}
}

func TestDeleteCourseNotFound(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewCourseController(service, db)
	)

	test.DeleteCourseNotFound(t, context.Background(), service, ctrl, 100)
}

func TestListCourseOK(t *testing.T) {
	var (
		db      = library.NewEmptyDB()
		service = NewService(db)
		ctrl    = NewCourseController(service, db)
		start   = time.Now()
		end     = time.Now().Add(time.Duration(1) * time.Hour)
	)
	payload := &app.CreateCoursePayload{
		Name:      name,
		Location:  loc,
		StartTime: start,
		EndTime:   end,
	}
	_, cmt := test.CreateCourseCreated(t, context.Background(), service, ctrl, payload)
	payload.Name = "name2"
	_, cmt2 := test.CreateCourseCreated(t, context.Background(), service, ctrl, payload)

	cases := map[string]struct {
		CourseName *string
		Expected   []*app.CourseMedia
	}{
		"no-filter": {nil, []*app.CourseMedia{cmt, cmt2}},
		"filter":    {&name, []*app.CourseMedia{cmt}},
	}

	for k, tc := range cases {
		_, mt := test.ListCourseOK(t, context.Background(), service, ctrl, tc.CourseName)

		if mt == nil {
			t.Errorf("list course %s: returned nil", k)
		} else {
			if len(mt) != len(tc.Expected) {
				t.Errorf("list course %s: expected %d entry(ies), got %d: %#v", k, len(tc.Expected), len(mt), toString(mt))
			} else {
				for i, e := range tc.Expected {
					if !reflect.DeepEqual(mt[i], e) {
						t.Errorf("list course %s: response media type invalid, expected %#+v, got %#+v", k, i, mt[i])
					}
				}
			}
		}
	}
}

func toString(mt app.CourseMediaCollection) string {
	elems := make([]string, len(mt))
	for i, m := range mt {
		if m != nil {
			elems[i] = fmt.Sprintf("%#+v", *m)
		} else {
			elems[i] = "<nil>"
		}
	}
	return "{ " + strings.Join(elems, ", ") + " }"
}
