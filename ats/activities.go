// This file was auto-generated by Fern from our API Definition.

package ats

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
	time "time"
)

type ActivityEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync     *bool            `json:"-"`
	Model        *ActivityRequest `json:"model,omitempty"`
	RemoteUserId string           `json:"remote_user_id"`
}

type ActivitiesListRequest struct {
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
	// Deprecated. Use show_enum_origins.
	RemoteFields *ActivitiesListRequestRemoteFields `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *ActivitiesListRequestShowEnumOrigins `json:"-"`
	// If provided, will only return activities done by this user.
	UserId *string `json:"-"`
}

type ActivitiesRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *string `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *ActivitiesRetrieveRequestRemoteFields `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *ActivitiesRetrieveRequestShowEnumOrigins `json:"-"`
}

type ActivitiesListRequestRemoteFields uint

const (
	ActivitiesListRequestRemoteFieldsActivityType ActivitiesListRequestRemoteFields = iota + 1
	ActivitiesListRequestRemoteFieldsActivityTypeVisibility
	ActivitiesListRequestRemoteFieldsVisibility
)

func (a ActivitiesListRequestRemoteFields) String() string {
	switch a {
	default:
		return strconv.Itoa(int(a))
	case ActivitiesListRequestRemoteFieldsActivityType:
		return "activity_type"
	case ActivitiesListRequestRemoteFieldsActivityTypeVisibility:
		return "activity_type,visibility"
	case ActivitiesListRequestRemoteFieldsVisibility:
		return "visibility"
	}
}

func (a ActivitiesListRequestRemoteFields) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *ActivitiesListRequestRemoteFields) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "activity_type":
		value := ActivitiesListRequestRemoteFieldsActivityType
		*a = value
	case "activity_type,visibility":
		value := ActivitiesListRequestRemoteFieldsActivityTypeVisibility
		*a = value
	case "visibility":
		value := ActivitiesListRequestRemoteFieldsVisibility
		*a = value
	}
	return nil
}

type ActivitiesListRequestShowEnumOrigins uint

const (
	ActivitiesListRequestShowEnumOriginsActivityType ActivitiesListRequestShowEnumOrigins = iota + 1
	ActivitiesListRequestShowEnumOriginsActivityTypeVisibility
	ActivitiesListRequestShowEnumOriginsVisibility
)

func (a ActivitiesListRequestShowEnumOrigins) String() string {
	switch a {
	default:
		return strconv.Itoa(int(a))
	case ActivitiesListRequestShowEnumOriginsActivityType:
		return "activity_type"
	case ActivitiesListRequestShowEnumOriginsActivityTypeVisibility:
		return "activity_type,visibility"
	case ActivitiesListRequestShowEnumOriginsVisibility:
		return "visibility"
	}
}

func (a ActivitiesListRequestShowEnumOrigins) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *ActivitiesListRequestShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "activity_type":
		value := ActivitiesListRequestShowEnumOriginsActivityType
		*a = value
	case "activity_type,visibility":
		value := ActivitiesListRequestShowEnumOriginsActivityTypeVisibility
		*a = value
	case "visibility":
		value := ActivitiesListRequestShowEnumOriginsVisibility
		*a = value
	}
	return nil
}

type ActivitiesRetrieveRequestRemoteFields uint

const (
	ActivitiesRetrieveRequestRemoteFieldsActivityType ActivitiesRetrieveRequestRemoteFields = iota + 1
	ActivitiesRetrieveRequestRemoteFieldsActivityTypeVisibility
	ActivitiesRetrieveRequestRemoteFieldsVisibility
)

func (a ActivitiesRetrieveRequestRemoteFields) String() string {
	switch a {
	default:
		return strconv.Itoa(int(a))
	case ActivitiesRetrieveRequestRemoteFieldsActivityType:
		return "activity_type"
	case ActivitiesRetrieveRequestRemoteFieldsActivityTypeVisibility:
		return "activity_type,visibility"
	case ActivitiesRetrieveRequestRemoteFieldsVisibility:
		return "visibility"
	}
}

func (a ActivitiesRetrieveRequestRemoteFields) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *ActivitiesRetrieveRequestRemoteFields) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "activity_type":
		value := ActivitiesRetrieveRequestRemoteFieldsActivityType
		*a = value
	case "activity_type,visibility":
		value := ActivitiesRetrieveRequestRemoteFieldsActivityTypeVisibility
		*a = value
	case "visibility":
		value := ActivitiesRetrieveRequestRemoteFieldsVisibility
		*a = value
	}
	return nil
}

type ActivitiesRetrieveRequestShowEnumOrigins uint

const (
	ActivitiesRetrieveRequestShowEnumOriginsActivityType ActivitiesRetrieveRequestShowEnumOrigins = iota + 1
	ActivitiesRetrieveRequestShowEnumOriginsActivityTypeVisibility
	ActivitiesRetrieveRequestShowEnumOriginsVisibility
)

func (a ActivitiesRetrieveRequestShowEnumOrigins) String() string {
	switch a {
	default:
		return strconv.Itoa(int(a))
	case ActivitiesRetrieveRequestShowEnumOriginsActivityType:
		return "activity_type"
	case ActivitiesRetrieveRequestShowEnumOriginsActivityTypeVisibility:
		return "activity_type,visibility"
	case ActivitiesRetrieveRequestShowEnumOriginsVisibility:
		return "visibility"
	}
}

func (a ActivitiesRetrieveRequestShowEnumOrigins) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *ActivitiesRetrieveRequestShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "activity_type":
		value := ActivitiesRetrieveRequestShowEnumOriginsActivityType
		*a = value
	case "activity_type,visibility":
		value := ActivitiesRetrieveRequestShowEnumOriginsActivityTypeVisibility
		*a = value
	case "visibility":
		value := ActivitiesRetrieveRequestShowEnumOriginsVisibility
		*a = value
	}
	return nil
}
