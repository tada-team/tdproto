package tdproto

func NewServerCallBuzzcancel(chat *JID, teamUid string, uid string) (r ServerCallBuzzcancel) {
	r.BaseEvent.Name = "server.call.buzzcancel"
	r.Params.Jid = *chat
	r.Params.Team = teamUid
	r.Params.Uid = uid
	return r
}

type ServerCallBuzzcancel struct {
	BaseEvent
	Params ServerCallBuzzcancelParams `json:"params"`
}

type ServerCallBuzzcancelParams struct {
	Jid  JID    `json:"jid"`
	Team string `json:"team"`
	Uid  string `json:"uid"`
}
