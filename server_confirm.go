package tdproto

func NewServerConfirm(v string) (r ServerConfirm) {
	r.Name = r.GetName()
	r.Params.ConfirmId = v
	return r
}

// Server got message from client.
type ServerConfirm struct {
	BaseEvent
	Params serverConfirmParams `json:"params"`
}

func (p ServerConfirm) GetName() string { return "server.confirm" }

type serverConfirmParams struct {
	ConfirmId string `json:"confirm_id"`
}
