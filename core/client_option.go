// This file was auto-generated by Fern from our API Definition.

package core

import (
	fmt "fmt"
	http "net/http"
)

// ClientOption adapts the behavior of the generated client.
type ClientOption func(*ClientOptions)

// ClientOptions defines all of the possible client options.
// This type is primarily used by the generated code and is
// not meant to be used directly; use ClientOption instead.
type ClientOptions struct {
	BaseURL      string
	HTTPClient   HTTPClient
	HTTPHeader   http.Header
	ApiKey       string
	AccountToken *string
}

// NewClientOptions returns a new *ClientOptions value.
// This function is primarily used by the generated code and is
// not meant to be used directly; use ClientOption instead.
func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		HTTPClient: http.DefaultClient,
		HTTPHeader: make(http.Header),
	}
}

// ToHeader maps the configured client options into a http.Header issued
// on every request.
func (c *ClientOptions) ToHeader() http.Header {
	header := c.cloneHeader()
	if c.ApiKey != "" {
		header.Set("Authorization", "Bearer "+c.ApiKey)
	}
	if c.AccountToken != nil {
		header.Set("X-Account-Token", fmt.Sprintf("%v", *c.AccountToken))
	}
	return header
}

func (c *ClientOptions) cloneHeader() http.Header {
	headers := c.HTTPHeader.Clone()
	headers.Set("X-Fern-Language", "Go")
	headers.Set("X-Fern-SDK-Name", "github.com/merge-api/merge-go-client")
	headers.Set("X-Fern-SDK-Version", "v1.0.7")
	return headers
}
