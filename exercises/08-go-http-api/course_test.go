package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bketelsen/buildingapis/exercises/library"
)

func init() {
	db = library.NewDB()
}

func TestGetOneCourse(t *testing.T) {
	// register the handler function with the httptest Server
	ts := httptest.NewServer(http.HandlerFunc(courses))
	defer ts.Close()

	// make a request
	// this test is a brittle - what happens if
	// there isn't a course with ID 1?
	resp, err := http.Get(ts.URL + "/api/courses/1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
	// this is the hard way
	got := new(Course)
	if err = json.NewDecoder(resp.Body).Decode(got); err != nil {
		t.Error(err.Error())
	}

	// Get the golden record from the datastore to compare with
	want, err := db.Get("courses", "id", "1")
	if err != nil {
		t.Errorf("Couldn't read from memdb: %v", err)
	}
	wanted := want.(*library.CourseModel)

	// now compare as many of the fields as you want
	// remembering that there may be pointers or string <-> int
	// conversions required (like for the ID field)
	if got.Location != wanted.Location {
		t.Errorf("Got %s, wanted %s Location", got.Location, wanted.Location)
	}
}

func TestGetCourses(t *testing.T) {
	// register the handler function with the httptest Server
	ts := httptest.NewServer(http.HandlerFunc(courses))
	defer ts.Close()

	// make a request
	resp, err := http.Get(ts.URL + "/api/courses/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
	// this is the hard way
	var got []Course
	if err = json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Error(err.Error())
	}

	want, err := db.List("courses", "id", nil)
	if err != nil {
		t.Errorf("Couldn't read from memdb: %v", err)
	}

	if len(got) != len(want) {
		t.Errorf("Expected %d records, got %d.", len(want), len(got))
	}
}
