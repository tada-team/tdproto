package tdproto

func NewServerRemindUpdated(remind Remind) (r ServerRemindUpdated) {
	r.BaseEvent.Name = "server.remind.updated"
	r.Params.Reminds = []Remind{remind}
	return r
}

type ServerRemindUpdated struct {
	BaseEvent
	Params ServerRemindUpdatedParams `json:"params"`
}

type ServerRemindUpdatedParams struct {
	Reminds []Remind `json:"reminds"`
}
