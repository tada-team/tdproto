package tdproto

type DeletedChat struct {
	Jid       JID      `json:"jid"`
	ChatType  ChatType `json:"chat_type"`
	Gentime   int64    `json:"gentime"`
	IsArchive bool     `json:"is_archive"`
}

type Chat struct {
	ChatType                  ChatType          `json:"chat_type"`
	BaseGentime               int64             `json:"base_gentime,omitempty"`
	Gentime                   int64             `json:"gentime"`
	Created                   string            `json:"created"`
	Jid                       JID               `json:"jid"`
	DisplayName               string            `json:"display_name"`
	Icons                     *IconData         `json:"icons"`
	CountersEnabled           bool              `json:"counters_enabled,omitempty"`
	CanCall                   bool              `json:"can_call,omitempty"`
	CanSendMessage            bool              `json:"can_send_message,omitempty"`
	CantSendMessageReason     string            `json:"cant_send_message_reason,omitempty"`
	Collapsed                 bool              `json:"collapsed,omitempty"`
	Draft                     string            `json:"draft,omitempty"`
	DraftNum                  int64             `json:"draft_num,omitempty"`
	Hidden                    bool              `json:"hidden,omitempty"`
	NotificationsEnabled      bool              `json:"notifications_enabled,omitempty"`
	NumImportants             int               `json:"num_importants,omitempty"`
	NumUnread                 uint              `json:"num_unread,omitempty"`
	NumUnreadNotices          uint              `json:"num_unread_notices,omitempty"`
	LastMessage               *Message          `json:"last_message,omitempty"`
	LastReadMessageId         string            `json:"last_read_message_id,omitempty"`
	Section                   string            `json:"section,omitempty"`
	ChangeableFields          []string          `json:"changeable_fields,omitempty"`
	Pinned                    bool              `json:"pinned,omitempty"`
	PinnedSortOrdering        int               `json:"pinned_sort_ordering,omitempty"`
	NumMembers                *uint             `json:"num_members,omitempty"`
	CanDelete                 bool              `json:"can_delete,omitempty"`
	Description               string            `json:"description,omitempty"`
	Feed                      bool              `json:"feed,omitempty"`
	PinnedMessage             *Message          `json:"pinned_message,omitempty"`
	ColorIndex                *uint16           `chattype:"task" json:"color_index,omitempty"`
	NumItems                  *uint             `chattype:"task" json:"num_items,omitempty"`
	NumCheckedItems           *uint             `chattype:"task" json:"num_checked_items,omitempty"`
	Assignee                  *JID              `chattype:"task" json:"assignee,omitempty"`
	Num                       uint              `chattype:"task" json:"num,omitempty"`
	Observers                 *[]JID            `chattype:"task" json:"observers,omitempty"`
	Owner                     *JID              `chattype:"task" json:"owner,omitempty"`
	TaskStatus                string            `chattype:"task" json:"task_status,omitempty"`
	Title                     string            `chattype:"task" json:"title,omitempty"`
	Done                      string            `chattype:"task" json:"done,omitempty"`
	DoneReason                string            `chattype:"task" json:"done_reason,omitempty"`
	Deadline                  string            `chattype:"task" json:"deadline,omitempty"`
	DeadlineExpired           bool              `chattype:"task" json:"deadline_expired,omitempty"`
	Links                     MessageLinks      `chattype:"task" json:"links,omitempty"`
	Tags                      []string          `chattype:"task" json:"tags,omitempty"`
	Importance                *int              `chattype:"task" json:"importance,omitempty"`
	Urgency                   *int              `chattype:"task" json:"urgency,omitempty"`
	SpentTime                 *int              `chattype:"task" json:"spent_time,omitempty"`
	Complexity                *int              `chattype:"task" json:"complexity,omitempty"`
	LinkedMessages            []interface{}     `chattype:"task" json:"linked_messages,omitempty"`
	Items                     []TaskItem        `chattype:"task" json:"items,omitempty"`
	Parents                   []Subtask         `chattype:"task" json:"parents,omitempty"`
	Tabs                      *[]TaskTabKey     `chattype:"task" json:"tabs,omitempty"`
	Status                    GroupStatus       `chattype:"group" json:"status,omitempty"`
	Members                   []GroupMembership `chattype:"group" json:"members,omitempty"`
	CanAddMember              bool              `chattype:"group" json:"can_add_member,omitempty"`
	CanRemoveMember           bool              `chattype:"group" json:"can_remove_member,omitempty"`
	CanChangeMemberStatus     bool              `chattype:"group" json:"can_change_member_status,omitempty"`
	CanChangeSettings         bool              `chattype:"group" json:"can_change_settings,omitempty"` // deprecated: use changeable fields
	DefaultForAll             bool              `chattype:"group" json:"default_for_all,omitempty"`
	ReadonlyForMembers        bool              `chattype:"group" json:"readonly_for_members,omitempty"`
	AutocleanupAge            *int              `chattype:"group" json:"autocleanup_age,omitempty"`
	Public                    bool              `chattype:"group,task" json:"public,omitempty"`
	CanJoin                   bool              `chattype:"group,task" json:"can_join,omitempty"`
	CanDeleteAnyMessage       *bool             `json:"can_delete_any_message,omitempty"`
	CanSetImportantAnyMessage *bool             `json:"can_set_important_any_message,omitempty"`
	//WikiPage              *ShortWikiPage `chattype:"group" json:"wiki_page,omitempty"`
}

type ChatShort struct {
	Jid         JID       `json:"jid"`
	DisplayName string    `json:"display_name"`
	Icons       *IconData `json:"icons"`
	ChatType    ChatType  `json:"chat_type"`
}

type Subtask struct {
	Jid         JID    `json:"jid"`
	Assignee    JID    `json:"assignee"`
	Title       string `json:"title"`
	Num         uint   `json:"num"`
	DisplayName string `json:"display_name"`
	Public      bool   `json:"public,omitempty"`
}

type TaskItem struct {
	Uid          string   `json:"uid,omitempty"`
	SortOrdering uint     `json:"sort_ordering,omitempty"`
	Text         string   `json:"text"`
	Checked      bool     `json:"checked,omitempty"`
	CanToggle    bool     `json:"can_toggle,omitempty"`
	Subtask      *Subtask `json:"subtask,omitempty"`
}

type GroupMembership struct {
	Jid       *JID        `json:"jid"`
	Status    GroupStatus `json:"status"`
	CanRemove bool        `json:"can_remove,omitempty"`
}
