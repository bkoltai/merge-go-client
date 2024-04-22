// This file was auto-generated by Fern from our API Definition.

package hris

import (
	fmt "fmt"
	time "time"
)

type PayrollRunsListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return payroll runs ended after this datetime.
	EndedAfter *time.Time `json:"-"`
	// If provided, will only return payroll runs ended before this datetime.
	EndedBefore *time.Time `json:"-"`
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
	RemoteFields *PayrollRunsListRequestRemoteFields `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return PayrollRun's with this status. Options: ('REGULAR', 'OFF_CYCLE', 'CORRECTION', 'TERMINATION', 'SIGN_ON_BONUS')
	//
	// - `REGULAR` - REGULAR
	// - `OFF_CYCLE` - OFF_CYCLE
	// - `CORRECTION` - CORRECTION
	// - `TERMINATION` - TERMINATION
	// - `SIGN_ON_BONUS` - SIGN_ON_BONUS
	RunType *PayrollRunsListRequestRunType `json:"-"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *PayrollRunsListRequestShowEnumOrigins `json:"-"`
	// If provided, will only return payroll runs started after this datetime.
	StartedAfter *time.Time `json:"-"`
	// If provided, will only return payroll runs started before this datetime.
	StartedBefore *time.Time `json:"-"`
}

type PayrollRunsRetrieveRequest struct {
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *PayrollRunsRetrieveRequestRemoteFields `json:"-"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *PayrollRunsRetrieveRequestShowEnumOrigins `json:"-"`
}

type PayrollRunsListRequestRemoteFields string

const (
	PayrollRunsListRequestRemoteFieldsRunState        PayrollRunsListRequestRemoteFields = "run_state"
	PayrollRunsListRequestRemoteFieldsRunStateRunType PayrollRunsListRequestRemoteFields = "run_state,run_type"
	PayrollRunsListRequestRemoteFieldsRunType         PayrollRunsListRequestRemoteFields = "run_type"
)

func NewPayrollRunsListRequestRemoteFieldsFromString(s string) (PayrollRunsListRequestRemoteFields, error) {
	switch s {
	case "run_state":
		return PayrollRunsListRequestRemoteFieldsRunState, nil
	case "run_state,run_type":
		return PayrollRunsListRequestRemoteFieldsRunStateRunType, nil
	case "run_type":
		return PayrollRunsListRequestRemoteFieldsRunType, nil
	}
	var t PayrollRunsListRequestRemoteFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PayrollRunsListRequestRemoteFields) Ptr() *PayrollRunsListRequestRemoteFields {
	return &p
}

type PayrollRunsListRequestRunType string

const (
	PayrollRunsListRequestRunTypeCorrection  PayrollRunsListRequestRunType = "CORRECTION"
	PayrollRunsListRequestRunTypeOffCycle    PayrollRunsListRequestRunType = "OFF_CYCLE"
	PayrollRunsListRequestRunTypeRegular     PayrollRunsListRequestRunType = "REGULAR"
	PayrollRunsListRequestRunTypeSignOnBonus PayrollRunsListRequestRunType = "SIGN_ON_BONUS"
	PayrollRunsListRequestRunTypeTermination PayrollRunsListRequestRunType = "TERMINATION"
)

func NewPayrollRunsListRequestRunTypeFromString(s string) (PayrollRunsListRequestRunType, error) {
	switch s {
	case "CORRECTION":
		return PayrollRunsListRequestRunTypeCorrection, nil
	case "OFF_CYCLE":
		return PayrollRunsListRequestRunTypeOffCycle, nil
	case "REGULAR":
		return PayrollRunsListRequestRunTypeRegular, nil
	case "SIGN_ON_BONUS":
		return PayrollRunsListRequestRunTypeSignOnBonus, nil
	case "TERMINATION":
		return PayrollRunsListRequestRunTypeTermination, nil
	}
	var t PayrollRunsListRequestRunType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PayrollRunsListRequestRunType) Ptr() *PayrollRunsListRequestRunType {
	return &p
}

type PayrollRunsListRequestShowEnumOrigins string

const (
	PayrollRunsListRequestShowEnumOriginsRunState        PayrollRunsListRequestShowEnumOrigins = "run_state"
	PayrollRunsListRequestShowEnumOriginsRunStateRunType PayrollRunsListRequestShowEnumOrigins = "run_state,run_type"
	PayrollRunsListRequestShowEnumOriginsRunType         PayrollRunsListRequestShowEnumOrigins = "run_type"
)

func NewPayrollRunsListRequestShowEnumOriginsFromString(s string) (PayrollRunsListRequestShowEnumOrigins, error) {
	switch s {
	case "run_state":
		return PayrollRunsListRequestShowEnumOriginsRunState, nil
	case "run_state,run_type":
		return PayrollRunsListRequestShowEnumOriginsRunStateRunType, nil
	case "run_type":
		return PayrollRunsListRequestShowEnumOriginsRunType, nil
	}
	var t PayrollRunsListRequestShowEnumOrigins
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PayrollRunsListRequestShowEnumOrigins) Ptr() *PayrollRunsListRequestShowEnumOrigins {
	return &p
}

type PayrollRunsRetrieveRequestRemoteFields string

const (
	PayrollRunsRetrieveRequestRemoteFieldsRunState        PayrollRunsRetrieveRequestRemoteFields = "run_state"
	PayrollRunsRetrieveRequestRemoteFieldsRunStateRunType PayrollRunsRetrieveRequestRemoteFields = "run_state,run_type"
	PayrollRunsRetrieveRequestRemoteFieldsRunType         PayrollRunsRetrieveRequestRemoteFields = "run_type"
)

func NewPayrollRunsRetrieveRequestRemoteFieldsFromString(s string) (PayrollRunsRetrieveRequestRemoteFields, error) {
	switch s {
	case "run_state":
		return PayrollRunsRetrieveRequestRemoteFieldsRunState, nil
	case "run_state,run_type":
		return PayrollRunsRetrieveRequestRemoteFieldsRunStateRunType, nil
	case "run_type":
		return PayrollRunsRetrieveRequestRemoteFieldsRunType, nil
	}
	var t PayrollRunsRetrieveRequestRemoteFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PayrollRunsRetrieveRequestRemoteFields) Ptr() *PayrollRunsRetrieveRequestRemoteFields {
	return &p
}

type PayrollRunsRetrieveRequestShowEnumOrigins string

const (
	PayrollRunsRetrieveRequestShowEnumOriginsRunState        PayrollRunsRetrieveRequestShowEnumOrigins = "run_state"
	PayrollRunsRetrieveRequestShowEnumOriginsRunStateRunType PayrollRunsRetrieveRequestShowEnumOrigins = "run_state,run_type"
	PayrollRunsRetrieveRequestShowEnumOriginsRunType         PayrollRunsRetrieveRequestShowEnumOrigins = "run_type"
)

func NewPayrollRunsRetrieveRequestShowEnumOriginsFromString(s string) (PayrollRunsRetrieveRequestShowEnumOrigins, error) {
	switch s {
	case "run_state":
		return PayrollRunsRetrieveRequestShowEnumOriginsRunState, nil
	case "run_state,run_type":
		return PayrollRunsRetrieveRequestShowEnumOriginsRunStateRunType, nil
	case "run_type":
		return PayrollRunsRetrieveRequestShowEnumOriginsRunType, nil
	}
	var t PayrollRunsRetrieveRequestShowEnumOrigins
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PayrollRunsRetrieveRequestShowEnumOrigins) Ptr() *PayrollRunsRetrieveRequestShowEnumOrigins {
	return &p
}
