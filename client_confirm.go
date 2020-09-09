package tdproto

func NewClientConfirm(confirmId string) (r ClientConfirm) {
	r.Name = "client.confirm"
	r.Params.ConfirmId = confirmId
	return r
}

type ClientConfirm struct {
	BaseEvent
	Params clientConfirmParams `json:"params"`
}

type clientConfirmParams struct {
	ConfirmId string `json:"confirm_id"`
}
