package tdapi

import "github.com/tada-team/tdproto"

type Task struct {
	CustomColorIndex *uint16       `json:"custom_color_index,omitempty"`
	Description      string        `json:"description,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	SectionUid       string        `json:"section,omitempty"`
	Observers        []tdproto.JID `json:"observers,omitempty"`
	Items            []string      `json:"items,omitempty"`
	Assignee         tdproto.JID   `json:"assignee,omitempty"`
	Deadline         string        `json:"deadline,omitempty"`
	Public           bool          `json:"public,omitempty"`
	RemindAt         string        `json:"remind_at,omitempty"`
	TaskStatus       string        `json:"task_status,omitempty"`
	Importance       *int          `json:"importance,omitempty"`
	Urgency          *int          `json:"urgency,omitempty"`
	Complexity       *int          `json:"complexity,omitempty"`
	SpentTime        *int          `json:"spent_time,omitempty"`
	LinkedMessages   []string      `json:"linked_messages,omitempty"`
}
