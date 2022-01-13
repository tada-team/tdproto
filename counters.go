package tdproto

// Unread messages counter
type ChatCounters struct {
	LastReadMessageUid *string           `json:"last_read_message_id"`
	Jid                JID               `json:"jid"`
	ChatType           ChatType          `json:"chat_type"`
	LastActivity       ISODateTimeString `json:"last_activity,omitempty"`
	Gentime            int64             `json:"gentime"`
	NumUnread          uint              `json:"num_unread"`
	NumUnreadNotices   uint              `json:"num_unread_notices"`
}

// Unread message counters
type Unread struct {
	// Total unread messages
	NumMessages uint `json:"messages"`

	// Total unread messages with mentions
	NumNoticeMessages uint `json:"notice_messages"`

	// Total chats with unread messages
	NumChats uint `json:"chats"`
}

type TeamUnread map[ChatType]*Unread

// Unread message counters
type TeamCounter struct {
	// Unread message counters
	Unreads TeamUnread `json:"unread"`

	// Team id
	Uid string `json:"uid"`
}

func EmptyTeamUnread() TeamUnread {
	return TeamUnread{
		DirectChatType: new(Unread),
		GroupChatType:  new(Unread),
		TaskChatType:   new(Unread),
	}
}

func (unread TeamUnread) Empty() bool {
	for _, v := range unread {
		if v.NumMessages > 0 {
			return false
		}
	}
	return true
}
