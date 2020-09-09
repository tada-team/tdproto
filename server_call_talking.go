package tdproto

func NewServerCallTalking(talking bool, chat, actor *JID) (r ServerCallTalking) {
	r.BaseEvent.Unimportant = true
	r.BaseEvent.Name = "server.call.talking"
	r.Params.Jid = *chat
	r.Params.Actor = *actor
	r.Params.Talking = talking
	return r
}

type ServerCallTalking struct {
	BaseEvent
	Params serverCallTalkingParams `json:"params"`
}

type serverCallTalkingParams struct {
	Jid     JID  `json:"jid"`
	Actor   JID  `json:"actor"`
	Talking bool `json:"talking"`
}
