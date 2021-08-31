package tdproto

type Payment struct {
	TariffUid       string `json:"tariff_uid"`
	TeamUid         string `json:"team_uid"`
	UserUid         string `json:"user_uid"`
	ConfirmationUrl string `json:"confirmation_url"`
}
