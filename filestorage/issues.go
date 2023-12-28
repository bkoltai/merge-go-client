// This file was auto-generated by Fern from our API Definition.

package filestorage

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
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
	// If True, will include muted issues
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

type IssuesListRequestStatus uint

const (
	IssuesListRequestStatusOngoing IssuesListRequestStatus = iota + 1
	IssuesListRequestStatusResolved
)

func (i IssuesListRequestStatus) String() string {
	switch i {
	default:
		return strconv.Itoa(int(i))
	case IssuesListRequestStatusOngoing:
		return "ONGOING"
	case IssuesListRequestStatusResolved:
		return "RESOLVED"
	}
}

func (i IssuesListRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", i.String())), nil
}

func (i *IssuesListRequestStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "ONGOING":
		value := IssuesListRequestStatusOngoing
		*i = value
	case "RESOLVED":
		value := IssuesListRequestStatusResolved
		*i = value
	}
	return nil
}
