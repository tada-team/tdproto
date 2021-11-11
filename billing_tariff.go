package tdproto

import "time"

// TariffStatus is status of tariff
type TariffStatus string

const (
	// ActiveTariff is active tariff status
	ActiveTariff TariffStatus = "Active"

	// ArchiveTariff is archive tariff status
	ArchiveTariff TariffStatus = "Archive"
)

// TariffBilling struct of billing api
type TariffBilling struct {

	// Tariff id
	TariffId int64 `json:"tariff_id,omitempty"`

	// Name of tariff
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

	// Status of tariff
	Status TariffStatus `json:"status,omitempty"`
}

// CreateTariffRequest request on create tariff
type CreateTariffRequest struct {
	TariffBilling
}

// CreateTariffResponse response on create tariff
type CreateTariffResponse struct {
	TariffBilling
}

// GetTariffByIdRequest request on get tariff by ID
type GetTariffByIdRequest struct {
	Id int64 `json:"id,omitempty"`
}

// GetTariffByIdResponse response on get tariff by ID
type GetTariffByIdResponse struct {
	TariffBilling
}

// GetTariffsListResponse response on get tariffs list
type GetTariffsListResponse struct {
	Tariffs []TariffBilling `json:"tariffs"`
}

// GetActiveTariffsListResponse response on get active tariffs list
type GetActiveTariffsListResponse struct {
	Tariffs []TariffBilling `json:"tariffs"`
}

// CloseTariffRequest request on close(archive) tariff
type CloseTariffRequest struct {
	TariffId  int64     `json:"tariff_id,omitempty"`
	CloseDate time.Time `json:"close_date,omitempty"`
}

// CloseTariffResponse response on close(archive) tariff
type CloseTariffResponse struct {
	Success bool `json:"success"`
}

// SetTariffAsDefaultRequest request on set tariff as default
type SetTariffAsDefaultRequest struct {
	TariffId int64 `json:"tariff_id,omitempty"`
}

// SetTariffAsDefaultResponse response on set tariff as default
type SetTariffAsDefaultResponse struct {
	Success bool `json:"success"`
}
