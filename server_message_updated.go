package tdproto

func NewServerMessageUpdated(messages []Message, delayed bool, counters *ChatCounters, teamUnread *TeamUnread, badge *uint) (r ServerMessageUpdated) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	r.Params.Messages = messages
	r.Params.Delayed = delayed

	r.Params.ChatCounters = make([]ChatCounters, 0)
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

type serverMessageUpdatedParams struct {
	Messages     []Message      `json:"messages"`
	Delayed      bool           `json:"delayed"`
	ChatCounters []ChatCounters `json:"chat_counters"` // TODO: omitempty
	TeamUnread   *TeamUnread    `json:"team_unread"`   // TODO: omitempty
	Badge        *uint          `json:"badge"`         // TODO: omitempty
}
