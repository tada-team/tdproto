package tdproto

func NewServerCallBuzzcancel(chat JID, teamUid string, uid string) (r ServerCallBuzzcancel) {
	r.Name = r.GetName()
	r.Params.Jid = chat
	r.Params.Team = teamUid
	r.Params.Uid = uid
	return r
}

// Call cancelled on buzzing
type ServerCallBuzzcancel struct {
	BaseEvent
	Params serverCallBuzzcancelParams `json:"params"`
}

func (p ServerCallBuzzcancel) GetName() string { return "server.call.buzzcancel" }

type serverCallBuzzcancelParams struct {
	// Chat or contact id
	Jid  JID    `json:"jid"`

	// Team id
	Team string `json:"team"`

	// Call id
	Uid  string `json:"uid"`
}
