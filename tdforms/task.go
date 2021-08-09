package tdforms

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

	// Is task or group public for non-guests
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
