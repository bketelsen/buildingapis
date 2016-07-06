package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bketelsen/buildingapis/exercises/library"
)

const (
	courseBase       = "/api/courses/"
	registrationBase = "/api/registrations/"
)

// ErrBadPath is the error returned when the request
// path isn't valid.
var ErrBadPath = errors.New("Bad Request Path")

// Validation Errors
var (
	// ErrCourseNameRequired is returned when there is no course name in the post
	ErrCourseNameRequired = errors.New("Course Name Required")
	// ErrCourseStartTimeRequired is returned when there is no start time in the post
	ErrCourseStartTimeRequired = errors.New("Course Start Time Required")
	// ErrCourseEndTimeRequired is returned when there is no end time in the post
	ErrCourseEndTimeRequired = errors.New("Course End Time Required")
	// ErrCourseLocationRequired is returned when there is no location in the post
	ErrCourseLocationRequired = errors.New("Course Location Required")
	// ErrcourseNameLength is returned when the course name is less than 3 characters long
	ErrCourseNameLength = errors.New("Course Name Too Short")
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

// Course is a class that can be taken
type Course struct {
	ID int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
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

// Registration is the record of someone signing up to take a course
type Registration struct {
	// ID of registration
	ID int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	// CourseID of registered course
	CourseID int `json:"course_id,omitempty" xml:"course_id,omitempty" form:"course_id,omitempty"`
	// Attendee address
	Address *Address `json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	// The href to the course resource that describes the course being taught
	// Attendee first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty" form:"first_name,omitempty"`
	// Attendee last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty" form:"last_name,omitempty"`
}

// CourseServer services requests for Courses.
// It is a struct so we can do dependency injection
// in tests and use a different database
type CourseServer struct {
	DB *library.MemDB
}

func main() {
	db := library.NewDB()
	cs := &CourseServer{
		DB: db,
	}
	http.Handle(courseBase, logMiddleware(cs))
	http.HandleFunc(registrationBase, registrations)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func (cs *CourseServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Allow", "OPTIONS,GET,POST")
	if r.Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "GET":
		//	get course or courses
		id, err := idOrList(courseBase, r.URL.Path)
		if err != nil {
			jsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		if id == "" {
			cc, err := cs.DB.List("courses", "id", nil)
			if err != nil {
				if err == library.ErrNotFound {
					jsonError(w, "Not Found", http.StatusNotFound)
					return
				}
				jsonError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			courselist := courselistToCourseSlice(cc)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(courselist); err != nil {
				log.Println("Encode:", err)
				return
			}
		} else {
			c, err := cs.DB.Get("courses", "id", id)
			if err != nil {
				if err == library.ErrNotFound {
					jsonError(w, "Not Found", http.StatusNotFound)
					return
				}
				jsonError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			course := courseToCourse(c)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(course); err != nil {
				log.Println("Encode:", err)
				return
			}
		}

	case "POST":
		var course Course
		if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
			jsonError(w, "Invalid JSON", 400)
			return
		}
		// validate it
		err := validateCourse(&course)
		if err != nil {
			jsonError(w, err.Error(), 400)
			return
		}
		// convert it
		lc := convertCourse(course)
		// save it
		err = cs.DB.Insert("courses", lc)
		if err != nil {
			jsonError(w, err.Error(), 400)
		}
		w.WriteHeader(http.StatusCreated)
		return
	default:
		w.Header().Set("Allow", "GET,POST")
		jsonError(w, "Method Not Allowed", 405)
	}
}

func registrations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET,POST")
	jsonError(w, "Not Implemented", http.StatusNotImplemented)
}

func convertCourse(c Course) *library.CourseModel {
	id := strconv.Itoa(c.ID)
	lc := &library.CourseModel{
		ID:          id,
		Name:        c.Name,
		Location:    c.Location,
		Description: *c.Description,
		StartTime:   c.StartTime,
		EndTime:     c.EndTime,
	}
	return lc
}
func jsonError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(struct {
		Message string
	}{
		Message: msg,
	})
	if err != nil {
		log.Println("jsonError:", err)
		return
	}
}

// idOrList returns the ID in the request if given, or
// an empty string if none.  It returns an error if the
// Requsted URL doesn't match the expected pattern
func idOrList(base, path string) (string, error) {
	// go's mux will send /api/courses and /api/courses/
	// to the same handler if the trailing slash is registered as the
	// path.  Special case that here.
	if path+"/" == base {
		path = path + "/"
	}
	remains := strings.Replace(path, base, "", -1)
	if strings.HasPrefix(remains, "/") {
		remains = remains[1:]
	}
	if remains == "" {
		return "", nil
	}

	if strings.HasSuffix(remains, "/") {
		remains = remains[0 : len(remains)-1]
	}
	if a := strings.Split(remains, "/"); len(a) > 1 {
		return "", ErrBadPath
	}
	return remains, nil
}

func courseToCourse(i interface{}) *Course {
	m := i.(*library.CourseModel)

	id, err := strconv.Atoi(m.ID)
	if err != nil {
		panic("invalid course ID - must be an int") // bug
	}
	mt := &Course{
		ID:        id,
		Name:      m.Name,
		Location:  m.Location,
		StartTime: m.StartTime,
		EndTime:   m.EndTime,
	}
	if m.Description != "" {
		mt.Description = &m.Description
	}
	return mt
}

func courselistToCourseSlice(i []interface{}) []*Course {
	cc := make([]*Course, len(i))
	for x, course := range i {
		mt := courseToCourse(course)
		cc[x] = mt
	}
	return cc
}

func validateCourse(c *Course) error {
	// Required:
	// name, start_time, end_time, location
	if c.Name == "" {
		return ErrCourseNameRequired
	}
	// compare start and end time to a nil time
	// to see if they're set
	var t time.Time
	if c.StartTime == t {
		return ErrCourseStartTimeRequired
	}
	if c.EndTime == t {
		return ErrCourseEndTimeRequired
	}
	if c.Location == "" {
		return ErrCourseLocationRequired
	}
	// Minimum Length
	// name: 3

	if len(c.Name) < 3 {
		return ErrCourseNameLength
	}
	return nil
}
