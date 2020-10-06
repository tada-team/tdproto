package tdproto

func NewServerCallTalking(talking bool, chat, actor *JID) (r ServerCallTalking) {
	r.Name = r.GetName()
	r.Params.Jid = *chat
	r.Params.Actor = *actor
	r.Params.Talking = talking
	return r
}

type ServerCallTalking struct {
	BaseEvent
	Params serverCallTalkingParams `json:"params"`
}

func (p ServerCallTalking) GetName() string { return "server.call.talking" }

type serverCallTalkingParams struct {
	Jid     JID  `json:"jid"`
	Actor   JID  `json:"actor"`
	Talking bool `json:"talking"`
}
