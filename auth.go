package tdproto

type Auth struct {
	Token       string     `json:"token,omitempty"`
	Me          UserWithMe `json:"me"`
	Required2fa bool       `json:"required2fa"`
	Recovery2fa bool       `json:"recovery2fa"`
	Method2fa   string     `json:"method2fa"`
}
