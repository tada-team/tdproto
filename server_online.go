package tdproto

func NewServerOnline(contacts []OnlineContact, calls []OnlineCall) (r ServerOnline) {
	r.Name = r.GetName()
	r.Params.Contacts = &contacts
	r.Params.Calls = &calls
	return r
}

type ServerOnline struct {
	BaseEvent
	Params serverOnlineParams `json:"params"`
}

func (p ServerOnline) GetName() string { return "server.online" }

type serverOnlineParams struct {
	Contacts *[]OnlineContact `json:"contacts"`
	Calls    *[]OnlineCall    `json:"calls"`
}

type OnlineContact struct {
	Jid    JID  `json:"jid"`
	Afk    bool `json:"afk,omitempty"`
	Mobile bool `json:"mobile"` // TODO: omitempty. 17feb2020
}

type OnlineCall struct {
	Jid         JID     `json:"jid"`
	Uid         string  `json:"uid"`
	Start       *string `json:"start,omitempty"`
	OnlineCount int     `json:"online_count,omitempty"`
}
