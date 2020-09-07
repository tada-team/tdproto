package tdproto

func NewServerRemindDeleted(uid string) (r ServerRemindDeleted) {
	r.BaseEvent.Name = "server.remind.deleted"
	r.Params.Remind = []DeletedRemind{
		{Uid: uid},
	}
	return r
}

type ServerRemindDeleted struct {
	BaseEvent
	Params ServerRemindDeletedParams `json:"params"`
}

type ServerRemindDeletedParams struct {
	Remind []DeletedRemind `json:"reminds"`
}
