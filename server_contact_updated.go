package tdproto

func NewServerContactUpdated(contacts ...Contact) (r ServerContactUpdated) {
	r.Name = r.GetName()
	r.Params.Contacts = contacts
	return r
}

// Contact created or updated
type ServerContactUpdated struct {
	BaseEvent
	Params serverContactUpdatedParams `json:"params"`
}

func (p ServerContactUpdated) GetName() string { return "server.contact.updated" }

// Params of the server.contact.updated event
type serverContactUpdatedParams struct {
	// Contact info
	Contacts []Contact `json:"contacts"`
}
