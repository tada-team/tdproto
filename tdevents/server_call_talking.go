package tdevents

import "github.com/tada-team/tdproto"

func NewServerCallTalking(talking bool, chat, actor tdproto.JID) (r ServerCallTalking) {
	r.Name = r.GetName()
	r.Params.Jid = chat
	r.Params.Actor = actor
	r.Params.Talking = talking
	return r
}

// Someone talks in call
type ServerCallTalking struct {
	BaseEvent
	Params serverCallTalkingParams `json:"params"`
}

func (p ServerCallTalking) GetName() string { return "server.call.talking" }

// Params of the server.call.talking event
type serverCallTalkingParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Actor id
	Actor tdproto.JID `json:"actor"`

	// Is talking
	Talking bool `json:"talking"`
}
