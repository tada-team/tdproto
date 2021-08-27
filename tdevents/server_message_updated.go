package tdevents

import "github.com/tada-team/tdproto"

func NewServerMessageUpdated(messages []tdproto.Message, delayed bool, counters *tdproto.ChatCounters, teamUnread *tdproto.TeamUnread, badge *uint) (r ServerMessageUpdated) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	r.Params.Messages = messages
	r.Params.Delayed = delayed

	r.Params.ChatCounters = make([]tdproto.ChatCounters, 0)
	if counters != nil {
		r.Params.ChatCounters = append(r.Params.ChatCounters, *counters)
	}

	if teamUnread != nil {
		if badge == nil {
			panic("programming error: empty badge")
		}
		r.Params.TeamUnread = teamUnread
		r.Params.Badge = badge
	}

	return r
}

// Chat message created, updated or deleted
type ServerMessageUpdated struct {
	BaseEvent
	Params serverMessageUpdatedParams `json:"params"`
}

func (p ServerMessageUpdated) GetName() string { return "server.message.updated" }

// Params of the server.message.updated event
type serverMessageUpdatedParams struct {
	// Messages data
	Messages []tdproto.Message `json:"messages"`

	// true = silently message update, false = new message
	Delayed bool `json:"delayed"`

	// Chat counters
	ChatCounters []tdproto.ChatCounters `json:"chat_counters"`

	// Current team counters
	TeamUnread *tdproto.TeamUnread `json:"team_unread"`

	// Total number of unreads, if changed
	Badge *uint `json:"badge"`
}
