package tdproto

func NewServerOnline(contacts []OnlineContact, calls []OnlineCall) (r ServerOnline) {
	r.BaseEvent.Unimportant = true
	r.BaseEvent.Name = "server.online"
	r.Params.Contacts = &contacts
	r.Params.Calls = &calls
	return r
}

type ServerOnline struct {
	BaseEvent
	Params ServerOnlineParams `json:"params"`
}

type ServerOnlineParams struct {
	Contacts *[]OnlineContact `json:"contacts"`
	Calls    *[]OnlineCall    `json:"calls"`
}

type OnlineContact struct {
	Jid    JID  `json:"jid"`
	Afk    bool `json:"afk,omitempty"`
	Mobile bool `json:"mobile"` // TODO: omitempty. 17feb2020
}

type OnlineCall struct {
	Jid JID    `json:"jid"`
	Uid string `json:"uid"`
}
