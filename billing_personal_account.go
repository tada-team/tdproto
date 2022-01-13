package tdproto

import "time"

// PersonalAccountBilling struct of billing api
type PersonalAccountBilling struct {
	// Status of personal account
	Status PersonalAccountStatus `json:"status"`

	// Date of next debiting funds
	NextBillingDate time.Time `json:"next_billing_date"`

	// PersonalAccountBilling ID
	PersonalAccountId string `json:"personal_account_id"` // TODO: must be int64

	// Full name of owner personal account
	FullName string `json:"full_name,omitempty"`

	// Phone number of owner account
	Phone string `json:"phone,omitempty"`

	// ID User who owns this personal account
	OwnerUuid string `json:"owner_uuid"`

	// ID Tariff on this personal account
	TariffId string `json:"tariff_id"` // TODO: must be int64

	// Name Tariff on this personal account
	TariffName string `json:"tariff_name"`

	// ID Discount on personal account
	DiscountId string `json:"discount_id"` // TODO: must be int64

	// Amount of Discount on personal account
	DiscountAmount int32 `json:"discount_amount"`

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

// UpdatePersonalAccountRequest request on update personal account
type UpdatePersonalAccountRequest struct {
	FullName string `json:"full_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// UpdatePersonalAccountResponse response on update personal account
type UpdatePersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// CheckActivePersonalAccountResponse response on check active personal account
type CheckActivePersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetPersonalAccountByIDResponse response on get personal account by ID
type GetPersonalAccountByIDResponse struct {
	PersonalAccountBilling
}

// GetPersonalAccountsListResponse response on get list of personal accounts
type GetPersonalAccountsListResponse struct {
	PersonalAccounts []PersonalAccountBilling `json:"personal_accounts,omitempty"`
}

// ActivatePersonalAccountResponse response on activate suspended personal account
type ActivatePersonalAccountResponse struct {
	Success bool `json:"success"`
}

// BlockPersonalAccountResponse response on block unblocked personal account
type BlockPersonalAccountResponse struct {
	Success bool `json:"success"`
}

// UnblockPersonalAccountResponse response on unblock blocked personal account
type UnblockPersonalAccountResponse struct {
	Success bool `json:"success"`
}

// SuspendPersonalAccountResponse response on suspend active personal account
type SuspendPersonalAccountResponse struct {
	Success bool `json:"success"`
}
