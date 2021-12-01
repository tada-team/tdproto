package tdproto

import "time"

// ChangeTariffBilling struct of changes tariffs by personal account
type ChangeTariffBilling struct {
	TariffId   string    `json:"tariff_id"` // TODO: must be int64
	OpenDate   time.Time `json:"open_date,omitempty"`
	CloseDate  time.Time `json:"close_date,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
}

// GetChangesTariffsByPersonalAccountResponse response on get changes tariffs by personal account
type GetChangesTariffsByPersonalAccountResponse struct {
	TariffsChanges []ChangeTariffBilling `json:"tariffs_changes,omitempty"`
}

// CreateChangeTariffOnPersonalAccountRequest request on create change tariff on personal account
type CreateChangeTariffOnPersonalAccountRequest struct {
	TariffId string    `json:"tariff_id"` // TODO: must be int64
	OpenDate time.Time `json:"open_date,omitempty"`
}

// CreateChangeTariffOnPersonalAccountResponse response on create change tariff on personal account
type CreateChangeTariffOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteChangeTariffOnPersonalAccountRequest request on delete change tariff on personal account
type DeleteChangeTariffOnPersonalAccountRequest struct {
	TariffId string `json:"tariff_id"` // TODO: must be int64
}

// DeleteChangeTariffOnPersonalAccountResponse response on delete change tariff om personal account
type DeleteChangeTariffOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}
