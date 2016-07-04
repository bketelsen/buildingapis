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

// in memory datastore
var db *library.MemDB

const (
	courseBase       = "/api/courses/"
	registrationBase = "/api/registrations/"
)

// ErrBadPath is the error returned when the request
// path isn't valid.
var ErrBadPath = errors.New("Bad Request Path")

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
	ID int `json:id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
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
	ID int `json:id,omitempty" xml:"id,omitempty" form:"id,omitempty"`

	CourseID int `json:course_id,omitempty" xml:"course_id,omitempty" form:"course_id,omitempty"`
	// Attendee address
	Address *Address `json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	// The href to the course resource that describes the course being taught
	// Attendee first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty" form:"first_name,omitempty"`
	// Attendee last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty" form:"last_name,omitempty"`
}

func main() {

	db = library.NewDB()

	http.HandleFunc(courseBase, courses)
	http.HandleFunc(registrationBase, registrations)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func courses(w http.ResponseWriter, r *http.Request) {
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
			cc, err := db.List("courses", "id", nil)
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
			c, err := db.Get("courses", "id", id)
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

		// save it
		//c := saveAndGetCourse()
		var c Course

		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			log.Println("serveEndpoint: POST: Encode:", err)
		}
	default:
		w.Header().Set("Allow", "GET,POST")
		jsonError(w, "Method Not Allowed", 405)
	}
}

func registrations(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Allow", "OPTIONS,GET,POST")
	if r.Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "GET":
		//	get registration or registrations
		id, err := idOrList(registrationBase, r.URL.Path)
		if err != nil {
			jsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		if id == "" {
			cc, err := db.List("registrations", "id", nil)
			if err != nil {
				if err == library.ErrNotFound {
					jsonError(w, "Not Found", http.StatusNotFound)
					return
				}
				jsonError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			reglist := reglistToRegistrationSlice(cc)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(reglist); err != nil {
				log.Println("Encode:", err)
				return
			}
		} else {
			c, err := db.Get("registrations", "id", id)
			if err != nil {
				if err == library.ErrNotFound {
					jsonError(w, "Not Found", http.StatusNotFound)
					return
				}
				jsonError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			reg := regToRegistration(c)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(reg); err != nil {
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

		// save it
		//c := saveAndGetCourse()
		var c Course

		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			log.Println("serveEndpoint: POST: Encode:", err)
		}
	default:
		w.Header().Set("Allow", "GET,POST")
		jsonError(w, "Method Not Allowed", 405)
	}
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

func regToRegistration(i interface{}) *Registration {
	m := i.(*library.RegistrationModel)
	id, err := strconv.Atoi(m.ID)
	if err != nil {
		panic("invalid Registration ID - must be an int") // bug
	}
	courseID, err := strconv.Atoi(m.CourseID)
	if err != nil {
		panic("invalid Course ID - must be an int") // bug
	}
	mt := &Registration{
		ID:        id,
		FirstName: &m.FirstName,
		LastName:  &m.LastName,
		CourseID:  courseID,
		Address: &Address{
			Number: m.Address.Number,
			Street: m.Address.Street,
			City:   m.Address.City,
			State:  m.Address.State,
			Zip:    m.Address.Zip,
		},
	}
	return mt
}

func courselistToCourseSlice(i []interface{}) []*Course {
	cc := make([]*Course, len(i))
	for x, course := range i {
		m := course.(*library.CourseModel)
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
		cc[x] = mt
	}
	return cc
}
func reglistToRegistrationSlice(i []interface{}) []*Registration {
	cc := make([]*Registration, len(i))
	for x, reg := range i {
		m := reg.(*library.RegistrationModel)
		id, err := strconv.Atoi(m.ID)
		if err != nil {
			panic("invalid Registration ID - must be an int") // bug
		}

		courseID, err := strconv.Atoi(m.CourseID)
		if err != nil {
			panic("invalid Course ID - must be an int") // bug
		}
		mt := &Registration{
			ID:        id,
			FirstName: &m.FirstName,
			LastName:  &m.LastName,
			CourseID:  courseID,
			Address: &Address{
				Number: m.Address.Number,
				Street: m.Address.Street,
				City:   m.Address.City,
				State:  m.Address.State,
				Zip:    m.Address.Zip,
			},
		}

		cc[x] = mt
	}
	return cc
}
