package tdproto

// Status in team
type GroupStatus string

const (
	// Group administrator
	GroupAdmin GroupStatus = "admin" // 2

	// Group member
	GroupMember GroupStatus = "member" // 3
)
