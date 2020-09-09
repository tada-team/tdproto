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
	Params serverRemindDeletedParams `json:"params"`
}

type serverRemindDeletedParams struct {
	Remind []DeletedRemind `json:"reminds"`
}
