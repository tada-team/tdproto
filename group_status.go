package tdproto

// Status in team
type GroupStatus string

const (
	// Group administrator
	GroupAdmin GroupStatus = "admin" // 3

	// Group member
	GroupMember GroupStatus = "member" // 2
)
