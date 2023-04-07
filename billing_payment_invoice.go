package tdproto

import "time"

type PaymentInvoice struct {
	Id                 string    `json:"id"`
	EnquiryId          string    `json:"enquiry_id"`
	CounterpartyId     string    `json:"counterparty_id"`
	Status             string    `json:"status"`
	TariffName         string    `json:"tariff_name,omitempty"`
	CostWorkplace      string    `json:"cost_workplace,omitempty"`
	WorkplaceCount     string    `json:"workplace,omitempty"`
	FreeWorkplaceCount string    `json:"free_workplace,omitempty"`
	PaidAt             time.Time `json:"paid_at,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	ActivateAt         time.Time `json:"activate_at,omitempty"`
	Amount             float64   `json:"amount"`
}

type PaymentInvoiceCreateRequest struct {
	EnquiryId      string    `json:"enquiry_id"`
	CounterpartyId string    `json:"counterparty_id"`
	PaidAt         time.Time `json:"paid_at,omitempty"`
	Status         string    `json:"status"`
}

type PaymentInvoiceCreateResponse struct {
	PaymentInvoice
}

type PaymentInvoiceUpdateRequest struct {
	EnquiryId      string    `json:"enquiry_id"`
	CounterpartyId string    `json:"counterparty_id"`
	PaidAt         time.Time `json:"paid_at"`
	Status         string    `json:"status"`
}

type PaymentInvoiceUpdateResponse struct {
	PaymentInvoice
}

type PaymentInvoiceGetRequest struct {
	PaymentInvoiceIds string `json:"payment_invoice_ids"`
}

type PaymentInvoiceGetListRequest struct {
	EnquiryId      string `json:"enquiry_id,omitempty"`
	CounterpartyId string `json:"counterparty_id,omitempty"`
	Limit          uint32 `json:"limit,omitempty"`
	Offset         uint32 `json:"offset,omitempty"`
	Status         string `json:"status,omitempty"`
}

type PaymentInvoiceGetListResponse struct {
	PaymentInvoiceList []PaymentInvoice `json:"payment_invoice_list"`
}
