// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// api client HTTP transport
//
// Command:
// $ goa gen github.com/jimmiebtlr/goav2-example/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the api service endpoint HTTP clients.
type Client struct {
	// Info Doer is the HTTP client used to make requests to the info endpoint.
	InfoDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the api service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		InfoDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Info returns an endpoint that makes HTTP requests to the api service info
// server.
func (c *Client) Info() goa.Endpoint {
	var (
		decodeResponse = DecodeInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.InfoDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("api", "info", err)
		}
		return decodeResponse(resp)
	}
}