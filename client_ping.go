package tdproto

func NewClientPing() (r ClientPing) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	return r
}

// Empty message for checking server connection
type ClientPing struct {
	BaseEvent
}

func (p ClientPing) GetName() string { return "client.ping" }
