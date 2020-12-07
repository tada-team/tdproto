package tdproto

// Team status
type TeamStatus string

const (
	// Team owner. Can do anything
	TeamOwner TeamStatus = "owner"

	// Team administrator
	TeamAdmin TeamStatus = "admin"

	// Team member
	TeamMember TeamStatus = "member"

	// Team guest. Restricted account
	TeamGuest TeamStatus = "guest"
)
