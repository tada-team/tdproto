package tdws

import "github.com/tada-team/tdproto"

func NewServerRemindUpdated(reminds ...tdproto.Remind) (r ServerRemindUpdated) {
	r.Name = r.GetName()
	r.Params.Reminds = reminds
	return r
}

// Task/group remind created or changed
type ServerRemindUpdated struct {
	BaseEvent
	Params serverRemindUpdatedParams `json:"params"`
}

func (p ServerRemindUpdated) GetName() string { return "server.remind.updated" }

// Params of the server.remind.updated event
type serverRemindUpdatedParams struct {
	// Remind information
	Reminds []tdproto.Remind `json:"reminds"`
}
