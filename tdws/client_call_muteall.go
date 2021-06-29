package tdws

import "github.com/tada-team/tdproto"

// Mute all other call participants
type ClientCallMuteAll struct {
	BaseEvent
	Params clientCallMuteAllParams `json:"params"`
}

func (p ClientCallMuteAll) GetName() string { return "client.call.muteall" }

// Params of the client.call.muteall event
type clientCallMuteAllParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`
}
