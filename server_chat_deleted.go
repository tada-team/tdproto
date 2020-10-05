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

type ServerChatDeleted struct {
	BaseEvent
	Params serverChatDeletedParams `json:"params"`
}

func (p ServerChatDeleted) GetName() string { return "server.chat.deleted" }

type serverChatDeletedParams struct {
	Chats      []DeletedChat `json:"chats"`
	TeamUnread *TeamUnread   `json:"team_unread"`
	Badge      uint          `json:"badge"`
}
