package tdproto

import "time"

// ChangeTariffBilling struct of changes tariffs by personal account
type ChangeTariffBilling struct {
	PersonalAccountId int64     `json:"personal_account_id"`
	TariffId          int64     `json:"tariff_id,omitempty"`
	OpenDate          time.Time `json:"open_date,omitempty"`
	CloseDate         time.Time `json:"close_date,omitempty"`
	CreateDate        time.Time `json:"create_date,omitempty"`
}

// GetChangesTariffsByPersonalAccountRequest request on get changes tariffs by personal account
type GetChangesTariffsByPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

// GetChangesTariffsByPersonalAccountResponse response on get changes tariffs by personal account
type GetChangesTariffsByPersonalAccountResponse struct {
	TariffsChanges []*ChangeTariffBilling `json:"tariffs_changes,omitempty"`
}

// CreateChangeTariffOnPersonalAccountRequest request on create change tariff on personal account
type CreateChangeTariffOnPersonalAccountRequest struct {
	PersonalAccountId int64     `json:"personal_account_id"`
	TariffId          int64     `json:"tariff_id"`
	OpenDate          time.Time `json:"open_date,omitempty"`
}

// CreateChangeTariffOnPersonalAccountResponse response on create change tariff on personal account
type CreateChangeTariffOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteChangeTariffOnPersonalAccountRequest request on delete change tariff on personal account
type DeleteChangeTariffOnPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
	TariffId          int64 `json:"tariff_id"`
}

// DeleteChangeTariffOnPersonalAccountResponse response on delete change tariff om personal account
type DeleteChangeTariffOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}
