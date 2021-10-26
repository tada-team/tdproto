package tdproto

import "time"

// Tariff for teams
type Tariff struct {
	// Tariff id
	Uid string `json:"uid"`

	// Title of tariff on enlish
	TitleEn string `json:"title_en"`

	// Title of tariff on russian
	TitleRu string `json:"title_ru"`

	// Cloud space reserved for storing team users uploads in megabytes
	CloudSpace int64 `json:"cloud_space,omitempty"`

	// Maximum allowed number of members in a team
	MaxMembersInTeam int `json:"max_members_in_team,omitempty"`

	// Maximum number of participants per call
	MaxParticipantsPerCall int `json:"max_participants_per_call,omitempty"`

	//maximum file size for uploading
	MaxUploadFilesize int64 `json:"max_upload_filesize,omitempty"`
}

// TariffBilling struct of billing api
type TariffBilling struct {

	// Tariff id
	TariffId int64 `json:"tariff_id"`

	// Name of tariff on russian
	TariffName string `json:"tariff_name"`

	// Count of free workspaces
	FreeWorkplace int64 `json:"free_workplace"`

	// Disk space limit per user
	DiskSpaceQuota string `json:"disk_space_quota"`

	// Flag of availability of free seats when exceeding FreeWorkplace
	BillingFree bool `json:"billing_free"`

	// Flag of accounting without looking at the number of days before the billing period
	BillingFullTime bool `json:"billing_full_time"`

	// Number of paid days
	PeriodDays int64 `json:"period_days"`

	// Cost of one workplace
	CostWorkplace string `json:"cost_workplace"`

	// Currency of tariff in ISO
	Currency string `json:"currency"`

	// Flag for accounting for unspent days when switching to a new tariff
	RecalcChangeTariff bool `json:"recalc_change_tariff"`

	// Maximum count of users in voice conference
	MaxVoiceUser int64 `json:"max_voice_user"`

	// Maximum count of users in video conference
	MaxVideoUser int64 `json:"max_video_user"`

	// Default tariff flag that is set when registering an account
	DefaultTariff bool `json:"default_tariff"`

	// Date of opening tariff
	OpenDate *time.Time `json:"open_date"`

	// Date of closing tariff
	CloseDate *time.Time `json:"close_date"`
}
