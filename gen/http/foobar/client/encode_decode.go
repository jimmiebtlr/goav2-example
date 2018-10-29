// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// foobar HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/jimmiebtlr/goav2-example/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	foobar "github.com/jimmiebtlr/goav2-example/gen/foobar"
	foobarviews "github.com/jimmiebtlr/goav2-example/gen/foobar/views"
	goahttp "goa.design/goa/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "foobar" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListFoobarPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("foobar", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the foobar
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("foobar", "list", err)
			}
			p := NewListFoobarCollectionOK(body)
			view := "default"
			vres := foobarviews.FoobarCollection{p, view}
			if err = vres.Validate(); err != nil {
				return nil, goahttp.ErrValidationError("foobar", "list", err)
			}
			return foobar.NewFoobarCollection(vres), nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("foobar", "list", resp.StatusCode, string(body))
		}
	}
}