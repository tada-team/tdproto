package tdproto

func NewServerChatDeleted(chat DeletedChat, teamUnread *TeamUnread, badge uint) (r ServerChatDeleted) {
	chat.IsArchive = true
	chat.Gentime = Gentime()

	r.BaseEvent.Name = "server.chat.deleted"
	r.Params.Chats = []DeletedChat{chat}
	r.Params.TeamUnread = teamUnread
	r.Params.Badge = badge
	return r
}

type ServerChatDeleted struct {
	BaseEvent
	Params ServerChatDeletedParams `json:"params"`
}

type ServerChatDeletedParams struct {
	Chats      []DeletedChat `json:"chats"`
	TeamUnread *TeamUnread   `json:"team_unread"`
	Badge      uint          `json:"badge"`
}
