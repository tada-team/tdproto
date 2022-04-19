package tdproto

type MeetingMemberStatus string

const (
	MeetingMemberStatusOwner          MeetingMemberStatus = "MEETING_MEMBER_STATUS_OWNER"
	MeetingMemberStatusAdmin          MeetingMemberStatus = "MEETING_MEMBER_STATUS_ADMIN"
	MeetingMemberStatusMember         MeetingMemberStatus = "MEETING_MEMBER_STATUS_MEMBER"
	MeetingMemberStatusOptionalMember MeetingMemberStatus = "MEETING_MEMBER_STATUS_OPTIONAL_MEMBER"
)

type MeetingPresenceStatus string

const (
	MeetingPresenceStatusAccepted MeetingPresenceStatus = "MEETING_PRESENCE_STATUS_ACCEPTED"
	MeetingPresenceStatusRejected MeetingPresenceStatus = "MEETING_PRESENCE_STATUS_REJECTED"
	MeetingPresenceStatusDoubts   MeetingPresenceStatus = "MEETING_PRESENCE_STATUS_DOUBTS"
	MeetingPresenceStatusWaiting  MeetingPresenceStatus = "MEETING_PRESENCE_STATUS_WAITING"
)

type MeetingRepeatabilityType string

const (
	MeetingRepeatabilityStatusUnspecified MeetingRepeatabilityType = "MEETING_PRESENCE_STATUS_UNSPECIFIED"
	MeetingRepeatabilityStatusDaily       MeetingRepeatabilityType = "MEETING_PRESENCE_STATUS_DAILY"
	MeetingRepeatabilityTypeWeekly        MeetingRepeatabilityType = "MEETING_PRESENCE_STATUS_WEEKLY"
	MeetingRepeatabilityTypeMonthly       MeetingRepeatabilityType = "MEETING_PRESENCE_STATUS_MONTHLY"
	MeetingRepeatabilityTypeNDOW          MeetingRepeatabilityType = "MEETING_PRESENCE_STATUS_NDOW"
	MeetingRepeatabilityTypeAnnually      MeetingRepeatabilityType = "MEETING_PRESENCE_STATUS_ANNUALLY"
)
