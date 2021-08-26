package tdapi

import "github.com/tada-team/tdproto"

// SMS code response
type SmsCode struct {
	// Normalized mobile phone number
	Phone string `json:"phone"`

	// Code expiration time
	CodeValidUntil tdproto.ISODateTimeString `json:"code_valid_until"`

	// Next code sending available time
	NextCodeAt tdproto.ISODateTimeString `json:"next_code_at"`

	// SMS code length
	CodeLength int `json:"code_length"`
}
