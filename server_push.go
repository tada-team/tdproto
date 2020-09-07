package tdproto

func NewServerMessagePush(p MessagePush) (r ServerMessagePush) {
	r.BaseEvent.Name = "server.message.push"
	r.Params = p
	return r
}

type ServerMessagePush struct {
	BaseEvent
	Params MessagePush `json:"params"`
}
