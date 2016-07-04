package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/labstack/echo"
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

func createRegistration(db *library.MemDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload registrationPayload
		err := c.Bind(&payload)
		if err != nil {
			return respondBadRequest(c, "invalid request body: %s", err)
		}
		model := &library.RegistrationModel{
			ID:        library.NewID(),
			CourseID:  courseIDFromHref(payload.CourseHref),
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Address:   addressFromPayload(payload.Address),
		}
		if err := db.Insert("registrations", model); err != nil {
			return fmt.Errorf("failed to insert registration: %s", err)
		}
		c.Response().Header().Set("Location", fmt.Sprintf("/registrations/%s", model.ID))
		return c.JSON(http.StatusCreated, registrationToMedia(model))
	}
}

func showRegistration(db *library.MemDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		im, err := db.Get("registrations", "id", id)
		if err != nil && err != library.ErrNotFound {
			return fmt.Errorf("failed to query registration: %s", err)
		}
		if im == nil {
			return respondNotFound(c)
		}
		return c.JSON(http.StatusOK, registrationToMedia(im.(*library.RegistrationModel)))
	}
}

func listRegistrations(db *library.MemDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		im, err := db.List("registrations", "id", nil)
		if err != nil && err != library.ErrNotFound {
			return fmt.Errorf("failed to query registrations: %s", err)
		}
		if im == nil {
			return respondNotFound(c)
		}
		med := make([]*registrationMedia, len(im))
		for i, m := range im {
			med[i] = registrationToMedia(m.(*library.RegistrationModel))
		}
		return c.JSON(http.StatusOK, med)
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
