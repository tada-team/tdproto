package tdproto

func NewServerCallRestart(chat *JID, teamUid string, uid string) (r ServerCallRestart) {
	r.Name = r.GetName()
	r.Params.Jid = *chat
	r.Params.Team = teamUid
	r.Params.Uid = uid
	return r
}

// Call restarted
type ServerCallRestart struct {
	BaseEvent
	Params serverCallRestartParams `json:"params"`
}

func (p ServerCallRestart) GetName() string { return "server.call.restart" }

type serverCallRestartParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// Team id
	Team string `json:"team"`
}
