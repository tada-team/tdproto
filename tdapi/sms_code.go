package tdapi

type SmsCode struct {
	Phone          string `json:"phone"`
	CodeValidUntil string `json:"code_valid_until"`
	NextCodeAt     string `json:"next_code_at"`
	CodeLength     int    `json:"code_length"`
}
