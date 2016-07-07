package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// CreateRegistrationPayload is the registration create action payload.
type CreateRegistrationPayload struct {
	// Attendee address
	Address *Address `json:"address" xml:"address" form:"address"`
	// The href to the course resource that describes the course being taught
	CourseHref string `json:"course_href" xml:"course_href" form:"course_href"`
	// Attendee first name
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	// Attendee last name
	LastName string `json:"last_name" xml:"last_name" form:"last_name"`
}

// CreateRegistrationPath computes a request path to the create action of registration.
func CreateRegistrationPath() string {
	return fmt.Sprintf("/api/registrations")
}

// Create a new registration
func (c *Client) CreateRegistration(ctx context.Context, path string, payload *CreateRegistrationPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateRegistrationRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateRegistrationRequest create the request corresponding to the create action endpoint of the registration resource.
func (c *Client) NewCreateRegistrationRequest(ctx context.Context, path string, payload *CreateRegistrationPayload, contentType string) (*http.Request, error) {
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

// ListRegistrationPath computes a request path to the list action of registration.
func ListRegistrationPath() string {
	return fmt.Sprintf("/api/registrations")
}

// List all registrations
func (c *Client) ListRegistration(ctx context.Context, path string, courseID *int) (*http.Response, error) {
	req, err := c.NewListRegistrationRequest(ctx, path, courseID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRegistrationRequest create the request corresponding to the list action endpoint of the registration resource.
func (c *Client) NewListRegistrationRequest(ctx context.Context, path string, courseID *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if courseID != nil {
		tmp7 := strconv.Itoa(*courseID)
		values.Set("course_id", tmp7)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowRegistrationPath computes a request path to the show action of registration.
func ShowRegistrationPath(id int) string {
	return fmt.Sprintf("/api/registrations/%v", id)
}

// Retrieve a registration by ID
func (c *Client) ShowRegistration(ctx context.Context, path string, view *string) (*http.Response, error) {
	req, err := c.NewShowRegistrationRequest(ctx, path, view)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowRegistrationRequest create the request corresponding to the show action endpoint of the registration resource.
func (c *Client) NewShowRegistrationRequest(ctx context.Context, path string, view *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if view != nil {
		values.Set("view", *view)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
