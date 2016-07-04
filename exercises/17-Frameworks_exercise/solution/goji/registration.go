package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"goji.io/pat"

	"goji.io"

	"github.com/bketelsen/buildingapis/exercises/library"

	"golang.org/x/net/context"
)

// Address is a street address
type address struct {
	Number int    `json:"number,omitempty" xml:"number,omitempty" form:"number,omitempty"`
	Street string `json:"street,omitempty" xml:"street,omitempty" form:"street,omitempty"`
	City   string `json:"city,omitempty" xml:"city,omitempty" form:"city,omitempty"`
	State  string `json:"state,omitempty" xml:"state,omitempty" form:"state,omitempty"`
	Zip    int    `json:"zip,omitempty" xml:"zip,omitempty" form:"zip,omitempty"`
}

// registrationPayload is the type used to create registrations
type registrationPayload struct {
	CourseHref string   `json:"course_href" xml:"course_href" form:"course_href"`
	FirstName  string   `json:"first_name" xml:"first_name" form:"first_name"`
	LastName   string   `json:"last_name" xml:"last_name" form:"last_name"`
	Address    *address `json:"address" xml:"address" form:"address"`
}

type registrationMedia struct {
	ID        int      `json:"id" xml:"id" form:"id"`
	Href      string   `json:"href" xml:"href" form:"href"`
	FirstName string   `json:"first_name" xml:"first_name" form:"first_name"`
	LastName  string   `json:"last_name" xml:"last_name" form:"last_name"`
	Address   *address `json:"address" xml:"address" form:"address"`
}

func createRegistration(db *library.MemDB) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		var payload registrationPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			respondBadRequest(w, "invalid request body: %s", err)
			return
		}
		model := &library.RegistrationModel{
			ID:        library.NewID(),
			CourseID:  courseIDFromHref(payload.CourseHref),
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Address:   addressFromPayload(payload.Address),
		}
		if err := db.Insert("registrations", model); err != nil {
			respondInternal(w, "failed to insert registration: %s", err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("/registrations/%d", model.ID))
		w.WriteHeader(201)
		err = json.NewEncoder(w).Encode(registrationToMedia(model))
		if err != nil {
			respondInternal(w, "failed to write response: %s", err)
			return
		}
	}
}

func showRegistration(db *library.MemDB) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		id := pat.Param(ctx, "id")
		im, err := db.Get("registrations", "id", id)
		if err != nil && err != library.ErrNotFound {
			respondInternal(w, "failed to query registration: %s", err)
			return
		}
		if im == nil {
			respondNotFound(w)
			return
		}
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(registrationToMedia(im.(*library.RegistrationModel)))
		if err != nil {
			respondInternal(w, "failed to write response: %s", err)
			return
		}
	}
}

func listRegistrations(db *library.MemDB) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		im, err := db.List("registrations", "id", nil)
		if err != nil && err != library.ErrNotFound {
			respondInternal(w, "failed to query registrations: %s", err)
			return
		}
		if im == nil {
			respondNotFound(w)
			return
		}
		med := make([]*registrationMedia, len(im))
		for i, m := range im {
			med[i] = registrationToMedia(m.(*library.RegistrationModel))
		}
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(med)
		if err != nil {
			respondInternal(w, "failed to write response: %s", err)
			return
		}
	}
}

// addressFromPayload creates an address model.
func addressFromPayload(payload *address) *library.Address {
	return &library.Address{
		Number: payload.Number,
		Street: payload.Street,
		City:   payload.City,
		State:  payload.State,
		Zip:    payload.Zip,
	}
}

// courseIDFromHref returns the ID of a course model given a resource href.
func courseIDFromHref(href string) string {
	elems := strings.Split(href, "/")
	return elems[len(elems)-1]
}

func registrationToMedia(m *library.RegistrationModel) *registrationMedia {
	id, err := strconv.Atoi(m.ID)
	if err != nil {
		panic("invalid registration ID - must be an int") // bug
	}
	mt := &registrationMedia{
		ID:        id,
		Href:      fmt.Sprintf("/registrations/%d", id),
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Address: &address{
			Number: m.Address.Number,
			Street: m.Address.Street,
			City:   m.Address.City,
			State:  m.Address.State,
			Zip:    m.Address.Zip,
		},
	}
	return mt
}
