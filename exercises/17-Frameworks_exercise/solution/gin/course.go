package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bketelsen/buildingapis/exercises/library"
	"github.com/gin-gonic/gin"
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

func createCourse(db *library.MemDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload coursePayload
		err := c.BindJSON(&payload)
		if err != nil {
			respondBadRequest(c, "invalid request body: %s", err)
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
			respondInternal(c, "failed to insert course: %s", err)
			return
		}
		c.Header("Location", fmt.Sprintf("/courses/%d", model.ID))
		c.JSON(http.StatusCreated, courseToMedia(model))
	}
}

func showCourse(db *library.MemDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		im, err := db.Get("courses", "id", id)
		if err != nil && err != library.ErrNotFound {
			respondInternal(c, "failed to query course: %s", err)
			return
		}
		if im == nil {
			respondNotFound(c)
			return
		}
		c.JSON(http.StatusOK, courseToMedia(im.(*library.CourseModel)))
	}
}

func deleteCourse(db *library.MemDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete("courses", "id", id); err != nil {
			if err == library.ErrNotFound {
				respondNotFound(c)
				return
			}
			respondInternal(c, "failed to delete course: %s", err)
			return
		}
		c.Status(http.StatusNoContent)
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
