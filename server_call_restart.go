package tdproto

func NewServerCallRestart(chat *JID, teamUid string, uid string) (r ServerCallRestart) {
	r.BaseEvent.Name = "server.call.restart"
	r.Params.Jid = *chat
	r.Params.Team = teamUid
	r.Params.Uid = uid
	return r
}

type ServerCallRestart struct {
	BaseEvent
	Params serverCallRestartParams `json:"params"`
}

type serverCallRestartParams struct {
	Jid  JID    `json:"jid"`
	Team string `json:"team"`
	Uid  string `json:"uid"`
}
