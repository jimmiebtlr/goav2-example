// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// api HTTP server
//
// Command:
// $ goa gen github.com/jimmiebtlr/goav2-example/design

package server

import (
	"context"
	"net/http"

	api "github.com/jimmiebtlr/goav2-example/gen/api"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the api service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Info   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the api service endpoints.
func New(
	e *api.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Info", "GET", "/api/info"},
		},
		Info: NewInfoHandler(e.Info, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "api" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Info = m(s.Info)
}

// Mount configures the mux to serve the api endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountInfoHandler(mux, h.Info)
}

// MountInfoHandler configures the mux to serve the "api" service "info"
// endpoint.
func MountInfoHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/info", f)
}

// NewInfoHandler creates a HTTP handler which loads the HTTP request and calls
// the "api" service "info" endpoint.
func NewInfoHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeInfoResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "info")
		ctx = context.WithValue(ctx, goa.ServiceKey, "api")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}