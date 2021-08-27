package tdresp

import "github.com/tada-team/tdproto"

type Auth struct {
	Token       string             `json:"token,omitempty"`
	Me          tdproto.UserWithMe `json:"me"`
	Required2fa bool               `json:"required2fa"`
	Recovery2fa bool               `json:"recovery2fa"`
	Method2fa   string             `json:"method2fa"`
}

type Auth2faMailRecovery struct {
	CodeValidUntil tdproto.ISODateTimeString `json:"code_valid_until"`
	NextCodeAt     tdproto.ISODateTimeString `json:"next_code_at"`
	CodeLength     int                       `json:"code_length"`
	Email          string                    `json:"email"`
}

type Auth2faSettingsResponse struct {
	Enabled        bool   `json:"enabled"`
	RecoveryStatus string `json:"recovery_status"`
}

type Auth2faSettingsMailValidation struct {
	Enabled        bool                      `json:"enabled"`
	RecoveryStatus string                    `json:"recovery_status"`
	CodeValidUntil tdproto.ISODateTimeString `json:"code_valid_until"`
	NextCodeAt     string                    `json:"next_code_at"`
	CodeLength     int                       `json:"code_length"`
}
