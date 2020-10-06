package tdproto

func NewServerRemindUpdated(remind Remind) (r ServerRemindUpdated) {
	r.Name = r.GetName()
	r.Params.Reminds = []Remind{remind}
	return r
}

type ServerRemindUpdated struct {
	BaseEvent
	Params serverRemindUpdatedParams `json:"params"`
}

func (p ServerRemindUpdated) GetName() string { return "server.remind.updated" }

type serverRemindUpdatedParams struct {
	Reminds []Remind `json:"reminds"`
}
