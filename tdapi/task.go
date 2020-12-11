package tdapi

import "github.com/tada-team/tdproto"

type Task struct {
	CustomColorIndex *uint16       `json:"custom_color_index,omitempty"`
	Description      string        `json:"description,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	SectionUid       string        `json:"section,omitempty"`
	Observers        []tdproto.JID `json:"observers,omitempty"` // TODO: rename to "followers"
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
	LinkedMessages   []string      `json:"linked_messages,omitempty"` // TODO: Message object
	Uploads          []string      `json:"uploads,omitempty"`
}

type TaskFilter struct {
	UserParams
	Paginator

	//* ?sort = [ "created" | "-created" | "last_message" | "-last_message" | "deadline" | "-deadline" ]
	Sort string `schema:"sort"`

	//* ?task_status = ["new" | "done" ]
	TaskStatus string `schema:"task_status"`

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
	Section string `schema:"section"`

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
}
