package tdproto

func NewServerRemindUpdated(remind Remind) (r ServerRemindUpdated) {
	r.BaseEvent.Name = "server.remind.updated"
	r.Params.Reminds = []Remind{remind}
	return r
}

type ServerRemindUpdated struct {
	BaseEvent
	Params serverRemindUpdatedParams `json:"params"`
}

type serverRemindUpdatedParams struct {
	Reminds []Remind `json:"reminds"`
}
