package tdws

import "github.com/tada-team/tdproto"

func NewServerChatLastread(counters tdproto.ChatCounters, teamUnread *tdproto.TeamUnread, badge uint) (r ServerChatLastread) {
	r.Name = r.GetName()
	r.Params.Chats = []tdproto.ChatCounters{counters}
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
	// Chat counters
	Chats []tdproto.ChatCounters `json:"chats"`

	// Current team counters
	TeamUnread *tdproto.TeamUnread `json:"team_unread"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
