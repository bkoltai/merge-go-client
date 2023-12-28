// This file was auto-generated by Fern from our API Definition.

package ticketing

import (
	time "time"
)

type TicketingAttachmentEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool              `json:"-"`
	Model    *AttachmentRequest `json:"model,omitempty"`
}

type AttachmentsDownloadRetrieveRequest struct {
	// If provided, specifies the export format of the file to be downloaded. For information on supported export formats, please refer to our <a href='https://help.merge.dev/en/articles/8615316-file-export-and-download-specification' target='_blank'>export format help center article</a>.
	MimeType *string `json:"-"`
}

type AttachmentsListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *string `json:"-"`
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
	// If provided, will only return attachments created in the third party platform after this datetime.
	RemoteCreatedAfter *time.Time `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return comments for this ticket.
	TicketId *string `json:"-"`
}

type AttachmentsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *string `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}
