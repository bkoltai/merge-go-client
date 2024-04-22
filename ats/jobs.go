// This file was auto-generated by Fern from our API Definition.

package ats

import (
	fmt "fmt"
	time "time"
)

type JobsListRequest struct {
	// If provided, will only return jobs with this code.
	Code *string `json:"-"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JobsListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// If provided, will only return jobs for this office; multiple offices can be separated by commas.
	Offices *string `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *string `json:"-"`
	// If provided, will only return jobs with this status. Options: ('OPEN', 'CLOSED', 'DRAFT', 'ARCHIVED', 'PENDING')
	//
	// - `OPEN` - OPEN
	// - `CLOSED` - CLOSED
	// - `DRAFT` - DRAFT
	// - `ARCHIVED` - ARCHIVED
	// - `PENDING` - PENDING
	Status *JobsListRequestStatus `json:"-"`
}

type JobsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JobsRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *string `json:"-"`
}

type JobsScreeningQuestionsListRequest struct {
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JobsScreeningQuestionsListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
}

type JobsListRequestExpand string

const (
	JobsListRequestExpandDepartments                                           JobsListRequestExpand = "departments"
	JobsListRequestExpandDepartmentsHiringManagers                             JobsListRequestExpand = "departments,hiring_managers"
	JobsListRequestExpandDepartmentsHiringManagersJobPostings                  JobsListRequestExpand = "departments,hiring_managers,job_postings"
	JobsListRequestExpandDepartmentsHiringManagersJobPostingsRecruiters        JobsListRequestExpand = "departments,hiring_managers,job_postings,recruiters"
	JobsListRequestExpandDepartmentsHiringManagersRecruiters                   JobsListRequestExpand = "departments,hiring_managers,recruiters"
	JobsListRequestExpandDepartmentsJobPostings                                JobsListRequestExpand = "departments,job_postings"
	JobsListRequestExpandDepartmentsJobPostingsRecruiters                      JobsListRequestExpand = "departments,job_postings,recruiters"
	JobsListRequestExpandDepartmentsOffices                                    JobsListRequestExpand = "departments,offices"
	JobsListRequestExpandDepartmentsOfficesHiringManagers                      JobsListRequestExpand = "departments,offices,hiring_managers"
	JobsListRequestExpandDepartmentsOfficesHiringManagersJobPostings           JobsListRequestExpand = "departments,offices,hiring_managers,job_postings"
	JobsListRequestExpandDepartmentsOfficesHiringManagersJobPostingsRecruiters JobsListRequestExpand = "departments,offices,hiring_managers,job_postings,recruiters"
	JobsListRequestExpandDepartmentsOfficesHiringManagersRecruiters            JobsListRequestExpand = "departments,offices,hiring_managers,recruiters"
	JobsListRequestExpandDepartmentsOfficesJobPostings                         JobsListRequestExpand = "departments,offices,job_postings"
	JobsListRequestExpandDepartmentsOfficesJobPostingsRecruiters               JobsListRequestExpand = "departments,offices,job_postings,recruiters"
	JobsListRequestExpandDepartmentsOfficesRecruiters                          JobsListRequestExpand = "departments,offices,recruiters"
	JobsListRequestExpandDepartmentsRecruiters                                 JobsListRequestExpand = "departments,recruiters"
	JobsListRequestExpandHiringManagers                                        JobsListRequestExpand = "hiring_managers"
	JobsListRequestExpandHiringManagersJobPostings                             JobsListRequestExpand = "hiring_managers,job_postings"
	JobsListRequestExpandHiringManagersJobPostingsRecruiters                   JobsListRequestExpand = "hiring_managers,job_postings,recruiters"
	JobsListRequestExpandHiringManagersRecruiters                              JobsListRequestExpand = "hiring_managers,recruiters"
	JobsListRequestExpandJobPostings                                           JobsListRequestExpand = "job_postings"
	JobsListRequestExpandJobPostingsRecruiters                                 JobsListRequestExpand = "job_postings,recruiters"
	JobsListRequestExpandOffices                                               JobsListRequestExpand = "offices"
	JobsListRequestExpandOfficesHiringManagers                                 JobsListRequestExpand = "offices,hiring_managers"
	JobsListRequestExpandOfficesHiringManagersJobPostings                      JobsListRequestExpand = "offices,hiring_managers,job_postings"
	JobsListRequestExpandOfficesHiringManagersJobPostingsRecruiters            JobsListRequestExpand = "offices,hiring_managers,job_postings,recruiters"
	JobsListRequestExpandOfficesHiringManagersRecruiters                       JobsListRequestExpand = "offices,hiring_managers,recruiters"
	JobsListRequestExpandOfficesJobPostings                                    JobsListRequestExpand = "offices,job_postings"
	JobsListRequestExpandOfficesJobPostingsRecruiters                          JobsListRequestExpand = "offices,job_postings,recruiters"
	JobsListRequestExpandOfficesRecruiters                                     JobsListRequestExpand = "offices,recruiters"
	JobsListRequestExpandRecruiters                                            JobsListRequestExpand = "recruiters"
)

func NewJobsListRequestExpandFromString(s string) (JobsListRequestExpand, error) {
	switch s {
	case "departments":
		return JobsListRequestExpandDepartments, nil
	case "departments,hiring_managers":
		return JobsListRequestExpandDepartmentsHiringManagers, nil
	case "departments,hiring_managers,job_postings":
		return JobsListRequestExpandDepartmentsHiringManagersJobPostings, nil
	case "departments,hiring_managers,job_postings,recruiters":
		return JobsListRequestExpandDepartmentsHiringManagersJobPostingsRecruiters, nil
	case "departments,hiring_managers,recruiters":
		return JobsListRequestExpandDepartmentsHiringManagersRecruiters, nil
	case "departments,job_postings":
		return JobsListRequestExpandDepartmentsJobPostings, nil
	case "departments,job_postings,recruiters":
		return JobsListRequestExpandDepartmentsJobPostingsRecruiters, nil
	case "departments,offices":
		return JobsListRequestExpandDepartmentsOffices, nil
	case "departments,offices,hiring_managers":
		return JobsListRequestExpandDepartmentsOfficesHiringManagers, nil
	case "departments,offices,hiring_managers,job_postings":
		return JobsListRequestExpandDepartmentsOfficesHiringManagersJobPostings, nil
	case "departments,offices,hiring_managers,job_postings,recruiters":
		return JobsListRequestExpandDepartmentsOfficesHiringManagersJobPostingsRecruiters, nil
	case "departments,offices,hiring_managers,recruiters":
		return JobsListRequestExpandDepartmentsOfficesHiringManagersRecruiters, nil
	case "departments,offices,job_postings":
		return JobsListRequestExpandDepartmentsOfficesJobPostings, nil
	case "departments,offices,job_postings,recruiters":
		return JobsListRequestExpandDepartmentsOfficesJobPostingsRecruiters, nil
	case "departments,offices,recruiters":
		return JobsListRequestExpandDepartmentsOfficesRecruiters, nil
	case "departments,recruiters":
		return JobsListRequestExpandDepartmentsRecruiters, nil
	case "hiring_managers":
		return JobsListRequestExpandHiringManagers, nil
	case "hiring_managers,job_postings":
		return JobsListRequestExpandHiringManagersJobPostings, nil
	case "hiring_managers,job_postings,recruiters":
		return JobsListRequestExpandHiringManagersJobPostingsRecruiters, nil
	case "hiring_managers,recruiters":
		return JobsListRequestExpandHiringManagersRecruiters, nil
	case "job_postings":
		return JobsListRequestExpandJobPostings, nil
	case "job_postings,recruiters":
		return JobsListRequestExpandJobPostingsRecruiters, nil
	case "offices":
		return JobsListRequestExpandOffices, nil
	case "offices,hiring_managers":
		return JobsListRequestExpandOfficesHiringManagers, nil
	case "offices,hiring_managers,job_postings":
		return JobsListRequestExpandOfficesHiringManagersJobPostings, nil
	case "offices,hiring_managers,job_postings,recruiters":
		return JobsListRequestExpandOfficesHiringManagersJobPostingsRecruiters, nil
	case "offices,hiring_managers,recruiters":
		return JobsListRequestExpandOfficesHiringManagersRecruiters, nil
	case "offices,job_postings":
		return JobsListRequestExpandOfficesJobPostings, nil
	case "offices,job_postings,recruiters":
		return JobsListRequestExpandOfficesJobPostingsRecruiters, nil
	case "offices,recruiters":
		return JobsListRequestExpandOfficesRecruiters, nil
	case "recruiters":
		return JobsListRequestExpandRecruiters, nil
	}
	var t JobsListRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (j JobsListRequestExpand) Ptr() *JobsListRequestExpand {
	return &j
}

type JobsListRequestStatus string

const (
	JobsListRequestStatusArchived JobsListRequestStatus = "ARCHIVED"
	JobsListRequestStatusClosed   JobsListRequestStatus = "CLOSED"
	JobsListRequestStatusDraft    JobsListRequestStatus = "DRAFT"
	JobsListRequestStatusOpen     JobsListRequestStatus = "OPEN"
	JobsListRequestStatusPending  JobsListRequestStatus = "PENDING"
)

func NewJobsListRequestStatusFromString(s string) (JobsListRequestStatus, error) {
	switch s {
	case "ARCHIVED":
		return JobsListRequestStatusArchived, nil
	case "CLOSED":
		return JobsListRequestStatusClosed, nil
	case "DRAFT":
		return JobsListRequestStatusDraft, nil
	case "OPEN":
		return JobsListRequestStatusOpen, nil
	case "PENDING":
		return JobsListRequestStatusPending, nil
	}
	var t JobsListRequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (j JobsListRequestStatus) Ptr() *JobsListRequestStatus {
	return &j
}

type JobsRetrieveRequestExpand string

const (
	JobsRetrieveRequestExpandDepartments                                           JobsRetrieveRequestExpand = "departments"
	JobsRetrieveRequestExpandDepartmentsHiringManagers                             JobsRetrieveRequestExpand = "departments,hiring_managers"
	JobsRetrieveRequestExpandDepartmentsHiringManagersJobPostings                  JobsRetrieveRequestExpand = "departments,hiring_managers,job_postings"
	JobsRetrieveRequestExpandDepartmentsHiringManagersJobPostingsRecruiters        JobsRetrieveRequestExpand = "departments,hiring_managers,job_postings,recruiters"
	JobsRetrieveRequestExpandDepartmentsHiringManagersRecruiters                   JobsRetrieveRequestExpand = "departments,hiring_managers,recruiters"
	JobsRetrieveRequestExpandDepartmentsJobPostings                                JobsRetrieveRequestExpand = "departments,job_postings"
	JobsRetrieveRequestExpandDepartmentsJobPostingsRecruiters                      JobsRetrieveRequestExpand = "departments,job_postings,recruiters"
	JobsRetrieveRequestExpandDepartmentsOffices                                    JobsRetrieveRequestExpand = "departments,offices"
	JobsRetrieveRequestExpandDepartmentsOfficesHiringManagers                      JobsRetrieveRequestExpand = "departments,offices,hiring_managers"
	JobsRetrieveRequestExpandDepartmentsOfficesHiringManagersJobPostings           JobsRetrieveRequestExpand = "departments,offices,hiring_managers,job_postings"
	JobsRetrieveRequestExpandDepartmentsOfficesHiringManagersJobPostingsRecruiters JobsRetrieveRequestExpand = "departments,offices,hiring_managers,job_postings,recruiters"
	JobsRetrieveRequestExpandDepartmentsOfficesHiringManagersRecruiters            JobsRetrieveRequestExpand = "departments,offices,hiring_managers,recruiters"
	JobsRetrieveRequestExpandDepartmentsOfficesJobPostings                         JobsRetrieveRequestExpand = "departments,offices,job_postings"
	JobsRetrieveRequestExpandDepartmentsOfficesJobPostingsRecruiters               JobsRetrieveRequestExpand = "departments,offices,job_postings,recruiters"
	JobsRetrieveRequestExpandDepartmentsOfficesRecruiters                          JobsRetrieveRequestExpand = "departments,offices,recruiters"
	JobsRetrieveRequestExpandDepartmentsRecruiters                                 JobsRetrieveRequestExpand = "departments,recruiters"
	JobsRetrieveRequestExpandHiringManagers                                        JobsRetrieveRequestExpand = "hiring_managers"
	JobsRetrieveRequestExpandHiringManagersJobPostings                             JobsRetrieveRequestExpand = "hiring_managers,job_postings"
	JobsRetrieveRequestExpandHiringManagersJobPostingsRecruiters                   JobsRetrieveRequestExpand = "hiring_managers,job_postings,recruiters"
	JobsRetrieveRequestExpandHiringManagersRecruiters                              JobsRetrieveRequestExpand = "hiring_managers,recruiters"
	JobsRetrieveRequestExpandJobPostings                                           JobsRetrieveRequestExpand = "job_postings"
	JobsRetrieveRequestExpandJobPostingsRecruiters                                 JobsRetrieveRequestExpand = "job_postings,recruiters"
	JobsRetrieveRequestExpandOffices                                               JobsRetrieveRequestExpand = "offices"
	JobsRetrieveRequestExpandOfficesHiringManagers                                 JobsRetrieveRequestExpand = "offices,hiring_managers"
	JobsRetrieveRequestExpandOfficesHiringManagersJobPostings                      JobsRetrieveRequestExpand = "offices,hiring_managers,job_postings"
	JobsRetrieveRequestExpandOfficesHiringManagersJobPostingsRecruiters            JobsRetrieveRequestExpand = "offices,hiring_managers,job_postings,recruiters"
	JobsRetrieveRequestExpandOfficesHiringManagersRecruiters                       JobsRetrieveRequestExpand = "offices,hiring_managers,recruiters"
	JobsRetrieveRequestExpandOfficesJobPostings                                    JobsRetrieveRequestExpand = "offices,job_postings"
	JobsRetrieveRequestExpandOfficesJobPostingsRecruiters                          JobsRetrieveRequestExpand = "offices,job_postings,recruiters"
	JobsRetrieveRequestExpandOfficesRecruiters                                     JobsRetrieveRequestExpand = "offices,recruiters"
	JobsRetrieveRequestExpandRecruiters                                            JobsRetrieveRequestExpand = "recruiters"
)

func NewJobsRetrieveRequestExpandFromString(s string) (JobsRetrieveRequestExpand, error) {
	switch s {
	case "departments":
		return JobsRetrieveRequestExpandDepartments, nil
	case "departments,hiring_managers":
		return JobsRetrieveRequestExpandDepartmentsHiringManagers, nil
	case "departments,hiring_managers,job_postings":
		return JobsRetrieveRequestExpandDepartmentsHiringManagersJobPostings, nil
	case "departments,hiring_managers,job_postings,recruiters":
		return JobsRetrieveRequestExpandDepartmentsHiringManagersJobPostingsRecruiters, nil
	case "departments,hiring_managers,recruiters":
		return JobsRetrieveRequestExpandDepartmentsHiringManagersRecruiters, nil
	case "departments,job_postings":
		return JobsRetrieveRequestExpandDepartmentsJobPostings, nil
	case "departments,job_postings,recruiters":
		return JobsRetrieveRequestExpandDepartmentsJobPostingsRecruiters, nil
	case "departments,offices":
		return JobsRetrieveRequestExpandDepartmentsOffices, nil
	case "departments,offices,hiring_managers":
		return JobsRetrieveRequestExpandDepartmentsOfficesHiringManagers, nil
	case "departments,offices,hiring_managers,job_postings":
		return JobsRetrieveRequestExpandDepartmentsOfficesHiringManagersJobPostings, nil
	case "departments,offices,hiring_managers,job_postings,recruiters":
		return JobsRetrieveRequestExpandDepartmentsOfficesHiringManagersJobPostingsRecruiters, nil
	case "departments,offices,hiring_managers,recruiters":
		return JobsRetrieveRequestExpandDepartmentsOfficesHiringManagersRecruiters, nil
	case "departments,offices,job_postings":
		return JobsRetrieveRequestExpandDepartmentsOfficesJobPostings, nil
	case "departments,offices,job_postings,recruiters":
		return JobsRetrieveRequestExpandDepartmentsOfficesJobPostingsRecruiters, nil
	case "departments,offices,recruiters":
		return JobsRetrieveRequestExpandDepartmentsOfficesRecruiters, nil
	case "departments,recruiters":
		return JobsRetrieveRequestExpandDepartmentsRecruiters, nil
	case "hiring_managers":
		return JobsRetrieveRequestExpandHiringManagers, nil
	case "hiring_managers,job_postings":
		return JobsRetrieveRequestExpandHiringManagersJobPostings, nil
	case "hiring_managers,job_postings,recruiters":
		return JobsRetrieveRequestExpandHiringManagersJobPostingsRecruiters, nil
	case "hiring_managers,recruiters":
		return JobsRetrieveRequestExpandHiringManagersRecruiters, nil
	case "job_postings":
		return JobsRetrieveRequestExpandJobPostings, nil
	case "job_postings,recruiters":
		return JobsRetrieveRequestExpandJobPostingsRecruiters, nil
	case "offices":
		return JobsRetrieveRequestExpandOffices, nil
	case "offices,hiring_managers":
		return JobsRetrieveRequestExpandOfficesHiringManagers, nil
	case "offices,hiring_managers,job_postings":
		return JobsRetrieveRequestExpandOfficesHiringManagersJobPostings, nil
	case "offices,hiring_managers,job_postings,recruiters":
		return JobsRetrieveRequestExpandOfficesHiringManagersJobPostingsRecruiters, nil
	case "offices,hiring_managers,recruiters":
		return JobsRetrieveRequestExpandOfficesHiringManagersRecruiters, nil
	case "offices,job_postings":
		return JobsRetrieveRequestExpandOfficesJobPostings, nil
	case "offices,job_postings,recruiters":
		return JobsRetrieveRequestExpandOfficesJobPostingsRecruiters, nil
	case "offices,recruiters":
		return JobsRetrieveRequestExpandOfficesRecruiters, nil
	case "recruiters":
		return JobsRetrieveRequestExpandRecruiters, nil
	}
	var t JobsRetrieveRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (j JobsRetrieveRequestExpand) Ptr() *JobsRetrieveRequestExpand {
	return &j
}

type JobsScreeningQuestionsListRequestExpand string

const (
	JobsScreeningQuestionsListRequestExpandJob        JobsScreeningQuestionsListRequestExpand = "job"
	JobsScreeningQuestionsListRequestExpandOptions    JobsScreeningQuestionsListRequestExpand = "options"
	JobsScreeningQuestionsListRequestExpandOptionsJob JobsScreeningQuestionsListRequestExpand = "options,job"
)

func NewJobsScreeningQuestionsListRequestExpandFromString(s string) (JobsScreeningQuestionsListRequestExpand, error) {
	switch s {
	case "job":
		return JobsScreeningQuestionsListRequestExpandJob, nil
	case "options":
		return JobsScreeningQuestionsListRequestExpandOptions, nil
	case "options,job":
		return JobsScreeningQuestionsListRequestExpandOptionsJob, nil
	}
	var t JobsScreeningQuestionsListRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (j JobsScreeningQuestionsListRequestExpand) Ptr() *JobsScreeningQuestionsListRequestExpand {
	return &j
}
