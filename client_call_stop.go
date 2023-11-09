package tdproto

// Stop call
type ClientCallStop struct {
	BaseEvent
	Params clientCallStopParams `json:"params"`
}

func (p ClientCallStop) GetName() string { return "client.call.stop" }

// Params of the client.call.stop event
type clientCallStopParams struct {
	// Chat id
	Jid JID `json:"jid"`
}
