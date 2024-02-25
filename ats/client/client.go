// This file was auto-generated by Fern from our API Definition.

package client

import (
	accountdetails "github.com/merge-api/merge-go-client/ats/accountdetails"
	accounttoken "github.com/merge-api/merge-go-client/ats/accounttoken"
	activities "github.com/merge-api/merge-go-client/ats/activities"
	applications "github.com/merge-api/merge-go-client/ats/applications"
	asyncpassthrough "github.com/merge-api/merge-go-client/ats/asyncpassthrough"
	attachments "github.com/merge-api/merge-go-client/ats/attachments"
	audittrail "github.com/merge-api/merge-go-client/ats/audittrail"
	availableactions "github.com/merge-api/merge-go-client/ats/availableactions"
	candidates "github.com/merge-api/merge-go-client/ats/candidates"
	deleteaccount "github.com/merge-api/merge-go-client/ats/deleteaccount"
	departments "github.com/merge-api/merge-go-client/ats/departments"
	eeocs "github.com/merge-api/merge-go-client/ats/eeocs"
	fieldmapping "github.com/merge-api/merge-go-client/ats/fieldmapping"
	forceresync "github.com/merge-api/merge-go-client/ats/forceresync"
	generatekey "github.com/merge-api/merge-go-client/ats/generatekey"
	interviews "github.com/merge-api/merge-go-client/ats/interviews"
	issues "github.com/merge-api/merge-go-client/ats/issues"
	jobinterviewstages "github.com/merge-api/merge-go-client/ats/jobinterviewstages"
	jobs "github.com/merge-api/merge-go-client/ats/jobs"
	linkedaccounts "github.com/merge-api/merge-go-client/ats/linkedaccounts"
	linktoken "github.com/merge-api/merge-go-client/ats/linktoken"
	offers "github.com/merge-api/merge-go-client/ats/offers"
	offices "github.com/merge-api/merge-go-client/ats/offices"
	passthrough "github.com/merge-api/merge-go-client/ats/passthrough"
	regeneratekey "github.com/merge-api/merge-go-client/ats/regeneratekey"
	rejectreasons "github.com/merge-api/merge-go-client/ats/rejectreasons"
	scopes "github.com/merge-api/merge-go-client/ats/scopes"
	scorecards "github.com/merge-api/merge-go-client/ats/scorecards"
	selectivesync "github.com/merge-api/merge-go-client/ats/selectivesync"
	syncstatus "github.com/merge-api/merge-go-client/ats/syncstatus"
	tags "github.com/merge-api/merge-go-client/ats/tags"
	users "github.com/merge-api/merge-go-client/ats/users"
	webhookreceivers "github.com/merge-api/merge-go-client/ats/webhookreceivers"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	AccountDetails     *accountdetails.Client
	AccountToken       *accounttoken.Client
	Activities         *activities.Client
	Applications       *applications.Client
	AsyncPassthrough   *asyncpassthrough.Client
	Attachments        *attachments.Client
	AuditTrail         *audittrail.Client
	AvailableActions   *availableactions.Client
	Candidates         *candidates.Client
	Scopes             *scopes.Client
	DeleteAccount      *deleteaccount.Client
	Departments        *departments.Client
	Eeocs              *eeocs.Client
	FieldMapping       *fieldmapping.Client
	GenerateKey        *generatekey.Client
	Interviews         *interviews.Client
	Issues             *issues.Client
	JobInterviewStages *jobinterviewstages.Client
	Jobs               *jobs.Client
	LinkToken          *linktoken.Client
	LinkedAccounts     *linkedaccounts.Client
	Offers             *offers.Client
	Offices            *offices.Client
	Passthrough        *passthrough.Client
	RegenerateKey      *regeneratekey.Client
	RejectReasons      *rejectreasons.Client
	Scorecards         *scorecards.Client
	SelectiveSync      *selectivesync.Client
	SyncStatus         *syncstatus.Client
	ForceResync        *forceresync.Client
	Tags               *tags.Client
	Users              *users.Client
	WebhookReceivers   *webhookreceivers.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:            options.BaseURL,
		caller:             core.NewCaller(options.HTTPClient),
		header:             options.ToHeader(),
		AccountDetails:     accountdetails.NewClient(opts...),
		AccountToken:       accounttoken.NewClient(opts...),
		Activities:         activities.NewClient(opts...),
		Applications:       applications.NewClient(opts...),
		AsyncPassthrough:   asyncpassthrough.NewClient(opts...),
		Attachments:        attachments.NewClient(opts...),
		AuditTrail:         audittrail.NewClient(opts...),
		AvailableActions:   availableactions.NewClient(opts...),
		Candidates:         candidates.NewClient(opts...),
		Scopes:             scopes.NewClient(opts...),
		DeleteAccount:      deleteaccount.NewClient(opts...),
		Departments:        departments.NewClient(opts...),
		Eeocs:              eeocs.NewClient(opts...),
		FieldMapping:       fieldmapping.NewClient(opts...),
		GenerateKey:        generatekey.NewClient(opts...),
		Interviews:         interviews.NewClient(opts...),
		Issues:             issues.NewClient(opts...),
		JobInterviewStages: jobinterviewstages.NewClient(opts...),
		Jobs:               jobs.NewClient(opts...),
		LinkToken:          linktoken.NewClient(opts...),
		LinkedAccounts:     linkedaccounts.NewClient(opts...),
		Offers:             offers.NewClient(opts...),
		Offices:            offices.NewClient(opts...),
		Passthrough:        passthrough.NewClient(opts...),
		RegenerateKey:      regeneratekey.NewClient(opts...),
		RejectReasons:      rejectreasons.NewClient(opts...),
		Scorecards:         scorecards.NewClient(opts...),
		SelectiveSync:      selectivesync.NewClient(opts...),
		SyncStatus:         syncstatus.NewClient(opts...),
		ForceResync:        forceresync.NewClient(opts...),
		Tags:               tags.NewClient(opts...),
		Users:              users.NewClient(opts...),
		WebhookReceivers:   webhookreceivers.NewClient(opts...),
	}
}
