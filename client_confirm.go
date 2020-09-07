package tdproto

func NewClientConfirm(confirmId string) (r ClientConfirm) {
	r.Name = "client.confirm"
	r.Params.ConfirmId = confirmId
	return r
}

type ClientConfirm struct {
	BaseEvent
	Params ClientConfirmParams `json:"params"`
}

type ClientConfirmParams struct {
	ConfirmId string `json:"confirm_id"`
}
