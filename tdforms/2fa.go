package tdforms


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
