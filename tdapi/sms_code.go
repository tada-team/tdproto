package tdapi

import "github.com/tada-team/tdproto"

type SmsCode struct {
	Phone          string                    `json:"phone"`
	CodeValidUntil tdproto.ISODateTimeString `json:"code_valid_until"`
	NextCodeAt     tdproto.ISODateTimeString `json:"next_code_at"`
	CodeLength     int                       `json:"code_length"`
}
