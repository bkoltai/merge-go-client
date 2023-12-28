// This file was auto-generated by Fern from our API Definition.

package accounting

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
	time "time"
)

type ItemsListRequest struct {
	// If provided, will only return items for this company.
	CompanyId *string `json:"-"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *ItemsListRequestExpand `json:"-"`
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
	RemoteFields *string `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *string `json:"-"`
}

type ItemsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *ItemsRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *string `json:"-"`
}

type ItemsListRequestExpand uint

const (
	ItemsListRequestExpandCompany ItemsListRequestExpand = iota + 1
	ItemsListRequestExpandPurchaseAccount
	ItemsListRequestExpandPurchaseAccountCompany
	ItemsListRequestExpandPurchaseAccountSalesAccount
	ItemsListRequestExpandPurchaseAccountSalesAccountCompany
	ItemsListRequestExpandSalesAccount
	ItemsListRequestExpandSalesAccountCompany
)

func (i ItemsListRequestExpand) String() string {
	switch i {
	default:
		return strconv.Itoa(int(i))
	case ItemsListRequestExpandCompany:
		return "company"
	case ItemsListRequestExpandPurchaseAccount:
		return "purchase_account"
	case ItemsListRequestExpandPurchaseAccountCompany:
		return "purchase_account,company"
	case ItemsListRequestExpandPurchaseAccountSalesAccount:
		return "purchase_account,sales_account"
	case ItemsListRequestExpandPurchaseAccountSalesAccountCompany:
		return "purchase_account,sales_account,company"
	case ItemsListRequestExpandSalesAccount:
		return "sales_account"
	case ItemsListRequestExpandSalesAccountCompany:
		return "sales_account,company"
	}
}

func (i ItemsListRequestExpand) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", i.String())), nil
}

func (i *ItemsListRequestExpand) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "company":
		value := ItemsListRequestExpandCompany
		*i = value
	case "purchase_account":
		value := ItemsListRequestExpandPurchaseAccount
		*i = value
	case "purchase_account,company":
		value := ItemsListRequestExpandPurchaseAccountCompany
		*i = value
	case "purchase_account,sales_account":
		value := ItemsListRequestExpandPurchaseAccountSalesAccount
		*i = value
	case "purchase_account,sales_account,company":
		value := ItemsListRequestExpandPurchaseAccountSalesAccountCompany
		*i = value
	case "sales_account":
		value := ItemsListRequestExpandSalesAccount
		*i = value
	case "sales_account,company":
		value := ItemsListRequestExpandSalesAccountCompany
		*i = value
	}
	return nil
}

type ItemsRetrieveRequestExpand uint

const (
	ItemsRetrieveRequestExpandCompany ItemsRetrieveRequestExpand = iota + 1
	ItemsRetrieveRequestExpandPurchaseAccount
	ItemsRetrieveRequestExpandPurchaseAccountCompany
	ItemsRetrieveRequestExpandPurchaseAccountSalesAccount
	ItemsRetrieveRequestExpandPurchaseAccountSalesAccountCompany
	ItemsRetrieveRequestExpandSalesAccount
	ItemsRetrieveRequestExpandSalesAccountCompany
)

func (i ItemsRetrieveRequestExpand) String() string {
	switch i {
	default:
		return strconv.Itoa(int(i))
	case ItemsRetrieveRequestExpandCompany:
		return "company"
	case ItemsRetrieveRequestExpandPurchaseAccount:
		return "purchase_account"
	case ItemsRetrieveRequestExpandPurchaseAccountCompany:
		return "purchase_account,company"
	case ItemsRetrieveRequestExpandPurchaseAccountSalesAccount:
		return "purchase_account,sales_account"
	case ItemsRetrieveRequestExpandPurchaseAccountSalesAccountCompany:
		return "purchase_account,sales_account,company"
	case ItemsRetrieveRequestExpandSalesAccount:
		return "sales_account"
	case ItemsRetrieveRequestExpandSalesAccountCompany:
		return "sales_account,company"
	}
}

func (i ItemsRetrieveRequestExpand) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", i.String())), nil
}

func (i *ItemsRetrieveRequestExpand) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "company":
		value := ItemsRetrieveRequestExpandCompany
		*i = value
	case "purchase_account":
		value := ItemsRetrieveRequestExpandPurchaseAccount
		*i = value
	case "purchase_account,company":
		value := ItemsRetrieveRequestExpandPurchaseAccountCompany
		*i = value
	case "purchase_account,sales_account":
		value := ItemsRetrieveRequestExpandPurchaseAccountSalesAccount
		*i = value
	case "purchase_account,sales_account,company":
		value := ItemsRetrieveRequestExpandPurchaseAccountSalesAccountCompany
		*i = value
	case "sales_account":
		value := ItemsRetrieveRequestExpandSalesAccount
		*i = value
	case "sales_account,company":
		value := ItemsRetrieveRequestExpandSalesAccountCompany
		*i = value
	}
	return nil
}
