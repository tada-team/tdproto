package tdevents

import "github.com/tada-team/tdproto"

func NewServerMessagePush(p tdproto.MessagePush) (r ServerMessagePush) {
	r.Name = r.GetName()
	r.Params = p
	return r
}

// Push replacement for desktop application
type ServerMessagePush struct {
	BaseEvent
	Params tdproto.MessagePush `json:"params"`
}

func (p ServerMessagePush) GetName() string { return "server.message.push" }
