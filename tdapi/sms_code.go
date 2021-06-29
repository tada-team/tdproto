package tdapi

import "github.com/tada-team/tdproto"

// Sms code response
type SmsCode struct {
	// Phone number
	Phone string `json:"phone"`

	// Code expiration date
	CodeValidUntil tdproto.ISODateTimeString `json:"code_valid_until"`

	// Next code date
	NextCodeAt tdproto.ISODateTimeString `json:"next_code_at"`

	// Code length, symbols
	CodeLength int `json:"code_length"`
}
