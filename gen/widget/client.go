// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// widget client
//
// Command:
// $ goa gen github.com/jimmiebtlr/goav2-example/design

package widget

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "widget" service client.
type Client struct {
	ListEndpoint goa.Endpoint
}

// NewClient initializes a "widget" service client given the endpoints.
func NewClient(list goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
	}
}

// List calls the "list" endpoint of the "widget" service.
func (c *Client) List(ctx context.Context) (res WidgetCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(WidgetCollection), nil
}
