package tdevents

import "github.com/tada-team/tdproto"

func NewServerRemindDeleted(uid string) (r ServerRemindDeleted) {
	r.Name = r.GetName()
	r.Params.Remind = []tdproto.DeletedRemind{
		{Uid: uid},
	}
	return r
}

// Task or group remind deleted
type ServerRemindDeleted struct {
	BaseEvent
	Params serverRemindDeletedParams `json:"params"`
}

func (p ServerRemindDeleted) GetName() string { return "server.remind.deleted" }

// Params of the server.remind.deleted event
type serverRemindDeletedParams struct {
	// Remind information
	Remind []tdproto.DeletedRemind `json:"reminds"`
}
