package tdproto

func NewServerContactUpdated(contacts ...Contact) (r ServerContactUpdated) {
	r.Name = r.GetName()
	r.Params.Contacts = contacts
	return r
}

type ServerContactUpdated struct {
	BaseEvent
	Params serverContactUpdatedParams `json:"params"`
}

func (p ServerContactUpdated) GetName() string { return "server.contact.updated" }

type serverContactUpdatedParams struct {
	Contacts []Contact `json:"contacts"`
}
