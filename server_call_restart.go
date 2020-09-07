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
	Params ServerCallRestartParams `json:"params"`
}

type ServerCallRestartParams struct {
	Jid  JID    `json:"jid"`
	Team string `json:"team"`
	Uid  string `json:"uid"`
}
