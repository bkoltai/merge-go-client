// This file was auto-generated by Fern from our API Definition.

package interviews

import (
	context "context"
	fmt "fmt"
	ats "github.com/merge-api/merge-go-client/ats"
	core "github.com/merge-api/merge-go-client/core"
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

// Returns a list of `ScheduledInterview` objects.
func (c *Client) List(ctx context.Context, request *ats.InterviewsListRequest) (*ats.PaginatedScheduledInterviewList, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ats/v1/interviews"

	queryParams := make(url.Values)
	if request.ApplicationId != nil {
		queryParams.Add("application_id", fmt.Sprintf("%v", *request.ApplicationId))
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
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.JobId != nil {
		queryParams.Add("job_id", fmt.Sprintf("%v", *request.JobId))
	}
	if request.JobInterviewStageId != nil {
		queryParams.Add("job_interview_stage_id", fmt.Sprintf("%v", *request.JobInterviewStageId))
	}
	if request.ModifiedAfter != nil {
		queryParams.Add("modified_after", fmt.Sprintf("%v", request.ModifiedAfter.Format(time.RFC3339)))
	}
	if request.ModifiedBefore != nil {
		queryParams.Add("modified_before", fmt.Sprintf("%v", request.ModifiedBefore.Format(time.RFC3339)))
	}
	if request.OrganizerId != nil {
		queryParams.Add("organizer_id", fmt.Sprintf("%v", *request.OrganizerId))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.RemoteFields != nil {
		queryParams.Add("remote_fields", fmt.Sprintf("%v", request.RemoteFields))
	}
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if request.ShowEnumOrigins != nil {
		queryParams.Add("show_enum_origins", fmt.Sprintf("%v", request.ShowEnumOrigins))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ats.PaginatedScheduledInterviewList
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

// Creates a `ScheduledInterview` object with the given values.
func (c *Client) Create(ctx context.Context, request *ats.ScheduledInterviewEndpointRequest) (*ats.ScheduledInterviewResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ats/v1/interviews"

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

	var response *ats.ScheduledInterviewResponse
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

// Returns a `ScheduledInterview` object with the given `id`.
func (c *Client) Retrieve(ctx context.Context, id string, request *ats.InterviewsRetrieveRequest) (*ats.ScheduledInterview, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"ats/v1/interviews/%v", id)

	queryParams := make(url.Values)
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.RemoteFields != nil {
		queryParams.Add("remote_fields", fmt.Sprintf("%v", request.RemoteFields))
	}
	if request.ShowEnumOrigins != nil {
		queryParams.Add("show_enum_origins", fmt.Sprintf("%v", request.ShowEnumOrigins))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ats.ScheduledInterview
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

// Returns metadata for `ScheduledInterview` POSTs.
func (c *Client) MetaPostRetrieve(ctx context.Context) (*ats.MetaResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ats/v1/interviews/meta/post"

	var response *ats.MetaResponse
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
