// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// foobar HTTP server types
//
// Command:
// $ goa gen github.com/jimmiebtlr/goav2-example/design

package server

import (
	foobarviews "github.com/jimmiebtlr/goav2-example/gen/foobar/views"
)

// FoobarResponseBodyCollection is the type of the "foobar" service "list"
// endpoint HTTP response body.
type FoobarResponseBodyCollection []*FoobarResponseBody

// FoobarResponseBody is used to define fields on response body types.
type FoobarResponseBody struct {
	// ID of the service
	ID string `form:"id" json:"id" xml:"id"`
}

// NewFoobarResponseBodyCollection builds the HTTP response body from the
// result of the "list" endpoint of the "foobar" service.
func NewFoobarResponseBodyCollection(res foobarviews.FoobarCollectionView) FoobarResponseBodyCollection {
	body := make([]*FoobarResponseBody, len(res))
	for i, val := range res {
		body[i] = &FoobarResponseBody{
			ID: *val.ID,
		}
	}
	return body
}
