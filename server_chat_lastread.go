package tdproto

func NewServerChatLastread(counters ChatCounters, teamUnread *TeamUnread, badge uint) (r ServerChatLastread) {
	r.Name = r.GetName()
	r.Params.Chats = []ChatCounters{counters}
	r.Params.TeamUnread = teamUnread
	r.Params.Badge = badge
	return r
}

type ServerChatLastread struct {
	BaseEvent
	Params serverChatLastreadParams `json:"params"`
}

func (p ServerChatLastread) GetName() string { return "server.chat.lastread" }

type serverChatLastreadParams struct {
	Chats      []ChatCounters `json:"chats"`
	TeamUnread *TeamUnread    `json:"team_unread"`
	Badge      uint           `json:"badge"`
}
