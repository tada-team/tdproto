package tdproto

import "time"

// TariffBilling struct of billing api
type TariffBilling struct {

	// Tariff id
	TariffId int64 `json:"tariff_id"`

	// Name of tariff on russian
	TariffName string `json:"tariff_name,omitempty"`

	// Count of free workspaces
	FreeWorkplace int32 `json:"free_workplace,omitempty"`

	// Disk space limit per user
	DiskSpaceQuota string `json:"disk_space_quota,omitempty"`

	// Flag of availability of free seats when exceeding FreeWorkplace
	IsBillingFree bool `json:"billing_free,omitempty"`

	// Flag of accounting without looking at the number of days before the billing period
	IsBillingFullTime bool `json:"billing_full_time,omitempty"`

	// Number of paid days
	PeriodDays int32 `json:"period_days,omitempty"`

	// Cost of one workplace
	CostWorkplace string `json:"cost_workplace,omitempty"`

	// Currency of tariff in ISO
	Currency string `json:"currency,omitempty"`

	// Flag for accounting for unspent days when switching to a new tariff
	IsRecalcChangeTariff bool `json:"recalc_change_tariff,omitempty"`

	// Maximum count of users in voice conference
	MaxVoiceUser int32 `json:"max_voice_user,omitempty"`

	// Maximum count of users in video conference
	MaxVideoUser int32 `json:"max_video_user,omitempty"`

	// Default tariff flag that is set when registering an account
	IsDefaultTariff bool `json:"default_tariff,omitempty"`

	// Date of opening tariff
	OpenDate *time.Time `json:"open_date,omitempty"`

	// Date of closing tariff
	CloseDate *time.Time `json:"close_date,omitempty"`
}

type TariffBillingWithStatus struct {
	TariffBilling

	// Status
	Status string `json:"status,omitempty"`
}

type CreateTariffRequest struct {
	TariffBilling
}

type CreateTariffResponse struct {
	TariffBilling
}

type GetTariffByIdRequest struct {
	Id int64 `json:"id"`
}

type GetTariffByIdResponse struct {
	TariffBilling
}

type GetTariffsListResponse struct {
	Tariffs []TariffBillingWithStatus `json:"tariffs"`
}

type GetActiveTariffsListResponse struct {
	Tariffs []TariffBillingWithStatus `json:"tariffs"`
}

type CloseTariffRequest struct {
	TariffId  int64     `json:"tariff_id"`
	CloseDate time.Time `json:"close_date,omitempty"`
}

type CloseTariffResponse struct {
	Success bool `json:"success"`
}

type SetTariffAsDefaultRequest struct {
	TariffId int64 `json:"tariff_id"`
}

type SetTariffAsDefaultResponse struct {
	Success bool `json:"success"`
}

type PersonalAccount struct {
	PersonalAccountId int64     `json:"personal_account_id"`
	OwnerUuid         string    `json:"owner_uuid"`
	TariffId          int64     `json:"tariff_id"`
	TariffName        string    `json:"tariff_name"`
	DiscountId        int64     `json:"discount_id"`
	DiscountAmount    int32     `json:"discount_amount"`
	Status            string    `json:"status"`
	NextBillingDate   time.Time `json:"next_billing_date"`
	TeamCount         int32     `json:"team_count"`
	WorkplaceCount    int32     `json:"workplace_count"`
	UsersCount        int32     `json:"users_count"`
	FreeWorkplaces    int32     `json:"free_workplaces"`
	PaidWorkplaces    int32     `json:"paid_workplaces"`
}

type CreatePersonalAccountRequest struct {
	OwnerUuid string `json:"owner_uuid"`
	TeamUuid  string `json:"team_uuid"`
}

type CreatePersonalAccountResponse struct {
	PersonalAccount
}

type GetPersonalAccountByIDRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

type GetPersonalAccountByIDResponse struct {
	PersonalAccount
}

type GetPersonalAccountsListRequest struct {
	PersonalAccountId int64    `json:"personal_account_id,omitempty"`
	Options           *Options `json:"options,omitempty"`
}

type Options struct {
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

type GetPersonalAccountsListResponse struct {
	PersonalAccounts []PersonalAccount `json:"personal_accounts"`
}

type ActivatePersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

type ActivatePersonalAccountResponse struct {
	Success bool `json:"success"`
}

type BlockPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

type BlockPersonalAccountResponse struct {
	Success bool `json:"success"`
}

type UnblockPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

type UnblockPersonalAccountResponse struct {
	Success bool `json:"success"`
}

type SuspendPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

type SuspendPersonalAccountResponse struct {
	Success bool `json:"success"`
}
