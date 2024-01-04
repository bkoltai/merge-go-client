// This file was auto-generated by Fern from our API Definition.

package asyncpassthrough

import (
	context "context"
	fmt "fmt"
	accounting "github.com/merge-api/merge-go-client/accounting"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client interface {
	Create(ctx context.Context, request *accounting.DataPassthroughRequest) (*accounting.AsyncPassthroughReciept, error)
	Retrieve(ctx context.Context, asyncPassthroughReceiptId string) (*accounting.RemoteResponse, error)
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Asynchronously pull data from an endpoint not currently supported by Merge.
func (c *client) Create(ctx context.Context, request *accounting.DataPassthroughRequest) (*accounting.AsyncPassthroughReciept, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/accounting/v1/async-passthrough"

	var response *accounting.AsyncPassthroughReciept
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Retrieves data from earlier async-passthrough POST request
func (c *client) Retrieve(ctx context.Context, asyncPassthroughReceiptId string) (*accounting.RemoteResponse, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/accounting/v1/async-passthrough/%v", asyncPassthroughReceiptId)

	var response *accounting.RemoteResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
