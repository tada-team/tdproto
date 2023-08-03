package tdproto

// NewServerChatHistoryCleared returns the new ServerChatHistoryCleared instance.
func NewServerChatHistoryCleared(jid JID, lastCleared *string) (r ServerChatHistoryCleared) {
	r.Name = r.GetName()
	r.Params.JID = jid
	if lastCleared != nil {
		r.Params.LastClear = lastCleared
	}

	return r
}

// ServerChatHistoryCleared represents the event about clearing the chat messages history for user.
type ServerChatHistoryCleared struct {
	BaseEvent
	Params serverChatHistoryClearedParams `json:"params"`
}

// GetName returns the name of ServerChatHistoryCleared instance.
func (p ServerChatHistoryCleared) GetName() string { return "server.chat.history.cleared" }

// Params of the server.chat.history.cleared event
type serverChatHistoryClearedParams struct {
	// Chat jid
	JID JID `json:"jid"`

	// LastClear last clear chat history
	LastClear *string `json:"last_clear"`
}
