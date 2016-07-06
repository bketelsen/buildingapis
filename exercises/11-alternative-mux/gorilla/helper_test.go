package main

import (
	"testing"
	"time"
)

func TestCourseValidation(t *testing.T) {
	// table driven tests are awesome
	// use them every chance you get
	// to save typing!
	tests := []struct {
		Course Course
		Valid  bool
	}{
		{Course: Course{
			ID:        1,
			Name:      "My Course",
			Location:  "Denver",
			StartTime: time.Now(),
			EndTime:   time.Now(),
		}, Valid: true},
		{Course: Course{
			ID:        2,
			Name:      "My", // min length
			Location:  "Denver",
			StartTime: time.Now(),
			EndTime:   time.Now(),
		}, Valid: false},
		{Course: Course{ // missing name
			ID:        2,
			Location:  "Denver",
			StartTime: time.Now(),
			EndTime:   time.Now(),
		}, Valid: false},
		{Course: Course{ // missing location
			ID:        2,
			Name:      "My Course",
			StartTime: time.Now(),
			EndTime:   time.Now(),
		}, Valid: false},
		{Course: Course{ // missing start time
			ID:       2,
			Name:     "My Course",
			Location: "Denver",
			EndTime:  time.Now(),
		}, Valid: false},
		{Course: Course{ // missing end time
			ID:        1,
			Name:      "My Course",
			Location:  "Denver",
			StartTime: time.Now(),
		}, Valid: false},
	}
	for _, ct := range tests {
		err := validateCourse(&ct.Course)
		if !ct.Valid && err == nil {
			t.Errorf("Expected Validation Error")
		}
	}
}
