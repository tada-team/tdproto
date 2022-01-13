package tdproto

import (
	"time"
)

// Task Preview
type TaskPreview struct {
	Deadline    *time.Time  `json:"deadline"`
	Error       string      `json:"_error,omitempty"`
	Assignee    JID         `json:"assignee"`
	Description string      `json:"description"`
	Section     string      `json:"section"`
	Tags        []string    `json:"tags"`
	Items       []TaskItems `json:"items"`
	Public      bool        `json:"public"`
}

// Task item
type TaskItems struct {
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}
