package tdproto

func NewServerCallTalking(talking bool, chat, actor *JID) (r ServerCallTalking) {
	r.Name = r.GetName()
	r.Params.Jid = *chat
	r.Params.Actor = *actor
	r.Params.Talking = talking
	return r
}

// Someone talks in call
type ServerCallTalking struct {
	BaseEvent
	Params serverCallTalkingParams `json:"params"`
}

func (p ServerCallTalking) GetName() string { return "server.call.talking" }

type serverCallTalkingParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Actor id
	Actor JID `json:"actor"`

	// Is talking
	Talking bool `json:"talking"`
}
