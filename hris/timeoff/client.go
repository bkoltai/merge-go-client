// This file was auto-generated by Fern from our API Definition.

package timeoff

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	hris "github.com/merge-api/merge-go-client/hris"
	http "net/http"
	url "net/url"
	time "time"
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

// Returns a list of `TimeOff` objects.
func (c *Client) List(ctx context.Context, request *hris.TimeOffListRequest) (*hris.PaginatedTimeOffList, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "hris/v1/time-off"

	queryParams := make(url.Values)
	if request.ApproverId != nil {
		queryParams.Add("approver_id", fmt.Sprintf("%v", *request.ApproverId))
	}
	if request.CreatedAfter != nil {
		queryParams.Add("created_after", fmt.Sprintf("%v", request.CreatedAfter.Format(time.RFC3339)))
	}
	if request.CreatedBefore != nil {
		queryParams.Add("created_before", fmt.Sprintf("%v", request.CreatedBefore.Format(time.RFC3339)))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.EmployeeId != nil {
		queryParams.Add("employee_id", fmt.Sprintf("%v", *request.EmployeeId))
	}
	if request.EndedAfter != nil {
		queryParams.Add("ended_after", fmt.Sprintf("%v", request.EndedAfter.Format(time.RFC3339)))
	}
	if request.EndedBefore != nil {
		queryParams.Add("ended_before", fmt.Sprintf("%v", request.EndedBefore.Format(time.RFC3339)))
	}
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.ModifiedAfter != nil {
		queryParams.Add("modified_after", fmt.Sprintf("%v", request.ModifiedAfter.Format(time.RFC3339)))
	}
	if request.ModifiedBefore != nil {
		queryParams.Add("modified_before", fmt.Sprintf("%v", request.ModifiedBefore.Format(time.RFC3339)))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.RemoteFields != nil {
		queryParams.Add("remote_fields", fmt.Sprintf("%v", *request.RemoteFields))
	}
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if request.RequestType != nil {
		queryParams.Add("request_type", fmt.Sprintf("%v", *request.RequestType))
	}
	if request.ShowEnumOrigins != nil {
		queryParams.Add("show_enum_origins", fmt.Sprintf("%v", *request.ShowEnumOrigins))
	}
	if request.StartedAfter != nil {
		queryParams.Add("started_after", fmt.Sprintf("%v", request.StartedAfter.Format(time.RFC3339)))
	}
	if request.StartedBefore != nil {
		queryParams.Add("started_before", fmt.Sprintf("%v", request.StartedBefore.Format(time.RFC3339)))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", *request.Status))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hris.PaginatedTimeOffList
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

// Creates a `TimeOff` object with the given values.
func (c *Client) Create(ctx context.Context, request *hris.TimeOffEndpointRequest) (*hris.TimeOffResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "hris/v1/time-off"

	queryParams := make(url.Values)
	if request.IsDebugMode != nil {
		queryParams.Add("is_debug_mode", fmt.Sprintf("%v", *request.IsDebugMode))
	}
	if request.RunAsync != nil {
		queryParams.Add("run_async", fmt.Sprintf("%v", *request.RunAsync))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hris.TimeOffResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns a `TimeOff` object with the given `id`.
func (c *Client) Retrieve(ctx context.Context, id string, request *hris.TimeOffRetrieveRequest) (*hris.TimeOff, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"hris/v1/time-off/%v", id)

	queryParams := make(url.Values)
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.RemoteFields != nil {
		queryParams.Add("remote_fields", fmt.Sprintf("%v", *request.RemoteFields))
	}
	if request.ShowEnumOrigins != nil {
		queryParams.Add("show_enum_origins", fmt.Sprintf("%v", *request.ShowEnumOrigins))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hris.TimeOff
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

// Returns metadata for `TimeOff` POSTs.
func (c *Client) MetaPostRetrieve(ctx context.Context) (*hris.MetaResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "hris/v1/time-off/meta/post"

	var response *hris.MetaResponse
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
