package tdproto

type MeetingMemberStatus string

const (
	MeetingMemberStatusOwner  MeetingMemberStatus = "owner"
	MeetingMemberStatusAdmin  MeetingMemberStatus = "admin"
	MeetingMemberStatusMember MeetingMemberStatus = "member"
)

type MeetingMemberPresence string

const (
	MeetingMemberStatusAccepted MeetingMemberStatus = "accepted"
	MeetingMemberStatusRejected MeetingMemberStatus = "rejected"
	MeetingMemberStatusDoubts   MeetingMemberStatus = "doubts"
	MeetingMemberStatusWaiting  MeetingMemberStatus = "waiting"
)
