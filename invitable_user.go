package tdproto

// Account from other team, Active Directory or server
type InvitableUser struct {
	// Account id
	Uid string `json:"uid"`

	// Node uid for external users
	Node string `json:"node,omitempty"`

	// Full name
	DisplayName string `json:"display_name"`

	// Icons
	Icons IconData `json:"icons"`
}
