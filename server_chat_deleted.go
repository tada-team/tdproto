package tdproto

func NewServerChatDeleted(chat DeletedChat, teamUnread *TeamUnread, badge uint) (r ServerChatDeleted) {
	chat.IsArchive = true
	chat.Gentime = Gentime()

	r.Name = r.GetName()
	r.Params.Chats = []DeletedChat{chat}
	r.Params.TeamUnread = teamUnread
	r.Params.Badge = badge
	return r
}

// Chat deleted
type ServerChatDeleted struct {
	BaseEvent
	Params serverChatDeletedParams `json:"params"`
}

func (p ServerChatDeleted) GetName() string { return "server.chat.deleted" }

// Params of the server.chat.deleted event
type serverChatDeletedParams struct {
	// Current team counters
	TeamUnread *TeamUnread `json:"team_unread"`

	// List of deleted chats
	Chats []DeletedChat `json:"chats"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
