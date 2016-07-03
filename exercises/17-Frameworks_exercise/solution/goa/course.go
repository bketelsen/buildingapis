package main

import (
	"strconv"

	"github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa/app"
	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/goadesign/goa"
)

// CourseController implements the course resource.
type CourseController struct {
	*goa.Controller
	db *library.MemDB
}

// NewCourseController creates a course controller.
func NewCourseController(service *goa.Service, db *library.MemDB) *CourseController {
	return &CourseController{
		Controller: service.NewController("CourseController"),
		db:         db,
	}
}

// Create runs the create action.
func (c *CourseController) Create(ctx *app.CreateCourseContext) error {
	payload := ctx.Payload
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
	if err := c.db.Insert("courses", model); err != nil {
		return err // internal error
	}
	ctx.ResponseData.Header().Set("Location", app.CourseHref(model.ID))
	return ctx.Created(courseToMedia(model))
}

// Delete runs the delete action.
func (c *CourseController) Delete(ctx *app.DeleteCourseContext) error {
	if err := c.db.Delete("courses", "id", strconv.Itoa(ctx.ID)); err != nil {
		if err == library.ErrNotFound {
			return ctx.NotFound()
		}
		return err // internal error
	}
	return ctx.NoContent()
}

// Show runs the show action.
func (c *CourseController) Show(ctx *app.ShowCourseContext) error {
	im, err := c.db.Get("courses", "id", strconv.Itoa(ctx.ID))
	if err != nil && err != library.ErrNotFound {
		return err // internal error
	}
	if im == nil {
		return ctx.NotFound()
	}
	return ctx.OK(courseToMedia(im))
}

func courseToMedia(i interface{}) *app.CourseMedia {
	m := i.(*library.CourseModel)
	id, err := strconv.Atoi(m.ID)
	if err != nil {
		panic("invalid course ID - must be an int") // bug
	}
	mt := &app.CourseMedia{
		ID:        id,
		Href:      app.CourseHref(id),
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
