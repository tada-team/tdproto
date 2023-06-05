package tdproto

// Account from other team, Active Directory or node
type InvitableUser struct {
	// Account id
	Uid string `json:"uid"`

	// Node uid for external users
	Node string `json:"node,omitempty"`

	// Full name
	DisplayName string `json:"display_name"`

	// Icons
	Icons IconData `json:"icons"`

	// Common team uids, if any
	Teams []string `json:"teams,omitempty"`

	// Флаг нахождения пользователя на другом аккаунте
	FromAnotherAccount bool `json:"from_another_account"`
}
