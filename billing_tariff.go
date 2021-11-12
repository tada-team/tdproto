package tdproto

import "time"

// Tariff struct of billing API
type TariffBilling struct {
	// Date of closing tariff
	CloseDate *time.Time `json:"close_date,omitempty"`

	// Cost of one workplace
	CostWorkplace string `json:"cost_workplace"`

	// Currency of tariff
	Currency Currency `json:"currency"`

	// Disk space limit per user
	DiskSpaceQuotaMb string `json:"disk_space_quota_mb"`

	// Count of free workspaces
	FreeWorkplaces int32 `json:"free_workplaces"`

	// Flag of availability of free seats when exceeding FreeWorkplace
	IsBillingFree bool `json:"is_billing_free"`

	// Flag of accounting without looking at the number of days before the billing period
	IsBillingFullTime bool `json:"is_billing_full_time"`

	// Default tariff flag that is set when registering an account
	IsDefaultTariff bool `json:"is_default_tariff"`

	// Flag for accounting for unspent days when switching to a new tariff
	IsRecalcChangeTariff bool `json:"is_recalc_change_tariff"`

	// Maximum count of users in voice conference
	MaxVoiceUser int32 `json:"max_voice_user"`

	// Maximum count of users in video conference
	MaxVideoUser int32 `json:"max_video_user"`

	// Date of opening tariff
	OpenDate *time.Time `json:"open_date"`

	// Number of paid days
	PeriodDays int32 `json:"period_days"`

	// Status of tariff
	Status TariffStatus `json:"status"`

	// Tariff id
	TariffId int64 `json:"tariff_id"`

	// Name of tariff
	TariffName string `json:"tariff_name"`
}

// Request to create the tariff
type CreateTariffRequest struct {
	// Date of closing tariff
	CloseDate *time.Time `json:"close_date,omitempty"`

	// Cost of one workplace
	CostWorkplace string `json:"cost_workplace,omitempty"`

	// Currency of tariff
	Currency Currency `json:"currency"`

	// Disk space limit per user
	DiskSpaceQuotaMb string `json:"disk_space_quota_mb,omitempty"`

	// Count of free workspaces
	FreeWorkplaces int32 `json:"free_workplaces,omitempty"`

	// Flag of availability of free seats when exceeding FreeWorkplace
	IsBillingFree bool `json:"is_billing_free,omitempty"`

	// Flag of accounting without looking at the number of days before the billing period
	IsBillingFullTime bool `json:"is_billing_full_time,omitempty"`

	// Default tariff flag that is set when registering an account
	IsDefaultTariff bool `json:"is_default_tariff,omitempty"`

	// Flag for accounting for unspent days when switching to a new tariff
	IsRecalcChangeTariff bool `json:"is_recalc_change_tariff,omitempty"`

	// Maximum count of users in voice conference
	MaxVoiceUser int32 `json:"max_voice_user,omitempty"`

	// Maximum count of users in video conference
	MaxVideoUser int32 `json:"max_video_user,omitempty"`

	// Date of opening tariff
	OpenDate *time.Time `json:"open_date,omitempty"`

	// Number of paid days
	PeriodDays int32 `json:"period_days"`

	// Name of tariff
	TariffName string `json:"tariff_name"`
}

// Response from getting a list of tariffs
type GetTariffsListResponse struct {
	Tariffs []TariffBilling `json:"tariffs"`
}

// Response from getting a list of active tariffs
type GetActiveTariffsListResponse struct {
	Tariffs []TariffBilling `json:"tariffs"`
}

// Request to close (archive) the tariff
type CloseTariffRequest struct {
	CloseDate time.Time `json:"close_date,omitempty"`
}

// Response from closing (archiving) the tariff
type CloseTariffResponse struct {
	Success bool `json:"success"`
}

// Response from setting the tariff as default
type SetTariffAsDefaultResponse struct {
	Success bool `json:"success"`
}
