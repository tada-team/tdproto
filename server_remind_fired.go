package tdproto

func NewServerRemindFired(remind Remind) (r ServerRemindFired) {
	r.Name = r.GetName()
	r.Params.Reminds = []Remind{remind}
	return r
}

// Task or group remind fired
type ServerRemindFired struct {
	BaseEvent
	Params serverRemindFiredParams `json:"params"`
}

func (p ServerRemindFired) GetName() string { return "server.remind.fired" }

// Params of the server.remind.fired event
type serverRemindFiredParams struct {
	// Remind information
	Reminds []Remind `json:"reminds"`
}
