package tdproto

type ChatCounters struct {
	Jid                JID      `json:"jid"`
	ChatType           ChatType `json:"chat_type"`
	Gentime            int64    `json:"gentime"`
	NumUnread          uint     `json:"num_unread"`
	NumUnreadNotices   uint     `json:"num_unread_notices"`
	LastReadMessageUid *string  `json:"last_read_message_id"`
}

type Unread struct {
	NumMessages       uint `json:"messages"`
	NumNoticeMessages uint `json:"notice_messages"`
	NumChats          uint `json:"chats"`
}

type TeamUnread map[ChatType]*Unread

type TeamCounter struct {
	Uid     string     `json:"uid"`
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
