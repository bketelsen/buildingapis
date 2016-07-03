//************************************************************************//
// API "GoWorkshop": Application Controllers
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/bketelsen/buildingapis/workshop/service/design
// --out=$(GOPATH)/src/github.com/bketelsen/buildingapis/workshop/service
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// CourseController is the controller interface for the Course actions.
type CourseController interface {
	goa.Muxer
	Create(*CreateCourseContext) error
	Delete(*DeleteCourseContext) error
	List(*ListCourseContext) error
	Patch(*PatchCourseContext) error
	Show(*ShowCourseContext) error
}

// MountCourseController "mounts" a Course resource controller on the given service.
func MountCourseController(service *goa.Service, ctrl CourseController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCourseContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCoursePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("POST", "/api/courses", ctrl.MuxHandler("Create", h, unmarshalCreateCoursePayload))
	service.LogInfo("mount", "ctrl", "Course", "action", "Create", "route", "POST /api/courses", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCourseContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("DELETE", "/api/courses/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Course", "action", "Delete", "route", "DELETE /api/courses/:id", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCourseContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("GET", "/api/courses", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Course", "action", "List", "route", "GET /api/courses", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPatchCourseContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PatchCoursePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Patch(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("PATCH", "/api/courses/:id", ctrl.MuxHandler("Patch", h, unmarshalPatchCoursePayload))
	service.LogInfo("mount", "ctrl", "Course", "action", "Patch", "route", "PATCH /api/courses/:id", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCourseContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("GET", "/api/courses/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Course", "action", "Show", "route", "GET /api/courses/:id", "security", "JWTAuth")
}

// unmarshalCreateCoursePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCoursePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCoursePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalPatchCoursePayload unmarshals the request body into the context request data Payload field.
func unmarshalPatchCoursePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &patchCoursePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger/*file", "public/")
	service.Mux.Handle("GET", "/swagger/*file", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/", "route", "GET /swagger/*file")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger.yaml", "swagger/swagger.yaml")
	service.Mux.Handle("GET", "/swagger.yaml", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "swagger/swagger.yaml", "route", "GET /swagger.yaml")

	h = ctrl.FileHandler("/swagger/", "public/index.html")
	service.Mux.Handle("GET", "/swagger/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/index.html", "route", "GET /swagger/")
}

// RegistrationController is the controller interface for the Registration actions.
type RegistrationController interface {
	goa.Muxer
	Create(*CreateRegistrationContext) error
	List(*ListRegistrationContext) error
	Patch(*PatchRegistrationContext) error
	Show(*ShowRegistrationContext) error
}

// MountRegistrationController "mounts" a Registration resource controller on the given service.
func MountRegistrationController(service *goa.Service, ctrl RegistrationController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateRegistrationContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateRegistrationPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("POST", "/api/registrations", ctrl.MuxHandler("Create", h, unmarshalCreateRegistrationPayload))
	service.LogInfo("mount", "ctrl", "Registration", "action", "Create", "route", "POST /api/registrations", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRegistrationContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("GET", "/api/registrations", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Registration", "action", "List", "route", "GET /api/registrations", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPatchRegistrationContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PatchRegistrationPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Patch(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("PATCH", "/api/registrations/:id", ctrl.MuxHandler("Patch", h, unmarshalPatchRegistrationPayload))
	service.LogInfo("mount", "ctrl", "Registration", "action", "Patch", "route", "PATCH /api/registrations/:id", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRegistrationContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	service.Mux.Handle("GET", "/api/registrations/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Registration", "action", "Show", "route", "GET /api/registrations/:id", "security", "JWTAuth")
}

// unmarshalCreateRegistrationPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateRegistrationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createRegistrationPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalPatchRegistrationPayload unmarshals the request body into the context request data Payload field.
func unmarshalPatchRegistrationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &patchRegistrationPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SessionController is the controller interface for the Session actions.
type SessionController interface {
	goa.Muxer
	Login(*LoginSessionContext) error
}

// MountSessionController "mounts" a Session resource controller on the given service.
func MountSessionController(service *goa.Service, ctrl SessionController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginSessionContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Login(rctx)
	}
	h = handleSecurity("BasicAuth", h)
	service.Mux.Handle("GET", "/api/token", ctrl.MuxHandler("Login", h, nil))
	service.LogInfo("mount", "ctrl", "Session", "action", "Login", "route", "GET /api/token", "security", "BasicAuth")
}
