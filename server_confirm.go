package tdproto

func NewServerConfirm(v string) (r ServerConfirm) {
	r.BaseEvent.Name = "server.confirm"
	r.Params.ConfirmId = v
	return r
}

type ServerConfirm struct {
	BaseEvent
	Params serverConfirmParams `json:"params"`
}

type serverConfirmParams struct {
	ConfirmId string `json:"confirm_id"`
}
