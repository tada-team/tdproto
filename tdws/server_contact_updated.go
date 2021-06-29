package tdws

import "github.com/tada-team/tdproto"

func NewServerContactUpdated(contacts ...tdproto.Contact) (r ServerContactUpdated) {
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
	Contacts []tdproto.Contact `json:"contacts"`
}
