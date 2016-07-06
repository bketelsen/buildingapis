package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bketelsen/buildingapis/exercises/library"
)

func TestGetOneCourse(t *testing.T) {
	db := library.NewDB()
	cs := &CourseServer{
		DB: db,
	}
	// register the handler function with the httptest Server
	ts := httptest.NewServer(http.Handler(cs))
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

func TestGetOneCourseNotFound(t *testing.T) {
	db := library.NewDB()
	cs := &CourseServer{
		DB: db,
	}
	// register the handler function with the httptest Server
	ts := httptest.NewServer(http.Handler(cs))
	defer ts.Close()

	// make a request
	resp, err := http.Get(ts.URL + "/api/courses/30")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 404 {
		t.Fatalf("Received %d response, expected  %d\n", resp.StatusCode, http.StatusNotFound)
	}
	defer resp.Body.Close()
	// this is the hard way
	var got map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Error(err.Error())
	}
	var msg interface{}
	var ok bool
	if msg, ok = got["Message"]; !ok {
		t.Error("expected a message key in response json")
	}
	if msg.(string) != "Not Found" {
		t.Error("Expected not found message")
	}

}
func TestGetCourses(t *testing.T) {
	db := library.NewDB()
	cs := &CourseServer{
		DB: db,
	}
	// register the handler function with the httptest Server
	ts := httptest.NewServer(http.Handler(cs))
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

func TestPostCourseGood(t *testing.T) {
	db := library.NewEmptyDB()
	cs := &CourseServer{
		DB: db,
	}
	// register the handler function with the httptest Server
	ts := httptest.NewServer(http.Handler(cs))
	defer ts.Close()

	var desc string
	desc = "Course Description"
	course := Course{
		ID:          1,
		Name:        "Best Course Ever",
		Location:    "Denver",
		Description: &desc,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(10 * time.Minute),
	}
	b, err := json.Marshal(course)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("POST", ts.URL+"/api/courses/", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		t.Fatalf("Received non-201 response: %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
	c, err := cs.DB.Get("courses", "id", "1")
	if err != nil {
		t.Errorf("Expected course, got error: %v", err)
	}
	rc, ok := c.(*library.CourseModel)
	if !ok {
		t.Errorf("invalid response object")
	}
	if rc.Location != course.Location {
		t.Errorf("Expected %s, got %s", course.Location, rc.Location)
	}
	// Test more or all of these fields if you wish to be complete
}

func TestPostCourseBad(t *testing.T) {
	db := library.NewEmptyDB()
	cs := &CourseServer{
		DB: db,
	}
	// register the handler function with the httptest Server
	ts := httptest.NewServer(http.Handler(cs))
	defer ts.Close()

	var desc string
	desc = "Course Description"
	course := Course{
		ID:          1,
		Name:        "1",
		Location:    "Denver",
		Description: &desc,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(10 * time.Minute),
	}
	b, err := json.Marshal(course)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("POST", ts.URL+"/api/courses/", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected response wanted %d, got %d\n", http.StatusBadRequest, resp.StatusCode)
	}
	defer resp.Body.Close()

	var got map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Error(err.Error())
	}
	var msg interface{}
	var ok bool
	if msg, ok = got["Message"]; !ok {
		t.Error("expected a message key in response json")
	}
	if msg.(string) != "Course Name Too Short" {
		t.Error("Expected validation error")
	}
}
