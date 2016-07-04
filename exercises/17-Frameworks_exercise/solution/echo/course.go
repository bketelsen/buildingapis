package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/labstack/echo"
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

func createCourse(db *library.MemDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload coursePayload
		err := c.Bind(&payload)
		if err != nil {
			return respondBadRequest(c, "invalid request body: %s", err)
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
			return fmt.Errorf("failed to insert course: %s", err)
		}
		c.Response().Header().Set("Location", fmt.Sprintf("/courses/%s", model.ID))
		return c.JSON(http.StatusCreated, courseToMedia(model))
	}
}

func showCourse(db *library.MemDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		im, err := db.Get("courses", "id", id)
		if err != nil && err != library.ErrNotFound {
			return fmt.Errorf("failed to query course: %s", err)
		}
		if im == nil {
			return respondNotFound(c)
		}
		return c.JSON(http.StatusOK, courseToMedia(im.(*library.CourseModel)))
	}
}

func deleteCourse(db *library.MemDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := db.Delete("courses", "id", id); err != nil {
			if err == library.ErrNotFound {
				return respondNotFound(c)
			}
			return fmt.Errorf("failed to delete course: %s", err)
		}
		return c.NoContent(http.StatusNoContent)
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
