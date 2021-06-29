package tdevents

import "github.com/tada-team/tdproto"

func NewServerChatDeleted(chat tdproto.DeletedChat, teamUnread *tdproto.TeamUnread, badge uint) (r ServerChatDeleted) {
	chat.IsArchive = true
	chat.Gentime = tdproto.Gentime()

	r.Name = r.GetName()
	r.Params.Chats = []tdproto.DeletedChat{chat}
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
	// List of deleted chats
	Chats []tdproto.DeletedChat `json:"chats"`

	// Current team counters
	TeamUnread *tdproto.TeamUnread `json:"team_unread"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
