package tdproto

func NewServerChatUpdated(chat Chat, teamUnread *TeamUnread, badge uint) (r ServerChatUpdated) {
	r.BaseEvent.Name = "server.chat.updated"
	r.Params.Chats = []Chat{chat}
	r.Params.TeamUnread = teamUnread
	r.Params.Badge = badge
	return r
}

type ServerChatUpdated struct {
	BaseEvent
	Params serverChatUpdatedParams `json:"params"`
}

type serverChatUpdatedParams struct {
	Chats      []Chat      `json:"chats"`
	TeamUnread *TeamUnread `json:"team_unread"`
	Badge      uint        `json:"badge"`
}
