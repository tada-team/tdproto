package tdproto

// NewServerChatHistoryCleared returns the new ServerChatHistoryCleared instance.
func NewServerChatHistoryCleared(jid JID, lastCleared *string) (r ServerChatHistoryCleared) {
	r.Name = r.GetName()
	r.Params.JID = jid
	if lastCleared != nil {
		r.Params.LastCleared = lastCleared
	}

	return r
}

// ServerChatHistoryCleared represents the event about clearing the chat messages history for user.
type ServerChatHistoryCleared struct {
	BaseEvent
	Params serverChatMessagesHistoryClearedParams `json:"params"`
}

// GetName returns the name of ServerChatHistoryCleared instance.
func (p ServerChatHistoryCleared) GetName() string { return "server.chat.history.cleared" }

// Params of the server.chat.history.cleared event
type serverChatMessagesHistoryClearedParams struct {
	// Chat jid
	JID JID `json:"jid"`

	// Last clear chat history
	LastCleared *string `json:"last_cleared,omitempty"`
}
