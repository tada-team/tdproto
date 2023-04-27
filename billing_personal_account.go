package tdproto

import "time"

// PersonalAccountBilling struct of billing api
type PersonalAccountBilling struct {

	// PersonalAccountBilling ID
	PersonalAccountId string `json:"personal_account_id"`

	// ID User who owns this personal account
	OwnerID string `json:"owner_id"`

	// UUID of User who owns this personal account
	OwnerUuid string `json:"owner_uuid"`

	// Count of teams on personal account
	TeamsCount uint32 `json:"teams_count"`

	// Count of workplaces on personal account
	WorkplaceCount uint32 `json:"workplace_count"`

	// Count of empty workplaces on personal account
	EmptyWorkplaceCount uint32 `json:"empty_workplace_count"`

	// Count of occupied workplaces on personal account
	OccupiedWorkplaceCount uint32 `json:"occupied_workplace_count"`

	// Count of free workplaces on personal account
	FreeWorkplaceCount uint32 `json:"free_workplace_count"`

	// Count of paid workplaces on personal account
	PaidWorkplaceCount uint32 `json:"paid_workplace_count"`

	// Is the account blocked
	IsBlocked bool `json:"is_blocked"`

	// Is the account suspended
	IsSuspended bool `json:"is_suspended"`

	// Date of next debiting funds
	NextBillingDate *time.Time `json:"next_billing_date,omitempty"`

	// Account blocking date
	BlockDate *time.Time `json:"block_date,omitempty"`

	// Account suspending date
	SuspendDate *time.Time `json:"suspend_date,omitempty"`

	// Status of personal account
	Status PersonalAccountStatus `json:"status"`

	// Tariff on this personal account
	Tariff TariffBilling `json:"tariff"`

	// Owner of this personal account
	Owner Contact `json:"owner,omitempty"`
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

type GetTeamsOnPersonalAccountResponse struct {
	Teams []GetTeamOnPersonalAccountResponse `json:"teams"`
}

type GetTeamOnPersonalAccountResponse struct {
	PersonalAccountId string    `json:"personal_account_id"`
	TeamId            string    `json:"team_id"`
	TeamUuid          string    `json:"team_uuid"`
	OpenDate          time.Time `json:"open_date"`
	CloseDate         time.Time `json:"close_date,omitempty"`
}
