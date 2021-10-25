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
	TariffId           int64      `json:"tariff_id"`
	TariffName         string     `json:"tariff_name"`
	FreeWorkplace      int64      `json:"free_workplace"`
	DiskSpaceQuota     string     `json:"disk_space_quota"`
	BillingFree        bool       `json:"billing_free"`
	BillingFullTime    bool       `json:"billing_full_time"`
	PeriodDays         int64      `json:"period_days"`
	CostWorkplace      string     `json:"cost_workplace"`
	Currency           string     `json:"currency"`
	RecalcChangeTariff bool       `json:"recalc_change_tariff"`
	MaxVoiceUser       int64      `json:"max_voice_user"`
	MaxVideoUser       int64      `json:"max_video_user"`
	DefaultTariff      bool       `json:"default_tariff"`
	OpenDate           *time.Time `json:"open_date"`
	CloseDate          *time.Time `json:"close_date"`
}
