// This file was auto-generated by Fern from our API Definition.

package filestorage

import (
	time "time"
)

type FileStorageFolderEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool          `json:"-"`
	Model    *FolderRequest `json:"model,omitempty"`
}

type FoldersListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return folders in this drive.
	DriveId *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *FoldersListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// If provided, will only return folders with this name. This performs an exact match.
	Name *string `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// If provided, will only return folders in this parent folder. If null, will return folders in root directory.
	ParentFolderId *string `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
}

type FoldersRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *FoldersRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}
