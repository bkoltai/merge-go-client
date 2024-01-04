// This file was auto-generated by Fern from our API Definition.

package accounting

import (
	time "time"
)

type JournalEntryEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool                `json:"-"`
	Model    *JournalEntryRequest `json:"model,omitempty"`
}

type JournalEntriesListRequest struct {
	// If provided, will only return journal entries for this company.
	CompanyId *string `json:"-"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JournalEntriesListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return objects created after this datetime.
	TransactionDateAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	TransactionDateBefore *time.Time `json:"-"`
}

type JournalEntriesRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JournalEntriesRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}
