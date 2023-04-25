package tdproto

import "time"

const (
	EnquiryTypeUnspecified EnquiryType = "ENQUIRY_TYPE_UNSPECIFIED"
	EnquiryTypeRenewal     EnquiryType = "ENQUIRY_TYPE_RENEWAL"
	EnquiryTypeBasic       EnquiryType = "ENQUIRY_TYPE_BASIC"
	EnquiryTypeExtension   EnquiryType = "ENQUIRY_TYPE_EXTENSION"

	EnquiryStatusUnspecified EnquiryStatus = "ENQUIRY_STATUS_UNSPECIFIED"
	EnquiryStatusWaiting     EnquiryStatus = "ENQUIRY_STATUS_WAITING"
	EnquiryStatusCancelled   EnquiryStatus = "ENQUIRY_STATUS_CANCELLED"
	EnquiryStatusActive      EnquiryStatus = "ENQUIRY_STATUS_ACTIVE"
	EnquiryStatusExpired     EnquiryStatus = "ENQUIRY_STATUS_EXPIRED"
	EnquiryStatusDone        EnquiryStatus = "ENQUIRY_STATUS_DONE"

	EnquiryPaymentStatusUnspecified         EnquiryPaymentStatus = "ENQUIRY_PAYMENT_STATUS"
	EnquiryPaymentStatusWaitingConfirmation EnquiryPaymentStatus = "ENQUIRY_PAYMENT_STATUS_WAITING_CONFIRMATION"
	EnquiryPaymentStatusWaitingCrediting    EnquiryPaymentStatus = "ENQUIRY_PAYMENT_STATUS_WAITING_CREDITING"
	EnquiryPaymentStatusPaid                EnquiryPaymentStatus = "ENQUIRY_PAYMENT_STATUS_PAID"

	PaymentTypeUnspecified                  PaymentType = "PAYMENT_TYPE_UNSPECIFIED"
	PaymentTypePaperDocumentManagement      PaymentType = "PAYMENT_TYPE_PAPER_DOCUMENT_MANAGEMENT"
	PaymentTypeElectronicDocumentManagement PaymentType = "PAYMENT_TYPE_ELECTRONIC_DOCUMENT_MANAGEMENT"
	PaymentTypeBankCard                     PaymentType = "PAYMENT_TYPE_BANK_CARD"
	PaymentTypeTechnical                    PaymentType = "PAYMENT_TYPE_TECHNICAL"

	EnquiryStatusFilterPaid                EnquiryStatusFilter = "ENQUIRY_STATUS_FILTER_PAID"
	EnquiryStatusFilterExpired             EnquiryStatusFilter = "ENQUIRY_STATUS_FILTER_EXPIRED"
	EnquiryStatusFilterCancelled           EnquiryStatusFilter = "ENQUIRY_STATUS_FILTER_CANCELLED"
	EnquiryStatusFilterWaitingConfirmation EnquiryStatusFilter = "ENQUIRY_STATUS_FILTER_WAITING_CONFIRMATION"
	EnquiryStatusFilterWaitingCrediting    EnquiryStatusFilter = "ENQUIRY_STATUS_FILTER_WAITING_CREDITING"
)

type EnquiryType string
type EnquiryStatus string
type EnquiryPaymentStatus string
type PaymentType string
type EnquiryStatusFilter string

type Enquiry struct {
	Id                       string  `json:"id"`
	PersonalAccountId        string  `json:"personal_account_id"`
	SelectableWorkplaceCount uint32  `json:"selectable_workplace_count"`
	Amount                   float64 `json:"amount"`

	TariffName         string  `json:"tariff_name"`
	WorkplacePrice     float64 `json:"workplace_price"`
	PeriodDays         uint32  `json:"period_days"`
	FreeWorkplaceCount uint32  `json:"free_workplace_count"`

	ActivationDate   *time.Time `json:"activation_date,omitempty"`
	DeactivationDate *time.Time `json:"deactivation_date,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	CreditedAt       *time.Time `json:"credited_at,omitempty"`
	ActuallyPaidAt   *time.Time `json:"actually_paid_at,omitempty"`
	FixationPaidAt   *time.Time `json:"fixation_paid_at,omitempty"`
	ExpirationDate   *time.Time `json:"expiration_date,omitempty"`
	ExpiredAt        *time.Time `json:"expired_at,omitempty"`
	ActivateAt       *time.Time `json:"activate_at,omitempty"`
	DeactivateAt     *time.Time `json:"deactivate_at,omitempty"`

	EnquiryType          EnquiryType          `json:"enquiry_type"`
	EnquiryStatus        EnquiryStatus        `json:"enquiry_status"`
	EnquiryPaymentStatus EnquiryPaymentStatus `json:"enquiry_payment_status"`
	PaymentType          PaymentType          `json:"payment_type"`

	FileName string `json:"file_name,omitempty"`
	MediaUrl string `json:"media_url,omitempty"`
}

type EnquiryCreateRequest struct {
	SelectableTariffId       string `json:"selectable_tariff_id"`
	SelectableWorkplaceCount uint32 `json:"selectable_workplace_count"`
	CounterpartyId           string `json:"counterparty_id"`
}

type EnquiryCreateResponse struct {
	Enquiry
}

type EnquiryGetListRequest struct {
	DateCreateFrom      *time.Time          `json:"date_create_from,omitempty"`
	DateCreateTo        *time.Time          `json:"date_create_to,omitempty"`
	DateActivateTo      *time.Time          `json:"date_activate_to,omitempty"`
	DateDeactivateTo    *time.Time          `json:"date_deactivate_to,omitempty"`
	EnquiryStatusFilter EnquiryStatusFilter `json:"enquiry_status_filter,omitempty"`
	Limit               *uint32             `json:"limit,omitempty"`
	Offset              *uint32             `json:"offset,omitempty"`
}

type EnquiryGetListResponse struct {
	PaginatedBillingEnquiries
}
