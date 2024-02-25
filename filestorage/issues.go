// This file was auto-generated by Fern from our API Definition.

package filestorage

import (
	fmt "fmt"
	time "time"
)

type IssuesListRequest struct {
	AccountToken *string `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If included, will only include issues whose most recent action occurred before this time
	EndDate                 *string `json:"-"`
	EndUserOrganizationName *string `json:"-"`
	// If provided, will only return issues whose first incident time was after this datetime.
	FirstIncidentTimeAfter *time.Time `json:"-"`
	// If provided, will only return issues whose first incident time was before this datetime.
	FirstIncidentTimeBefore *time.Time `json:"-"`
	// If true, will include muted issues
	IncludeMuted    *string `json:"-"`
	IntegrationName *string `json:"-"`
	// If provided, will only return issues whose last incident time was after this datetime.
	LastIncidentTimeAfter *time.Time `json:"-"`
	// If provided, will only return issues whose last incident time was before this datetime.
	LastIncidentTimeBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// If included, will only include issues whose most recent action occurred after this time
	StartDate *string `json:"-"`
	// Status of the issue. Options: ('ONGOING', 'RESOLVED')
	//
	// - `ONGOING` - ONGOING
	// - `RESOLVED` - RESOLVED
	Status *IssuesListRequestStatus `json:"-"`
}

type IssuesListRequestStatus string

const (
	IssuesListRequestStatusOngoing  IssuesListRequestStatus = "ONGOING"
	IssuesListRequestStatusResolved IssuesListRequestStatus = "RESOLVED"
)

func NewIssuesListRequestStatusFromString(s string) (IssuesListRequestStatus, error) {
	switch s {
	case "ONGOING":
		return IssuesListRequestStatusOngoing, nil
	case "RESOLVED":
		return IssuesListRequestStatusResolved, nil
	}
	var t IssuesListRequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i IssuesListRequestStatus) Ptr() *IssuesListRequestStatus {
	return &i
}
