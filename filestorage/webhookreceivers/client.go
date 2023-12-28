// This file was auto-generated by Fern from our API Definition.

package webhookreceivers

import (
	context "context"
	core "github.com/merge-api/merge-go-client/core"
	filestorage "github.com/merge-api/merge-go-client/filestorage"
	http "net/http"
)

type Client interface {
	List(ctx context.Context) ([]*filestorage.WebhookReceiver, error)
	Create(ctx context.Context, request *filestorage.WebhookReceiverRequest) (*filestorage.WebhookReceiver, error)
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

// Returns a list of `WebhookReceiver` objects.
func (c *client) List(ctx context.Context) ([]*filestorage.WebhookReceiver, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "webhook-receivers"

	var response []*filestorage.WebhookReceiver
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

// Creates a `WebhookReceiver` object with the given values.
func (c *client) Create(ctx context.Context, request *filestorage.WebhookReceiverRequest) (*filestorage.WebhookReceiver, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "webhook-receivers"

	var response *filestorage.WebhookReceiver
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
