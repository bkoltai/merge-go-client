// This file was auto-generated by Fern from our API Definition.

package forceresync

import (
	context "context"
	accounting "github.com/merge-api/merge-go-client/accounting"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Force re-sync of all models. This is available for all organizations via the dashboard. Force re-sync is also available programmatically via API for monthly, quarterly, and highest sync frequency customers on the Launch, Professional, or Enterprise plans. Doing so will consume a sync credit for the relevant linked account.
func (c *Client) SyncStatusResyncCreate(ctx context.Context) ([]*accounting.SyncStatus, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/accounting/v1/sync-status/resync"

	var response []*accounting.SyncStatus
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
