package tdproto

func NewClientPing() (r ClientPing) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	return r
}

type ClientPing struct {
	BaseEvent
}

func (p ClientPing) GetName() string { return "client.ping" }
