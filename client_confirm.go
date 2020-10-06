package tdproto

func NewClientConfirm(confirmId string) (r ClientConfirm) {
	r.Name = r.GetName()
	r.Params.ConfirmId = confirmId
	return r
}

type ClientConfirm struct {
	BaseEvent
	Params clientConfirmParams `json:"params"`
}

func (p ClientConfirm) GetName() string { return "client.confirm" }

type clientConfirmParams struct {
	ConfirmId string `json:"confirm_id"`
}
