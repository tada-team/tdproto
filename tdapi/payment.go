package tdapi

// Form for create payment
type Payment struct {
	TariffUid string `json:"tariff_uid"`
	TeamUid   string `json:"team_uid"`
	ReturnUrl string `json:"return_url,omitempty"`
}
