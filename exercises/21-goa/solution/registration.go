package main

import (
	"strconv"
	"strings"

	"github.com/bketelsen/buildingapis/exercises/21-goa/solution/app"
	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/goadesign/goa"
)

// RegistrationController implements the registration resource.
type RegistrationController struct {
	*goa.Controller
	db *library.MemDB
}

// NewRegistrationController creates a registration controller.
func NewRegistrationController(service *goa.Service, db *library.MemDB) *RegistrationController {
	return &RegistrationController{
		Controller: service.NewController("RegistrationController"),
		db:         db,
	}
}

// List runs the List action.
func (c *RegistrationController) List(ctx *app.ListRegistrationContext) error {
	es, err := c.db.List("registrations", "id", nil)
	if err != nil {
		return err // internal error
	}
	res := make(app.RegistrationMediaCollection, len(es))
	for i, e := range es {
		res[i] = registrationToMedia(e)
	}
	return ctx.OK(res)
}

// Show runs the Show action.
func (c *RegistrationController) Show(ctx *app.ShowRegistrationContext) error {
	im, err := c.db.Get("registrations", "id", strconv.Itoa(ctx.ID))
	if err != nil && err != library.ErrNotFound {
		return err // internal error
	}
	if im == nil {
		return ctx.NotFound()
	}
	return ctx.OK(registrationToMedia(im))
}

// Create runs the Create action.
func (c *RegistrationController) Create(ctx *app.CreateRegistrationContext) error {
	payload := ctx.Payload
	model := &library.RegistrationModel{
		ID:        library.NewID(),
		CourseID:  courseIDFromHref(payload.CourseHref),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Address:   addressFromPayload(payload.Address),
	}
	if err := c.db.Insert("registrations", model); err != nil {
		return err // internal error
	}
	ctx.ResponseData.Header().Set("Location", app.RegistrationHref(model.ID))
	return ctx.Created(registrationToMedia(model))
}

// courseIDFromHref returns the ID of a course model given a resource href.
func courseIDFromHref(href string) string {
	elems := strings.Split(href, "/")
	return elems[len(elems)-1]
}

// addressFromPayload creates an address model.
func addressFromPayload(payload *app.Address) *library.Address {
	return &library.Address{
		Number: payload.Number,
		Street: payload.Street,
		City:   payload.City,
		State:  payload.State,
		Zip:    payload.Zip,
	}
}

func registrationToMedia(i interface{}) *app.RegistrationMedia {
	m := i.(*library.RegistrationModel)
	id, err := strconv.Atoi(m.ID)
	if err != nil {
		panic("invalid registration ID - must be an int") // bug
	}
	courseID, err := strconv.Atoi(m.CourseID)
	if err != nil {
		panic("invalid course ID - must be an int") // bug
	}
	mt := &app.RegistrationMedia{
		ID:        id,
		Href:      app.RegistrationHref(id),
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Address: &app.Address{
			Number: m.Address.Number,
			Street: m.Address.Street,
			City:   m.Address.City,
			State:  m.Address.State,
			Zip:    m.Address.Zip,
		},
		Links: &app.RegistrationMediaLinks{
			Course: &app.CourseMediaLink{
				Href: app.CourseHref(m.CourseID),
				ID:   courseID,
			},
		},
	}
	return mt
}
