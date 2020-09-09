package tdproto

func NewServerRemindFired(remind Remind) (r ServerRemindFired) {
	r.BaseEvent.Name = "server.remind.fired"
	r.Params.Reminds = []Remind{remind}
	return r
}

type ServerRemindFired struct {
	BaseEvent
	Params serverRemindFiredParams `json:"params"`
}

type serverRemindFiredParams struct {
	Reminds []Remind `json:"reminds"`
}
