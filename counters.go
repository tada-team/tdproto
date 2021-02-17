package tdproto

type ChatCounters struct {
	Jid                JID               `json:"jid"`
	ChatType           ChatType          `json:"chat_type"`
	Gentime            int64             `json:"gentime"`
	NumUnread          uint              `json:"num_unread"`
	NumUnreadNotices   uint              `json:"num_unread_notices"`
	LastReadMessageUid *string           `json:"last_read_message_id"`
	LastActivity       ISODateTimeString `json:"last_activity,omitempty"`
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
	// Team id
	Uid string `json:"uid"`

	// Unread message counters
	Unreads TeamUnread `json:"unread"`
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
