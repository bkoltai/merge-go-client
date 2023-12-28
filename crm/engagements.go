// This file was auto-generated by Fern from our API Definition.

package crm

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
	time "time"
)

type EngagementEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool              `json:"-"`
	Model    *EngagementRequest `json:"model,omitempty"`
}

type EngagementsListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *EngagementsListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return engagements started after this datetime.
	StartedAfter *time.Time `json:"-"`
	// If provided, will only return engagements started before this datetime.
	StartedBefore *time.Time `json:"-"`
}

type PatchedEngagementEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool                     `json:"-"`
	Model    *PatchedEngagementRequest `json:"model,omitempty"`
}

type EngagementsRemoteFieldClassesListRequest struct {
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
}

type EngagementsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *EngagementsRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-"`
}

type EngagementsListRequestExpand uint

const (
	EngagementsListRequestExpandAccount EngagementsListRequestExpand = iota + 1
	EngagementsListRequestExpandAccountEngagementType
	EngagementsListRequestExpandContacts
	EngagementsListRequestExpandContactsAccount
	EngagementsListRequestExpandContactsAccountEngagementType
	EngagementsListRequestExpandContactsEngagementType
	EngagementsListRequestExpandContactsOwner
	EngagementsListRequestExpandContactsOwnerAccount
	EngagementsListRequestExpandContactsOwnerAccountEngagementType
	EngagementsListRequestExpandContactsOwnerEngagementType
	EngagementsListRequestExpandEngagementType
	EngagementsListRequestExpandOwner
	EngagementsListRequestExpandOwnerAccount
	EngagementsListRequestExpandOwnerAccountEngagementType
	EngagementsListRequestExpandOwnerEngagementType
)

func (e EngagementsListRequestExpand) String() string {
	switch e {
	default:
		return strconv.Itoa(int(e))
	case EngagementsListRequestExpandAccount:
		return "account"
	case EngagementsListRequestExpandAccountEngagementType:
		return "account,engagement_type"
	case EngagementsListRequestExpandContacts:
		return "contacts"
	case EngagementsListRequestExpandContactsAccount:
		return "contacts,account"
	case EngagementsListRequestExpandContactsAccountEngagementType:
		return "contacts,account,engagement_type"
	case EngagementsListRequestExpandContactsEngagementType:
		return "contacts,engagement_type"
	case EngagementsListRequestExpandContactsOwner:
		return "contacts,owner"
	case EngagementsListRequestExpandContactsOwnerAccount:
		return "contacts,owner,account"
	case EngagementsListRequestExpandContactsOwnerAccountEngagementType:
		return "contacts,owner,account,engagement_type"
	case EngagementsListRequestExpandContactsOwnerEngagementType:
		return "contacts,owner,engagement_type"
	case EngagementsListRequestExpandEngagementType:
		return "engagement_type"
	case EngagementsListRequestExpandOwner:
		return "owner"
	case EngagementsListRequestExpandOwnerAccount:
		return "owner,account"
	case EngagementsListRequestExpandOwnerAccountEngagementType:
		return "owner,account,engagement_type"
	case EngagementsListRequestExpandOwnerEngagementType:
		return "owner,engagement_type"
	}
}

func (e EngagementsListRequestExpand) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", e.String())), nil
}

func (e *EngagementsListRequestExpand) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "account":
		value := EngagementsListRequestExpandAccount
		*e = value
	case "account,engagement_type":
		value := EngagementsListRequestExpandAccountEngagementType
		*e = value
	case "contacts":
		value := EngagementsListRequestExpandContacts
		*e = value
	case "contacts,account":
		value := EngagementsListRequestExpandContactsAccount
		*e = value
	case "contacts,account,engagement_type":
		value := EngagementsListRequestExpandContactsAccountEngagementType
		*e = value
	case "contacts,engagement_type":
		value := EngagementsListRequestExpandContactsEngagementType
		*e = value
	case "contacts,owner":
		value := EngagementsListRequestExpandContactsOwner
		*e = value
	case "contacts,owner,account":
		value := EngagementsListRequestExpandContactsOwnerAccount
		*e = value
	case "contacts,owner,account,engagement_type":
		value := EngagementsListRequestExpandContactsOwnerAccountEngagementType
		*e = value
	case "contacts,owner,engagement_type":
		value := EngagementsListRequestExpandContactsOwnerEngagementType
		*e = value
	case "engagement_type":
		value := EngagementsListRequestExpandEngagementType
		*e = value
	case "owner":
		value := EngagementsListRequestExpandOwner
		*e = value
	case "owner,account":
		value := EngagementsListRequestExpandOwnerAccount
		*e = value
	case "owner,account,engagement_type":
		value := EngagementsListRequestExpandOwnerAccountEngagementType
		*e = value
	case "owner,engagement_type":
		value := EngagementsListRequestExpandOwnerEngagementType
		*e = value
	}
	return nil
}

type EngagementsRetrieveRequestExpand uint

const (
	EngagementsRetrieveRequestExpandAccount EngagementsRetrieveRequestExpand = iota + 1
	EngagementsRetrieveRequestExpandAccountEngagementType
	EngagementsRetrieveRequestExpandContacts
	EngagementsRetrieveRequestExpandContactsAccount
	EngagementsRetrieveRequestExpandContactsAccountEngagementType
	EngagementsRetrieveRequestExpandContactsEngagementType
	EngagementsRetrieveRequestExpandContactsOwner
	EngagementsRetrieveRequestExpandContactsOwnerAccount
	EngagementsRetrieveRequestExpandContactsOwnerAccountEngagementType
	EngagementsRetrieveRequestExpandContactsOwnerEngagementType
	EngagementsRetrieveRequestExpandEngagementType
	EngagementsRetrieveRequestExpandOwner
	EngagementsRetrieveRequestExpandOwnerAccount
	EngagementsRetrieveRequestExpandOwnerAccountEngagementType
	EngagementsRetrieveRequestExpandOwnerEngagementType
)

func (e EngagementsRetrieveRequestExpand) String() string {
	switch e {
	default:
		return strconv.Itoa(int(e))
	case EngagementsRetrieveRequestExpandAccount:
		return "account"
	case EngagementsRetrieveRequestExpandAccountEngagementType:
		return "account,engagement_type"
	case EngagementsRetrieveRequestExpandContacts:
		return "contacts"
	case EngagementsRetrieveRequestExpandContactsAccount:
		return "contacts,account"
	case EngagementsRetrieveRequestExpandContactsAccountEngagementType:
		return "contacts,account,engagement_type"
	case EngagementsRetrieveRequestExpandContactsEngagementType:
		return "contacts,engagement_type"
	case EngagementsRetrieveRequestExpandContactsOwner:
		return "contacts,owner"
	case EngagementsRetrieveRequestExpandContactsOwnerAccount:
		return "contacts,owner,account"
	case EngagementsRetrieveRequestExpandContactsOwnerAccountEngagementType:
		return "contacts,owner,account,engagement_type"
	case EngagementsRetrieveRequestExpandContactsOwnerEngagementType:
		return "contacts,owner,engagement_type"
	case EngagementsRetrieveRequestExpandEngagementType:
		return "engagement_type"
	case EngagementsRetrieveRequestExpandOwner:
		return "owner"
	case EngagementsRetrieveRequestExpandOwnerAccount:
		return "owner,account"
	case EngagementsRetrieveRequestExpandOwnerAccountEngagementType:
		return "owner,account,engagement_type"
	case EngagementsRetrieveRequestExpandOwnerEngagementType:
		return "owner,engagement_type"
	}
}

func (e EngagementsRetrieveRequestExpand) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", e.String())), nil
}

func (e *EngagementsRetrieveRequestExpand) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "account":
		value := EngagementsRetrieveRequestExpandAccount
		*e = value
	case "account,engagement_type":
		value := EngagementsRetrieveRequestExpandAccountEngagementType
		*e = value
	case "contacts":
		value := EngagementsRetrieveRequestExpandContacts
		*e = value
	case "contacts,account":
		value := EngagementsRetrieveRequestExpandContactsAccount
		*e = value
	case "contacts,account,engagement_type":
		value := EngagementsRetrieveRequestExpandContactsAccountEngagementType
		*e = value
	case "contacts,engagement_type":
		value := EngagementsRetrieveRequestExpandContactsEngagementType
		*e = value
	case "contacts,owner":
		value := EngagementsRetrieveRequestExpandContactsOwner
		*e = value
	case "contacts,owner,account":
		value := EngagementsRetrieveRequestExpandContactsOwnerAccount
		*e = value
	case "contacts,owner,account,engagement_type":
		value := EngagementsRetrieveRequestExpandContactsOwnerAccountEngagementType
		*e = value
	case "contacts,owner,engagement_type":
		value := EngagementsRetrieveRequestExpandContactsOwnerEngagementType
		*e = value
	case "engagement_type":
		value := EngagementsRetrieveRequestExpandEngagementType
		*e = value
	case "owner":
		value := EngagementsRetrieveRequestExpandOwner
		*e = value
	case "owner,account":
		value := EngagementsRetrieveRequestExpandOwnerAccount
		*e = value
	case "owner,account,engagement_type":
		value := EngagementsRetrieveRequestExpandOwnerAccountEngagementType
		*e = value
	case "owner,engagement_type":
		value := EngagementsRetrieveRequestExpandOwnerEngagementType
		*e = value
	}
	return nil
}
