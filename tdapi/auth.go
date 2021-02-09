package tdapi

import "github.com/tada-team/tdproto"

type Auth2faRequiredResponse struct {
	Token       string             `json:"token,omitempty"`
	Me          tdproto.UserWithMe `json:"me"`
	Required2fa bool               `json:"required2fa"`
}

type Auth2faForm struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type Auth2faSettingsResponse struct {
	Enabled bool `json:"enabled"`
}

type Auth2faSettingsSetForm struct {
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
	Hint              string `json:"hint"`
}

type Auth2faSettingsUpdateForm struct {
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
	Hint              string `json:"hint"`
}

type Auth2faSettingsDeleteForm struct {
	OldPassword string `json:"old_password"`
}
