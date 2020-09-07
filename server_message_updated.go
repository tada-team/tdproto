package tdproto

import (
	"log"
)

func NewServerMessageUpdated(messages []Message, delayed bool, counters *ChatCounters, teamUnread *TeamUnread, badge *uint) (r ServerMessageUpdated) {
	r.Name = "server.message.updated"
	r.ConfirmId = ConfirmId()
	r.Params.Messages = messages
	r.Params.Delayed = delayed

	r.Params.ChatCounters = make([]ChatCounters, 0)
	if counters != nil {
		r.Params.ChatCounters = append(r.Params.ChatCounters, *counters)
	}

	if teamUnread != nil {
		if badge == nil {
			log.Panicln("empty badge")
		}
		r.Params.TeamUnread = teamUnread
		r.Params.Badge = badge
	}

	return r
}

type ServerMessageUpdated struct {
	BaseEvent
	Params ServerMessageUpdatedParams `json:"params"`
}

type ServerMessageUpdatedParams struct {
	Messages     []Message      `json:"messages"`
	Delayed      bool           `json:"delayed"`
	ChatCounters []ChatCounters `json:"chat_counters"` // TODO: omitempty
	TeamUnread   *TeamUnread    `json:"team_unread"`   // TODO: omitempty
	Badge        *uint          `json:"badge"`         // TODO: omitempty
}
