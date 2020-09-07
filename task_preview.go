package tdproto

import (
	"time"
)

type TaskPreview struct {
	Error       string      `json:"_error,omitempty"`
	Assignee    JID         `json:"assignee"`
	Deadline    *time.Time  `json:"deadline"`
	Description string      `json:"description"`
	Section     string      `json:"section"`
	Public      bool        `json:"public"`
	Tags        []string    `json:"tags"`
	Items       []TaskItems `json:"items"`
}

type TaskItems struct {
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}
