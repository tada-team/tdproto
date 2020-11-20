package tdproto

type (
	TaskFilterKey string
	TaskSortKey   string
	TaskTabKey    string
)

func (key TaskFilterKey) String() string {
	return string(key)
}

type TaskSort struct {
	Key   TaskSortKey `json:"key"`
	Title string      `json:"title"`
}

type TaskFilter struct {
	Field TaskFilterKey `json:"field"`
	Title string        `json:"title"`
}

// Tasks counters
type TaskCounters struct {
	// task jid
	Jid JID `json:"jid"`

	// unread counter
	NumUnread uint `json:"num_unread"`

	//unread notice counter
	NumUnreadNotice uint `json:"num_unread_notice"`
}

type TaskTab struct {
	Key         TaskTabKey   `json:"key"`
	Title       string       `json:"title"`
	HideEmpty   bool         `json:"hide_empty"`
	ShowCounter bool         `json:"show_counter"`
	Pagination  bool         `json:"pagination"`
	Filters     []TaskFilter `json:"filters"`
	Sort        []TaskSort   `json:"sort"`

	// unread tasks with jid and counters
	UnreadTasks []TaskCounters `json:"unread_tasks"`
}
