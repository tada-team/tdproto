package tdproto

func NewServerContactUpdated(contacts ...Contact) (r ServerContactUpdated) {
	r.BaseEvent.Name = "server.contact.updated"
	r.Params.Contacts = contacts
	return r
}

type ServerContactUpdated struct {
	BaseEvent
	Params ServerContactUpdatedParams `json:"params"`
}

type ServerContactUpdatedParams struct {
	Contacts []Contact `json:"contacts"`
}
