package tdproto

func NewServerRemindFired(remind Remind) (r ServerRemindFired) {
	r.Name = r.GetName()
	r.Params.Reminds = []Remind{remind}
	return r
}

type ServerRemindFired struct {
	BaseEvent
	Params serverRemindFiredParams `json:"params"`
}

func (p ServerRemindFired) GetName() string { return "server.remind.fired" }

type serverRemindFiredParams struct {
	Reminds []Remind `json:"reminds"`
}
