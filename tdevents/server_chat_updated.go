package tdevents

import "github.com/tada-team/tdproto"

func NewServerChatUpdated(chats []tdproto.Chat, teamUnread *tdproto.TeamUnread, badge uint) (r ServerChatUpdated) {
	r.Name = r.GetName()
	r.Params.Chats = chats
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
	Chats []tdproto.Chat `json:"chats"`

	// Current team counters
	TeamUnread *tdproto.TeamUnread `json:"team_unread"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
