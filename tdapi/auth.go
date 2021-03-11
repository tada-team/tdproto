package tdapi

import "github.com/tada-team/tdproto"

// RecoveryStatus:
const (
	Unknown2fa     = "unknown"
	Unconfirmed2fa = "unconfirmed"
	Confirmed2fa   = "confirmed"
	Declined2fa    = "declined"
)

type Auth struct {
	Token       string             `json:"token,omitempty"`
	Me          tdproto.UserWithMe `json:"me"`
	Required2fa bool               `json:"required2fa"`
	Recovery2fa bool               `json:"recovery2fa"`
	Method2fa   string             `json:"method2fa"`
}

type Auth2faSettingsResponse struct {
	Enabled        bool   `json:"enabled"`
	RecoveryStatus string `json:"recovery_status"`
}

type Auth2faMailRecovery struct {
	CodeValidUntil string `json:"code_valid_until"`
	NextCodeAt     string `json:"next_code_at"`
	CodeLength     int    `json:"code_length"`
}

type Auth2faSettingsMailValidation struct {
	Enabled        bool   `json:"enabled"`
	RecoveryStatus string `json:"recovery_status"`

	CodeValidUntil string `json:"code_valid_until"`
	NextCodeAt     string `json:"next_code_at"`
	CodeLength     int    `json:"code_length"`
}

type basicPasswordUpdate struct {
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
	Hint              string `json:"hint"`
}

type basicToken struct {
	Token string `json:"token"`
}

type basicPassword struct {
	Password string `json:"password"`
}

type basicEmail struct {
	Email string `json:"email"`
}

type basicCode struct {
	Code string `json:"code"`
}

type Auth2faForm struct {
	basicToken
	basicPassword
}

type Create2faPasswordForm struct {
	basicPasswordUpdate
}

type Update2faPasswordForm struct {
	basicPasswordUpdate
	basicPassword
}

type Internal2faPasswordForm struct {
	basicPassword
}

type SendMail2faConfirmForm struct {
	basicPassword
	basicEmail
}

type ConfirmMail2faForm struct {
	basicPassword
	basicCode
	basicEmail
}

type AuthToken2faForm struct {
	basicToken
}

type AuthCheckCode2faForm struct {
	basicToken
	basicCode
}

type AuthPasswordRecovery2faForm struct {
	basicPasswordUpdate
	basicToken
	basicCode
}
