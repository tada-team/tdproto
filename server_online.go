package tdproto

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
	Jid JID `json:"jid"`

	// Is away from keyboard
	Afk bool `json:"afk,omitempty"`

	// Focus mode enabled
	Focused bool `json:"focused,omitempty"`

	// Is mobile client
	Mobile bool `json:"mobile"` // TODO: omitempty. 17feb2020
}

// Active call status
type OnlineCall struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// Call start
	Start ISODateTimeString `json:"start,omitempty"`

	// Number participants in call
	OnlineCount int `json:"online_count,omitempty"`

	// CallType is a type of call("audio" - audio room, "video" - video room)
	CallType CallType `json:"call_type"`
}
