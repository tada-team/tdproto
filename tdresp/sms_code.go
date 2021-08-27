package tdresp

import "github.com/tada-team/tdproto"

// SMS code response
type SmsCode struct {
	// Phone number
	Phone string `json:"phone"`

	// Code expiration date
	CodeValidUntil tdproto.ISODateTimeString `json:"code_valid_until"`

	// Next code date
	NextCodeAt tdproto.ISODateTimeString `json:"next_code_at"`

	// SMS code length
	CodeLength int `json:"code_length"`
}
