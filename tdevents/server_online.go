package tdevents

import "github.com/tada-team/tdproto"

func NewServerOnline(contacts []OnlineContact, calls []OnlineCall) (r ServerOnline) {
	r.Name = r.GetName()
	r.Params.Contacts = contacts
	r.Params.Calls = calls
	return r
}

// Online team members and current active calls
type ServerOnline struct {
	BaseEvent
	Params serverOnlineParams `json:"params"`
}

func (p ServerOnline) GetName() string { return "server.online" }

// Params of the server.online event
type serverOnlineParams struct {
	// Online team members
	Contacts []OnlineContact `json:"contacts"`

	// Active calls
	Calls []OnlineCall `json:"calls,omitempty"`
}

// Contact online status
type OnlineContact struct {
	// Contact id
	Jid tdproto.JID `json:"jid"`

	// Is away from keyboard
	Afk bool `json:"afk,omitempty"`

	// Is mobile client
	Mobile bool `json:"mobile"` // TODO: omitempty. 17feb2020
}

// Active call status
type OnlineCall struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// Call start
	Start tdproto.ISODateTimeString `json:"start,omitempty"`

	// Number participants in call
	OnlineCount int `json:"online_count,omitempty"`
}
