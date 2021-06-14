package tdapi

import "github.com/tada-team/tdproto"

// Task
type Task struct {
	// Custom task color
	CustomColorIndex *uint16 `json:"custom_color_index,omitempty"`

	// Task description
	Description string `json:"description,omitempty"`

	// Task tags
	Tags []string `json:"tags,omitempty"`

	// Task section UID
	SectionUid string `json:"section,omitempty"`

	// User who follow the task
	Observers []tdproto.JID `json:"observers,omitempty"` // TODO: rename to "followers"

	// Items of the task
	Items []string `json:"items,omitempty"`

	// User who was assigned the task
	Assignee tdproto.JID `json:"assignee,omitempty"`

	// Deadline time, if any
	Deadline tdproto.ISODateTimeString `json:"deadline,omitempty"`

	// Is task public
	Public bool `json:"public,omitempty"`

	// Fire a reminder at this time
	RemindAt tdproto.ISODateTimeString `json:"remind_at,omitempty"`

	// Task status
	TaskStatus string `json:"task_status,omitempty"`

	// Task importance
	Importance *int `json:"importance,omitempty"`

	// Task urgency
	Urgency *int `json:"urgency,omitempty"`

	// Task complexity
	Complexity *int `json:"complexity,omitempty"`

	// Time spent
	SpentTime *int `json:"spent_time,omitempty"`

	// Linked messages
	LinkedMessages []string `json:"linked_messages,omitempty"` // TODO: Message object

	// Task uploads
	Uploads []string `json:"uploads,omitempty"`
}

type TaskFilter struct {
	UserParams
	Paginator

	//* ?sort = [ "created" | "-created" | "last_message" | "-last_message" | "deadline" | "-deadline" ]
	Sort string `schema:"sort"`

	//* ?task_status = new,done | new | any
	TaskStatus string `schema:"task_status"`

	//* ?exclude_task_status = new,done | new | any
	ExcludeTaskStatus string `schema:"exclude_task_status"`

	//* ?num=num1,num2,num3...
	Num string `schema:"num"`

	//* ?observer=jid,jid   // TODO: rename to ?follower=
	Observer string `schema:"observer"`

	//* ?member=jid,jid
	Member string `schema:"member"`

	//* ?assignee=jid,jid
	Assignee string `schema:"assignee"`

	//* ?owner=jid,jid
	Owner string `schema:"owner"`

	//* ?section=[ uid,uid... | "-" ]
	Section string `schema:"section"` // TODO: rename to ?project=

	//* ?tag=[ tag,tag,tag... | "-" ]
	Tag string `schema:"tag"`

	//* ?q=
	Q string `schema:"q"`

	//* ?public=true|false
	Public string `schema:"public"`

	//* ?deadline_gte=<isodate>
	DeadlineGTE string `schema:"deadline_gte"`

	//* ?deadline_lte=<isodate>
	DeadlineLTE string `schema:"deadline_lte"`

	//* ?done_gte=<isodate>
	DoneGTE string `schema:"done_gte"`

	//* ?done_lte=<isodate>
	DoneLTE string `schema:"done_lte"`

	//* ?created_gte=<isodate>
	CreatedGTE string `schema:"created_gte"`

	//* ?created_lte=<isodate>
	CreatedLTE string `schema:"created_lte"`

	//* ?short=true|false
	Short string `schema:"short"`

	// gentime great than group/chat
	GentimeGT int64 `schema:"gentime_gt"`
}
