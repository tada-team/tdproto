package tdproto

func NewClientPing() (r ClientPing) {
	r.Name = "client.ping"
	r.ConfirmId = ConfirmId()
	return r
}

type ClientPing struct {
	BaseEvent
}
