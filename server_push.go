package tdproto

func NewServerMessagePush(p MessagePush) (r ServerMessagePush) {
	r.Name = r.GetName()
	r.Params = p
	return r
}

// Push replacement for desktop application
type ServerMessagePush struct {
	BaseEvent
	Params MessagePush `json:"params"`
}

func (p ServerMessagePush) GetName() string { return "server.message.push" }
