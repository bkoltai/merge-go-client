// This file was auto-generated by Fern from our API Definition.

package accounts

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	ticketing "github.com/merge-api/merge-go-client/ticketing"
	http "net/http"
	url "net/url"
	time "time"
)

type Client interface {
	List(ctx context.Context, request *ticketing.AccountsListRequest) (*ticketing.PaginatedAccountList, error)
	Retrieve(ctx context.Context, id string, request *ticketing.AccountsRetrieveRequest) (*ticketing.Account, error)
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

// Returns a list of `Account` objects.
func (c *client) List(ctx context.Context, request *ticketing.AccountsListRequest) (*ticketing.PaginatedAccountList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/ticketing/v1/accounts"

	queryParams := make(url.Values)
	if request.CreatedAfter != nil {
		queryParams.Add("created_after", fmt.Sprintf("%v", request.CreatedAfter.Format(time.RFC3339)))
	}
	if request.CreatedBefore != nil {
		queryParams.Add("created_before", fmt.Sprintf("%v", request.CreatedBefore.Format(time.RFC3339)))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
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
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ticketing.PaginatedAccountList
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
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

// Returns an `Account` object with the given `id`.
func (c *client) Retrieve(ctx context.Context, id string, request *ticketing.AccountsRetrieveRequest) (*ticketing.Account, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/ticketing/v1/accounts/%v", id)

	queryParams := make(url.Values)
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ticketing.Account
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
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
