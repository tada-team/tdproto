package tdproto

func NewServerChatUpdated(chat Chat, teamUnread *TeamUnread, badge uint) (r ServerChatUpdated) {
	r.Name = r.GetName()
	r.Params.Chats = []Chat{chat}
	r.Params.TeamUnread = teamUnread
	r.Params.Badge = badge
	return r
}

// Chat created or updated
type ServerChatUpdated struct {
	BaseEvent
	Params serverChatUpdatedParams `json:"params"`
}

func (p ServerChatUpdated) GetName() string { return "server.chat.updated" }

// Params of the server.chat.updated event
type serverChatUpdatedParams struct {
	// Chat counters
	Chats []Chat `json:"chats"`

	// Current team counters
	TeamUnread *TeamUnread `json:"team_unread"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
