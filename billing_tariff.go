package tdproto

import "time"

// Tariff struct of billing API
type TariffBilling struct {
	// Tariff id
	Id uint64 `json:"id"`

	// Name of tariff
	Name string `json:"name"`

	// Nomenclature name of tariff
	NomenclatureName string `json:"nomenclature_name"`

	// Description of tariff
	Description string `json:"description"`

	// Benefit of tariff
	Benefit string `json:"benefit,omitempty"`

	// Currency of tariff
	Currency Currency `json:"currency"`

	// Priority of tariff
	Priority uint32 `json:"priority,omitempty"`

	// Cost of one workplace
	CostWorkplace float64 `json:"cost_workplace"`

	// Count of maximum workspaces on tariff
	MaxWorkplaceCount uint32 `json:"max_workplace_count,omitempty"`

	// Count of minimum workspaces on tariff
	MinWorkplaceCount uint32 `json:"min_workplace_count,omitempty"`

	// Count of free workspaces
	FreeWorkplaceCount uint32 `json:"free_workplace_count"`

	// Minimum step of change count workspaces on tariff
	StepIncreasingWorkplaces uint32 `json:"step_increasing_workplaces"`

	// Disk space limit per user
	DiskSpaceQuotaMb float64 `json:"disk_space_quota_mb"`

	// Number of paid days
	PeriodDays uint32 `json:"period_days"`

	// Maximum count of users in voice conference
	MaxVoiceUser uint32 `json:"max_voice_user"`

	// Maximum count of users in video conference
	MaxVideoUser uint32 `json:"max_video_user"`

	// Bitrate of video in video co
	VideoCallBitrate uint32 `json:"video_call_bitrate"`

	// Bitrate of video in video sharing
	VideoSharingBitrate uint32 `json:"video_sharing_bitrate"`

	// Flag of availability of free seats when exceeding FreeWorkplace
	IsFree bool `json:"is_free"`

	// Flag of publicity
	IsPublic bool `json:"is_public"`

	// Default tariff flag that is set when registering an account
	IsDefault bool `json:"is_default"`

	// Date of opening tariff
	OpenDate *time.Time `json:"open_date"`

	// Date of closing tariff
	CloseDate *time.Time `json:"close_date,omitempty"`

	// Status of tariff
	Status TariffStatus `json:"status"`
}

// Request to create the tariff
type CreateTariffRequest struct {
	// Name of tariff
	Name string `json:"name"`

	// Nomenclature name of tariff
	NomenclatureName string `json:"nomenclature_name"`

	// Description of tariff
	Description string `json:"description,omitempty"`

	// Benefit of tariff
	Benefit string `json:"benefit,omitempty"`

	// Currency of tariff
	Currency Currency `json:"currency"`

	// Cost of one workplace
	CostWorkplace string `json:"cost_workplace"`

	// Priority of tariff
	Priority uint32 `json:"priority,omitempty"`

	// Count of maximum workspaces on tariff
	MaxWorkplaceCount uint32 `json:"max_workplace_count,omitempty"`

	// Count of minimum workspaces on tariff
	MinWorkplaceCount uint32 `json:"min_workplace_count,omitempty"`

	// Count of free workspaces
	FreeWorkplaceCount uint32 `json:"free_workplace_countt,omitempty"`

	// Minimum step of change count workspaces on tariff
	StepIncreasingWorkplaces uint32 `json:"step_increasing_workplaces,omitempty"`

	// Disk space limit per user
	DiskSpaceQuotaMb string `json:"disk_space_quota_mb,omitempty"`

	// Number of paid days
	PeriodDays uint32 `json:"period_days"`

	// Maximum count of users in voice conference
	MaxVoiceUser uint32 `json:"max_voice_user,omitempty"`

	// Maximum count of users in video conference
	MaxVideoUser uint32 `json:"max_video_user,omitempty"`

	// Bitrate of video in video co
	VideoCallBitrate uint32 `json:"video_call_bitrate,omitempty"`

	// Bitrate of video in video sharing
	VideoSharingBitrate uint32 `json:"video_sharing_bitrate,omitempty"`

	// Flag of availability of free seats when exceeding FreeWorkplace
	IsFree bool `json:"is_free,omitempty"`

	// Flag of publicity
	IsPublic bool `json:"is_public,omitempty"`

	// Default tariff flag that is set when registering an account
	IsDefault bool `json:"is_default,omitempty"`

	// Date of opening tariff
	OpenDate *time.Time `json:"open_date,omitempty"`
}

// Request to update the tariff
type UpdateTariffRequest struct {
	// Tariff id
	TariffId string `json:"tariff_id"`

	// Date of closing tariff
	CloseDate *time.Time `json:"close_date,omitempty"`

	// Default tariff flag that is set when registering an account
	IsDefault bool `json:"is_default,omitempty"`

	// Status of tariff
	Status TariffStatus `json:"status,omitempty"`
}

// Response from getting a list of tariffs
type GetTariffsListResponse struct {
	Tariffs []TariffBilling `json:"tariffs"`
}

// Response from getting a list of active tariffs
type GetActiveTariffsListResponse struct {
	Tariffs []TariffBilling `json:"tariffs"`
}
