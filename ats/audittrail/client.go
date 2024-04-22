// This file was auto-generated by Fern from our API Definition.

package audittrail

import (
	context "context"
	fmt "fmt"
	ats "github.com/merge-api/merge-go-client/ats"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
	url "net/url"
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

// Gets a list of audit trail events.
func (c *Client) List(ctx context.Context, request *ats.AuditTrailListRequest) (*ats.PaginatedAuditLogEventList, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ats/v1/audit-trail"

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.EndDate != nil {
		queryParams.Add("end_date", fmt.Sprintf("%v", *request.EndDate))
	}
	if request.EventType != nil {
		queryParams.Add("event_type", fmt.Sprintf("%v", *request.EventType))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.StartDate != nil {
		queryParams.Add("start_date", fmt.Sprintf("%v", *request.StartDate))
	}
	if request.UserEmail != nil {
		queryParams.Add("user_email", fmt.Sprintf("%v", *request.UserEmail))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ats.PaginatedAuditLogEventList
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
