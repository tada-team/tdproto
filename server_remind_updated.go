package tdproto

func NewServerRemindUpdated(reminds ...Remind) (r ServerRemindUpdated) {
	r.Name = r.GetName()
	r.Params.Reminds = reminds
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
