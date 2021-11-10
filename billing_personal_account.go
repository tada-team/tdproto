package tdproto

import "time"

// PersonalAccountBilling struct of billing api
type PersonalAccountBilling struct {

	// PersonalAccountBilling ID
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`

	// Full name of owner personal account
	FullName string `json:"full_name,omitempty"`

	// Phone number of owner account
	Phone string `json:"phone,omitempty"`

	// ID User who owns this personal account
	OwnerUuid string `json:"owner_uuid"`

	// ID Tariff on this personal account
	TariffId int64 `json:"tariff_id"`

	// Name Tariff on this personal account
	TariffName string `json:"tariff_name"`

	// ID Discount on personal account
	DiscountId int64 `json:"discount_id"`

	// Amount of Discount on personal account
	DiscountAmount int32 `json:"discount_amount"`

	// Status of personal account
	Status string `json:"status"`

	// Date of next debiting funds
	NextBillingDate time.Time `json:"next_billing_date"`

	// Count of teams on personal account
	TeamCount int32 `json:"team_count"`

	// Count of workplaces on personal account
	WorkplaceCount int32 `json:"workplace_count"`

	// Count of user on personal account
	UsersCount int32 `json:"users_count"`

	// Count of free workplaces on personal account
	FreeWorkplaces int32 `json:"free_workplaces"`

	// Count of paid workplaces on personal account
	PaidWorkplaces int32 `json:"paid_workplaces"`
}

// CreatePersonalAccountRequest request on create personal account
type CreatePersonalAccountRequest struct {
	OwnerUuid string `json:"owner_uuid"`
	FullName  string `json:"full_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	TeamUuid  string `json:"team_uuid"`
}

// CreatePersonalAccountResponse response on create personal account
type CreatePersonalAccountResponse struct {
	PersonalAccountBilling
}

// GetPersonalAccountByIDRequest request on get personal account by ID
type GetPersonalAccountByIDRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
}

// GetPersonalAccountByIDResponse response on get personal account by ID
type GetPersonalAccountByIDResponse struct {
	PersonalAccountBilling
}

// GetPersonalAccountsListRequest request on get list of personal accounts
type GetPersonalAccountsListRequest struct {
	PersonalAccountId int64    `json:"personal_account_id,omitempty"`
	Options           *Options `json:"options,omitempty"`
}

// Options struct for pagination
type Options struct {
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

// GetPersonalAccountsListResponse response on get list of personal accounts
type GetPersonalAccountsListResponse struct {
	PersonalAccounts []PersonalAccountBilling `json:"personal_accounts,omitempty"`
}

// ActivatePersonalAccountRequest request on activate suspended personal account
type ActivatePersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
}

// ActivatePersonalAccountResponse response on activate suspended personal account
type ActivatePersonalAccountResponse struct {
	Success bool `json:"success"`
}

// BlockPersonalAccountRequest request on block unblocked personal account
type BlockPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
}

// BlockPersonalAccountResponse response on block unblocked personal account
type BlockPersonalAccountResponse struct {
	Success bool `json:"success"`
}

// UnblockPersonalAccountRequest request on unblock blocked personal account
type UnblockPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
}

// UnblockPersonalAccountResponse response on unblock blocked personal account
type UnblockPersonalAccountResponse struct {
	Success bool `json:"success"`
}

// SuspendPersonalAccountRequest request on suspend active personal account
type SuspendPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
}

// SuspendPersonalAccountResponse response on suspend active personal account
type SuspendPersonalAccountResponse struct {
	Success bool `json:"success"`
}
