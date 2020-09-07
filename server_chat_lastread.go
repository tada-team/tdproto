package tdproto

func NewServerChatLastread(counters ChatCounters, teamUnread *TeamUnread, badge uint) (r ServerChatLastread) {
	r.BaseEvent.Name = "server.chat.lastread"
	r.Params.Chats = []ChatCounters{counters}
	r.Params.TeamUnread = teamUnread
	r.Params.Badge = badge
	return r
}

type ServerChatLastread struct {
	BaseEvent
	Params ServerChatLastreadParams `json:"params"`
}

type ServerChatLastreadParams struct {
	Chats      []ChatCounters `json:"chats"`
	TeamUnread *TeamUnread    `json:"team_unread"`
	Badge      uint           `json:"badge"`
}
