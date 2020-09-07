package tdproto

type GroupStatus string

const (
	GroupAdmin  = GroupStatus("admin")  // 3
	GroupMember = GroupStatus("member") // 2
)
