package tdproto

type (
	TaskFilterKey string
	TaskSortKey   string
	TaskTabKey    string
)

func (key TaskFilterKey) String() string {
	return string(key)
}

// Task sort type
type TaskSort struct {
	// Field
	Key TaskSortKey `json:"key"`

	// Sort title
	Title string `json:"title"`
}

// Task filter
type TaskFilter struct {
	// Task filter field
	Field TaskFilterKey `json:"field"`

	// Filter title
	Title string `json:"title"`
}

// Tasks counters
type TaskCounters struct {
	// Task jid
	Jid JID `json:"jid"`

	// Unreads conuter
	NumUnread uint `json:"num_unread,omitempty"`

	// Mentions (@) counter
	NumUnreadNotices uint `json:"num_unread_notices,omitempty"`
}

// Task tab
type TaskTab struct {
	// Tab name
	Key TaskTabKey `json:"key"`

	// Tab title
	Title string `json:"title"`

	// Disable this tab when it has no contents
	HideEmpty bool `json:"hide_empty"`

	// Show unread badge
	ShowCounter bool `json:"show_counter"`

	// Enable pagination
	Pagination bool `json:"pagination"`

	// Filters inside tab
	Filters []TaskFilter `json:"filters"`

	// Sort available in tab
	Sort []TaskSort `json:"sort"`

	// Unread tasks with jid and counters
	UnreadTasks []TaskCounters `json:"unread_tasks"`
}
