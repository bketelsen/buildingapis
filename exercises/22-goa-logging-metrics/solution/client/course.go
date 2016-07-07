package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"time"
)

// CreateCoursePayload is the course create action payload.
type CreateCoursePayload struct {
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

// CreateCoursePath computes a request path to the create action of course.
func CreateCoursePath() string {
	return fmt.Sprintf("/api/courses")
}

// Create a new course
func (c *Client) CreateCourse(ctx context.Context, path string, payload *CreateCoursePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCourseRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCourseRequest create the request corresponding to the create action endpoint of the course resource.
func (c *Client) NewCreateCourseRequest(ctx context.Context, path string, payload *CreateCoursePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteCoursePath computes a request path to the delete action of course.
func DeleteCoursePath(id int) string {
	return fmt.Sprintf("/api/courses/%v", id)
}

// Delete a course
func (c *Client) DeleteCourse(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCourseRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCourseRequest create the request corresponding to the delete action endpoint of the course resource.
func (c *Client) NewDeleteCourseRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowCoursePath computes a request path to the show action of course.
func ShowCoursePath(id int) string {
	return fmt.Sprintf("/api/courses/%v", id)
}

// Retrieve a course by ID
func (c *Client) ShowCourse(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowCourseRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowCourseRequest create the request corresponding to the show action endpoint of the course resource.
func (c *Client) NewShowCourseRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
