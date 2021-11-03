package tdproto

import "time"

type CreateTariffRequest struct {
	TariffName           string    `json:"tariff_name,omitempty"`
	FreeWorkplaces       int64     `json:"free_workplaces,omitempty"`
	DiskSpaceQuotaMb     float64   `json:"disk_space_quota_mb,omitempty"`
	IsBillingFree        bool      `json:"is_billing_free,omitempty"`
	IsBillingFullTime    bool      `json:"is_billing_full_time,omitempty"`
	PeriodDays           int64     `json:"period_days,omitempty"`
	CostWorkplace        float64   `json:"cost_workplace,omitempty"`
	Currency             string    `json:"currency,omitempty"`
	IsRecalcChangeTariff bool      `json:"is_recalc_change_tariff,omitempty"`
	MaxVoiceUser         int64     `json:"max_voice_user,omitempty"`
	MaxVideoUser         int64     `json:"max_video_user,omitempty"`
	IsDefaultTariff      bool      `json:"is_default_tariff,omitempty"`
	OpenDate             time.Time `json:"open_date,omitempty"`
	CloseDate            time.Time `json:"close_date,omitempty"`
}

type CreateTariffResponse struct {
	TariffId             int64     `json:"tariff_id"`
	TariffName           string    `json:"tariff_name"`
	FreeWorkplaces       int64     `json:"free_workplaces"`
	DiskSpaceQuotaMb     float64   `json:"disk_space_quota_mb"`
	IsBillingFree        bool      `json:"is_billing_free"`
	IsBillingFullTime    bool      `json:"is_billing_full_time"`
	PeriodDays           int64     `json:"period_days"`
	CostWorkplace        float64   `json:"cost_workplace"`
	Currency             string    `json:"currency"`
	IsRecalcChangeTariff bool      `json:"is_recalc_change_tariff"`
	MaxVoiceUser         int64     `json:"max_voice_user"`
	MaxVideoUser         int64     `json:"max_video_user"`
	IsDefaultTariff      bool      `json:"is_default_tariff"`
	OpenDate             time.Time `json:"open_date"`
	CloseDate            time.Time `json:"close_date"`
}

type GetTariffByIdRequest struct {
	Id int64 `json:"id"`
}

type GetTariffByIdResponse struct {
	TariffId             int64     `json:"tariff_id"`
	TariffName           string    `json:"tariff_name"`
	FreeWorkplaces       int64     `json:"free_workplaces"`
	DiskSpaceQuotaMb     float64   `json:"disk_space_quota_mb"`
	IsBillingFree        bool      `json:"is_billing_free"`
	IsBillingFullTime    bool      `json:"is_billing_full_time"`
	PeriodDays           int64     `json:"period_days"`
	CostWorkplace        float64   `json:"cost_workplace"`
	Currency             string    `json:"currency"`
	IsRecalcChangeTariff bool      `json:"is_recalc_change_tariff"`
	MaxVoiceUser         int64     `json:"max_voice_user"`
	MaxVideoUser         int64     `json:"max_video_user"`
	IsDefaultTariff      bool      `json:"is_default_tariff"`
	OpenDate             time.Time `json:"open_date"`
	CloseDate            time.Time `json:"close_date"`
	Status               string    `json:"status"`
}

type GetTariffsListRequest struct{}

type GetTariffsListResponse struct {
	Tariffs []GetTariffByIdResponse `json:"tariffs"`
}

type GetActiveTariffsListRequest struct{}

type GetActiveTariffsListResponse struct {
	Tariffs []GetTariffByIdResponse `json:"tariffs"`
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

type CreatePersonalAccountRequest struct {
	OwnerUuid string `json:"owner_uuid"`
	TeamUuid  string `json:"team_uuid"`
}

type CreatePersonalAccountResponse struct {
	PersonalAccountId int64     `json:"personal_account_id"`
	OwnerUuid         string    `json:"owner_uuid"`
	TariffId          int64     `json:"tariff_id"`
	TariffName        string    `json:"tariff_name"`
	DiscountId        int64     `json:"discount_id"`
	DiscountAmount    int64     `json:"discount_amount"`
	Status            string    `json:"status"`
	NextBillingDate   time.Time `json:"next_billing_date"`
	TeamCount         int64     `json:"team_count"`
	WorkplaceCount    int64     `json:"workplace_count"`
	UsersCount        int64     `json:"users_count"`
	FreeWorkplaces    int64     `json:"free_workplaces"`
	PaidWorkplaces    int64     `json:"paid_workplaces"`
}

type GetPersonalAccountByIDRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

type GetPersonalAccountByIDResponse struct {
	PersonalAccountId int64     `json:"personal_account_id"`
	OwnerUuid         string    `json:"owner_uuid"`
	TariffId          int64     `json:"tariff_id"`
	TariffName        string    `json:"tariff_name"`
	DiscountId        int64     `json:"discount_id"`
	DiscountAmount    int64     `json:"discount_amount"`
	Status            string    `json:"status"`
	NextBillingDate   time.Time `json:"next_billing_date"`
	TeamCount         int64     `json:"team_count"`
	WorkplaceCount    int64     `json:"workplace_count"`
	UsersCount        int64     `json:"users_count"`
	FreeWorkplaces    int64     `json:"free_workplaces"`
	PaidWorkplaces    int64     `json:"paid_workplaces"`
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
	PersonalAccounts []GetPersonalAccountByIDResponse `json:"personal_accounts"`
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
