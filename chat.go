package tdproto

// Minimal chat representation
type ChatShort struct {
	// Group/Task/Contact id
	Jid JID `json:"jid"`

	// Chat type
	ChatType ChatType `json:"chat_type"`

	// Title
	DisplayName string `json:"display_name"`

	// Icon data
	Icons IconData `json:"icons"`
}

// Minimal chat representation for deletion
type DeletedChat struct {
	// Group/Task/Contact id
	Jid JID `json:"jid"`

	// Chat type
	ChatType ChatType `json:"chat_type"`

	// Chat fields (related to concrete participant) version
	Gentime int64 `json:"gentime"`

	// Archive flag. Always true for this structure
	IsArchive bool `json:"is_archive"`
}

// Chat (direct, group, task) representation
type Chat struct {
	// Task urgency, if available in team
	Urgency *int `chattype:"task" json:"urgency,omitempty"`

	// Checked items in checklist. Tasks only
	NumCheckedItems *uint `chattype:"task" json:"num_checked_items,omitempty"`

	// Task importance, if available in team
	Importance *int `chattype:"task" json:"importance,omitempty"`

	// Task complexity, number
	Complexity *int `chattype:"task" json:"complexity,omitempty"`

	// Non-archive participants number
	NumMembers *uint `json:"num_members,omitempty"`

	// Last message object
	LastMessage *Message `json:"last_message,omitempty"`

	// Task spent time, number
	SpentTime *int `chattype:"task" json:"spent_time,omitempty"`

	// Pinned message for this chat
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// My status in group chat
	Status *GroupStatus `chattype:"group" json:"status,omitempty"`

	// Delete messages in this chat in seconds. Experimental function
	AutocleanupAge *int `chattype:"group" json:"autocleanup_age,omitempty"`

	// Custom color index from table of colors. Tasks only
	ColorIndex *uint16 `chattype:"task" json:"color_index,omitempty"`

	// Items in checklist. Tasks only
	NumItems *uint `chattype:"task" json:"num_items,omitempty"`

	// Task done reason, if any
	DoneReason string `chattype:"task" json:"done_reason,omitempty"`

	// Last message draft, if any
	Draft string `json:"draft,omitempty"`

	// Why I can't send message to this chat (if can't)
	CantSendMessageReason string `json:"cant_send_message_reason,omitempty"`

	// Group/Task/Contact id
	Jid JID `json:"jid"`

	// Task title. Generated from number and description
	Title string `chattype:"task" json:"title,omitempty"`

	// Task status. May be custom
	TaskStatus string `chattype:"task" json:"task_status,omitempty"`

	// Date of the last message sent even if it was deleted
	LastActivity ISODateTimeString `json:"last_activity,omitempty"`

	// Title
	DisplayName string `json:"display_name"`

	// Last read message id, if any
	LastReadMessageId string `json:"last_read_message_id,omitempty"`

	// Project / section id, if any
	Section string `json:"section,omitempty"`

	// Task creator
	Owner JID `chattype:"task" json:"owner,omitempty"`

	// Assignee contact id. Tasks only
	Assignee JID `chattype:"task" json:"assignee,omitempty"`

	// Task deadline in iso format, if any
	Deadline ISODateTimeString `chattype:"task" json:"deadline,omitempty"`

	// Creation date, iso datetime
	Created ISODateTimeString `json:"created"`

	// Chat type
	ChatType ChatType `json:"chat_type"`

	// Group or task description
	Description string `json:"description,omitempty"`

	// Task done date in iso format, if any
	Done ISODateTimeString `chattype:"task" json:"done,omitempty"`

	// Tab names
	Tabs []TaskTabKey `chattype:"task" json:"tabs,omitempty"`

	// Parent tasks
	Parents []Subtask `chattype:"task" json:"parents,omitempty"`

	// Used for "Create task from messages..."
	LinkedMessages []interface{} `chattype:"task" json:"linked_messages,omitempty"`

	// Markup entities for description field. Experimental
	Markup []MarkupEntity `json:"markup,omitempty" tdproto:"readonly"`

	// Upload uids for request, upload objects for response
	Uploads []Upload `chattype:"task" json:"uploads,omitempty"`

	// Checklist items. Task only
	Items []TaskItem `chattype:"task" json:"items,omitempty"`

	// Group chat members
	Members []GroupMembership `chattype:"group" json:"members,omitempty"`

	// Task followers id's. TODO: rename to "followers"
	Observers []JID `chattype:"task" json:"observers,omitempty"`

	// List of editable fields
	ChangeableFields []string `json:"changeable_fields,omitempty"`

	// Links in description
	Links MessageLinks `chattype:"task" json:"links,omitempty"`

	// Task tags list, if any
	Tags []string `chattype:"task" json:"tags,omitempty"`

	// Icons info
	Icons IconData `json:"icons"`

	// Sort ordering for pinned chat
	PinnedSortOrdering int `json:"pinned_sort_ordering,omitempty"`

	// Last message draft version, if any
	DraftGentime int64 `json:"draft_gentime,omitempty"`

	// Chat fields related to concrete participant) version
	Gentime int64 `json:"gentime"`

	// Unread counter
	NumUnread uint `json:"num_unread,omitempty"`

	// Task number in this team
	Num uint `chattype:"task" json:"num,omitempty"`

	// Base fields (not related to concrete participant) version
	BaseGentime int64 `json:"base_gentime,omitempty"`

	// Mentions (@) counter
	NumUnreadNotices uint `json:"num_unread_notices,omitempty"`

	// Number of important messages
	NumImportants int `json:"num_importants,omitempty"`

	// Deprecated
	DraftNum int64 `json:"draft_num,omitempty"`

	// Push notifications enabled
	NotificationsEnabled bool `json:"notifications_enabled,omitempty"`

	// Can I delete this chat
	CanDelete bool `json:"can_delete,omitempty"`

	// Is chat pinned on top
	Pinned bool `json:"pinned,omitempty"`

	// Any new team member will be added to this group chat
	DefaultForAll bool `chattype:"group" json:"default_for_all,omitempty"`

	// Can I call to this chat
	CanCall bool `json:"can_call,omitempty"`

	// Can I change member status in this group chat
	CanChangeMemberStatus bool `chattype:"group" json:"can_change_member_status,omitempty"`

	// Hidden chat
	Hidden bool `json:"hidden,omitempty"`

	// Can I add member to this group chat
	CanAddMember bool `chattype:"group" json:"can_add_member,omitempty"`

	// Can I remove member from this group chat
	CanRemoveMember bool `chattype:"group" json:"can_remove_member,omitempty"`

	// Present in feed (main screen)
	Feed bool `json:"feed,omitempty"`

	// deprecated: use changeable fields
	CanChangeSettings bool `chattype:"group" json:"can_change_settings,omitempty"`

	// Description collapsed. Used for tasks only
	Collapsed bool `json:"collapsed,omitempty"`

	// Readonly for non-admins group chat (Like Channels in Telegram but switchable)
	ReadonlyForMembers bool `chattype:"group" json:"readonly_for_members,omitempty"`

	// Can I send message to this chat
	CanSendMessage bool `json:"can_send_message,omitempty"`

	// Can other team member see this task/group chat
	Public bool `chattype:"group,task" json:"public,omitempty"`

	// Can I join to this public group/task
	CanJoin bool `chattype:"group,task" json:"can_join,omitempty"`

	// Can I delete any message in this chat
	CanDeleteAnyMessage bool `json:"can_delete_any_message,omitempty"`

	// Can I change Important flag in any message in this chat
	CanSetImportantAnyMessage bool `json:"can_set_important_any_message,omitempty"`

	// Include unread messages to counters
	CountersEnabled bool `json:"counters_enabled,omitempty"`

	// Is task deadline expired
	DeadlineExpired bool `chattype:"task" json:"deadline_expired,omitempty"`
}

// Link to sub/sup task
type Subtask struct {
	// Task id
	Jid JID `json:"jid"`

	// Assignee contact id. Tasks only
	Assignee JID `json:"assignee"`

	// Task title. Generated from number and description
	Title string `json:"title"`

	// Title
	DisplayName string `json:"display_name"`

	// Subtask task status
	TaskStatus string `json:"task_status,omitempty"`

	// Task number in this team
	Num uint `json:"num"`

	// Is task or group public for non-guests
	Public bool `json:"public,omitempty"`
}

// Task checklist item
type TaskItem struct {
	// Link to subtask. Optional
	Subtask *Subtask `json:"subtask,omitempty"`

	// Is task or group public for non-guests
	Uid string `json:"uid,omitempty"`

	// Text or "#{OtherTaskNumber}"
	Text string `json:"text"`

	// Sort ordering
	SortOrdering uint `json:"sort_ordering,omitempty"`

	// Object version
	Gentime int64 `json:"gentime" tdproto:"readonly"`

	// Item checked
	Checked bool `json:"checked,omitempty"`

	// Can I toggle this item
	CanToggle bool `json:"can_toggle,omitempty"`

	// Can I change this item
	CanChange bool `json:"can_change,omitempty"`
}

// Group chat membership status
type GroupMembership struct {
	// Contact id
	Jid JID `json:"jid"`

	// Status in group
	Status GroupStatus `json:"status,omitempty"`

	// Can I remove this member
	CanRemove bool `json:"can_remove,omitempty"`
}
