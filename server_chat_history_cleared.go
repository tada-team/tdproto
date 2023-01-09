package tdproto

func NewServerChatHistoryCleared(jid JID, lastCleared *string) (r ServerChatHistoryCleared) {
	r.Name = r.GetName()
	r.Params.JID = jid
	if lastCleared != nil {
		r.Params.LastCleared = lastCleared
	}
	return r
}

// ServerChatHistoryCleared
type ServerChatHistoryCleared struct {
	BaseEvent
	Params serverChatHistoryClearedParams `json:"params"`
}

func (p ServerChatHistoryCleared) GetName() string { return "server.chat.history.cleared" }

// Params of the server.chat.history.cleared event
type serverChatHistoryClearedParams struct {
	// chat jid
	JID JID `json:"jid"`

	LastCleared *string `json:"last_cleared,omitempty"`
}
