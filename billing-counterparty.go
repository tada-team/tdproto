package tdproto

import "time"

type Counterparty struct {
	Id                                string           `json:"id"`
	PersonalAccountId                 string           `json:"personal_account_id"`
	ElectronicDocumentManagementId    string           `json:"electronic_document_management_id,omitempty"`
	FullName                          string           `json:"full_name"`
	TaxpayerIdentificationNumber      string           `json:"taxpayer_identification_number"`
	LegalAddress                      string           `json:"legal_address"`
	PhysicalAddress                   string           `json:"physical_address"`
	AccountingDictionaryCode          string           `json:"accounting_dictionary_code,omitempty"`
	ClassifierOfIndustrialEnterprises string           `json:"classifier_of_industrial_enterprises,omitempty"`
	CreatedAt                         time.Time        `json:"created_at"`
	CounterpartyType                  CounterpartyType `json:"counterparty_type"`
}

type CounterpartyCreateRequest struct {
	PersonalAccountId                 string           `json:"personal_account_id"`
	ElectronicDocumentManagementId    string           `json:"electronic_document_management_id,omitempty"`
	FullName                          string           `json:"full_name"`
	TaxpayerIdentificationNumber      string           `json:"taxpayer_identification_number"`
	LegalAddress                      string           `json:"legal_address"`
	PhysicalAddress                   string           `json:"physical_address"`
	AccountingDictionaryCode          string           `json:"accounting_dictionary_code,omitempty"`
	ClassifierOfIndustrialEnterprises string           `json:"classifier_of_industrial_enterprises,omitempty"`
	CounterpartyType                  CounterpartyType `json:"counterparty_type"`
}

type CounterpartyCreateResponse struct {
	Counterparty
}

type CounterpartyUpdateRequest struct {
	Id                                string           `json:"id"`
	PersonalAccountId                 string           `json:"personal_account_id"`
	ElectronicDocumentManagementId    string           `json:"electronic_document_management_id,omitempty"`
	FullName                          string           `json:"full_name"`
	TaxpayerIdentificationNumber      string           `json:"taxpayer_identification_number"`
	LegalAddress                      string           `json:"legal_address"`
	PhysicalAddress                   string           `json:"physical_address"`
	AccountingDictionaryCode          string           `json:"accounting_dictionary_code,omitempty"`
	ClassifierOfIndustrialEnterprises string           `json:"classifier_of_industrial_enterprises,omitempty"`
	CounterpartyType                  CounterpartyType `json:"counterparty_type"`
}

type CounterpartyUpdateResponse struct {
	Counterparty
}

type CounterpartyGetRequest struct {
	CounterpartyIds          string `json:"counterparty_ids,omitempty"`
	AccountingDictionaryCode string `json:"accounting_dictionary_code,omitempty"`
	PersonalAccountId        string `json:"personal_account_id,omitempty"`
	Limit                    uint32 `json:"limit,omitempty"`
	Offset                   uint32 `json:"offset,omitempty"`
}

type CounterpartyGetResponse struct {
	CounterpartyList []Counterparty `json:"counterparty_list"`
}

type CounterpartyDeleteRequest struct {
	CounterpartyIds string `json:"counterparty_ids,omitempty"`
}

type CounterpartyType string

const (
	CounterpartyTypeUnspecified  CounterpartyType = "COUNTERPARTY_TYPE_UNSPECIFIED"
	CounterpartyTypePhysical     CounterpartyType = "COUNTERPARTY_TYPE_PHYSICAL"
	CounterpartyTypeSelfemployed CounterpartyType = "COUNTERPARTY_TYPE_SELFEMPLOYED"
	CounterpartyTypeEntrepreneur CounterpartyType = "COUNTERPARTY_TYPE_ENTREPRENEUR"
	CounterpartyTypeJuridical    CounterpartyType = "COUNTERPARTY_TYPE_JURIDICAL"
)
