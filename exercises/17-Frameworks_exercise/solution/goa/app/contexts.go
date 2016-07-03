//************************************************************************//
// API "GoWorkshop": Application Contexts
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa/design
// --out=$(GOPATH)/src/github.com/bketelsen/buildingapis/exercises/16-Frameworks_exercise/solution/goa
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

// CreateCourseContext provides the course create action context.
type CreateCourseContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateCoursePayload
}

// NewCreateCourseContext parses the incoming request URL and body, performs validations and creates the
// context used by the course controller create action.
func NewCreateCourseContext(ctx context.Context, service *goa.Service) (*CreateCourseContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateCourseContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createCoursePayload is the course create action payload.
type createCoursePayload struct {
	// Course description
	Description *string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	// Course end date/time
	EndTime *time.Time `json:"end_time,omitempty" xml:"end_time,omitempty" form:"end_time,omitempty"`
	// Course location
	Location *string `json:"location,omitempty" xml:"location,omitempty" form:"location,omitempty"`
	// Course name
	Name *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	// Course start date/time
	StartTime *time.Time `json:"start_time,omitempty" xml:"start_time,omitempty" form:"start_time,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createCoursePayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.StartTime == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "start_time"))
	}
	if payload.EndTime == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "end_time"))
	}
	if payload.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "location"))
	}

	if payload.Location != nil {
		if len(*payload.Location) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.location`, *payload.Location, len(*payload.Location), 2, true))
		}
	}
	if payload.Name != nil {
		if len(*payload.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 3, true))
		}
	}
	return
}

// Publicize creates CreateCoursePayload from createCoursePayload
func (payload *createCoursePayload) Publicize() *CreateCoursePayload {
	var pub CreateCoursePayload
	if payload.Description != nil {
		pub.Description = payload.Description
	}
	if payload.EndTime != nil {
		pub.EndTime = *payload.EndTime
	}
	if payload.Location != nil {
		pub.Location = *payload.Location
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.StartTime != nil {
		pub.StartTime = *payload.StartTime
	}
	return &pub
}

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

// Validate runs the validation rules defined in the design.
func (payload *CreateCoursePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "location"))
	}

	if len(payload.Location) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.location`, payload.Location, len(payload.Location), 2, true))
	}
	if len(payload.Name) < 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 3, true))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateCourseContext) Created(r *CourseMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goworkshop.course+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateCourseContext) BadRequest(r *goa.Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.api.error+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// DeleteCourseContext provides the course delete action context.
type DeleteCourseContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewDeleteCourseContext parses the incoming request URL and body, performs validations and creates the
// context used by the course controller delete action.
func NewDeleteCourseContext(ctx context.Context, service *goa.Service) (*DeleteCourseContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteCourseContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
		if rctx.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`id`, rctx.ID, 1, true))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteCourseContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteCourseContext) BadRequest(r *goa.Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.api.error+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteCourseContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowCourseContext provides the course show action context.
type ShowCourseContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowCourseContext parses the incoming request URL and body, performs validations and creates the
// context used by the course controller show action.
func NewShowCourseContext(ctx context.Context, service *goa.Service) (*ShowCourseContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowCourseContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
		if rctx.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`id`, rctx.ID, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCourseContext) OK(r *CourseMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goworkshop.course+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowCourseContext) BadRequest(r *goa.Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.api.error+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCourseContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateRegistrationContext provides the registration create action context.
type CreateRegistrationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateRegistrationPayload
}

// NewCreateRegistrationContext parses the incoming request URL and body, performs validations and creates the
// context used by the registration controller create action.
func NewCreateRegistrationContext(ctx context.Context, service *goa.Service) (*CreateRegistrationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateRegistrationContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createRegistrationPayload is the registration create action payload.
type createRegistrationPayload struct {
	// Attendee address
	Address *address `json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	// The href to the course resource that describes the course being taught
	CourseHref *string `json:"course_href,omitempty" xml:"course_href,omitempty" form:"course_href,omitempty"`
	// Attendee first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty" form:"first_name,omitempty"`
	// Attendee last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty" form:"last_name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createRegistrationPayload) Validate() (err error) {
	if payload.CourseHref == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "course_href"))
	}
	if payload.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}
	if payload.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "address"))
	}

	if payload.Address != nil {
		if err2 := payload.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.CourseHref != nil {
		if ok := goa.ValidatePattern(`/courses/[0-9]+`, *payload.CourseHref); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.course_href`, *payload.CourseHref, `/courses/[0-9]+`))
		}
	}
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	return
}

// Publicize creates CreateRegistrationPayload from createRegistrationPayload
func (payload *createRegistrationPayload) Publicize() *CreateRegistrationPayload {
	var pub CreateRegistrationPayload
	if payload.Address != nil {
		pub.Address = payload.Address.Publicize()
	}
	if payload.CourseHref != nil {
		pub.CourseHref = *payload.CourseHref
	}
	if payload.FirstName != nil {
		pub.FirstName = *payload.FirstName
	}
	if payload.LastName != nil {
		pub.LastName = *payload.LastName
	}
	return &pub
}

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

// Validate runs the validation rules defined in the design.
func (payload *CreateRegistrationPayload) Validate() (err error) {
	if payload.CourseHref == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "course_href"))
	}
	if payload.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}
	if payload.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "address"))
	}

	if payload.Address != nil {
		if err2 := payload.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`/courses/[0-9]+`, payload.CourseHref); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.course_href`, payload.CourseHref, `/courses/[0-9]+`))
	}
	if len(payload.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, payload.FirstName, len(payload.FirstName), 2, true))
	}
	if len(payload.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, payload.LastName, len(payload.LastName), 2, true))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateRegistrationContext) Created(r *RegistrationMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goworkshop.registration+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateRegistrationContext) BadRequest(r *goa.Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.api.error+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// ListRegistrationContext provides the registration list action context.
type ListRegistrationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListRegistrationContext parses the incoming request URL and body, performs validations and creates the
// context used by the registration controller list action.
func NewListRegistrationContext(ctx context.Context, service *goa.Service) (*ListRegistrationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListRegistrationContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListRegistrationContext) OK(r RegistrationMediaCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goworkshop.registration+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowRegistrationContext provides the registration show action context.
type ShowRegistrationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowRegistrationContext parses the incoming request URL and body, performs validations and creates the
// context used by the registration controller show action.
func NewShowRegistrationContext(ctx context.Context, service *goa.Service) (*ShowRegistrationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowRegistrationContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
		if rctx.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`id`, rctx.ID, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRegistrationContext) OK(r *RegistrationMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goworkshop.registration+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowRegistrationContext) BadRequest(r *goa.Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.api.error+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRegistrationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
