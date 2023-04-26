package tdapi

const (
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
)

type EnquiryStatus string
type EnquiryPaymentStatus string
