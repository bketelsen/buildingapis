package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"goji.io"
	"goji.io/pat"

	"github.com/bketelsen/buildingapis/exercises/library"

	"golang.org/x/net/context"
)

type coursePayload struct {
	Name        string    `json:"name" xml:"name" form:"name"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	StartTime   time.Time `json:"start_time" xml:"start_time" form:"start_time"`
	EndTime     time.Time `json:"end_time" xml:"end_time" form:"end_time"`
	Location    string    `json:"location" xml:"location" form:"location"`
}

type courseMedia struct {
	ID          int       `json:"id" xml:"id" form:"id"`
	Href        string    `json:"href" xml:"href" form:"href"`
	Name        string    `json:"name" xml:"name" form:"name"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	StartTime   time.Time `json:"start_time" xml:"start_time" form:"start_time"`
	EndTime     time.Time `json:"end_time" xml:"end_time" form:"end_time"`
	Location    string    `json:"location" xml:"location" form:"location"`
}

func createCourse(db *library.MemDB) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		var payload coursePayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			respondBadRequest(w, "invalid request body: %s", err)
			return
		}
		var desc string
		if payload.Description != nil {
			desc = *payload.Description
		}
		model := &library.CourseModel{
			ID:          library.NewID(),
			Name:        payload.Name,
			Description: desc,
			StartTime:   payload.StartTime,
			EndTime:     payload.EndTime,
			Location:    payload.Location,
		}
		if err := db.Insert("courses", model); err != nil {
			respondInternal(w, "failed to insert course: %s", err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("/courses/%d", model.ID))
		w.WriteHeader(201)
		err = json.NewEncoder(w).Encode(courseToMedia(model))
		if err != nil {
			respondInternal(w, "failed to write response: %s", err)
			return
		}
	}
}

func showCourse(db *library.MemDB) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		id := pat.Param(ctx, "id")
		im, err := db.Get("courses", "id", id)
		if err != nil && err != library.ErrNotFound {
			respondInternal(w, "failed to query course: %s", err)
			return
		}
		if im == nil {
			respondNotFound(w)
			return
		}
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(courseToMedia(im.(*library.CourseModel)))
		if err != nil {
			respondInternal(w, "failed to write response: %s", err)
			return
		}
	}
}

func deleteCourse(db *library.MemDB) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		id := pat.Param(ctx, "id")
		if err := db.Delete("courses", "id", id); err != nil {
			if err == library.ErrNotFound {
				respondNotFound(w)
				return
			}
			respondInternal(w, "failed to delete course: %s", err)
			return
		}
		w.WriteHeader(204)
	}
}

func courseToMedia(m *library.CourseModel) *courseMedia {
	id, err := strconv.Atoi(m.ID)
	if err != nil {
		panic("invalid course ID - must be an int") // bug
	}
	mt := &courseMedia{
		ID:        id,
		Href:      fmt.Sprintf("/courses/%d", id),
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
