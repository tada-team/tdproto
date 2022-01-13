package tdproto

func NewServerChatLastread(counters ChatCounters, teamUnread *TeamUnread, badge uint) (r ServerChatLastread) {
	r.Name = r.GetName()
	r.Params.Chats = []ChatCounters{counters}
	r.Params.TeamUnread = teamUnread
	r.Params.Badge = badge
	return r
}

// Changed last read message in chat
type ServerChatLastread struct {
	BaseEvent
	Params serverChatLastreadParams `json:"params"`
}

func (p ServerChatLastread) GetName() string { return "server.chat.lastread" }

// Params of the server.chat.lastread event
type serverChatLastreadParams struct {
	// Current team counters
	TeamUnread *TeamUnread `json:"team_unread"`

	// Chat counters
	Chats []ChatCounters `json:"chats"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
