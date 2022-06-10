package tdproto

func NewServerMeetingDeleted(meetingIds ...string) (r ServerMeetingDeleted) {
	r.Name = r.GetName()
	r.Params.MeetingIds = meetingIds
	return r
}

// Meeting deleted
type ServerMeetingDeleted struct {
	BaseEvent
	Params serverMeetingDeletedParams `json:"params"`
}

func (p ServerMeetingDeleted) GetName() string { return "server.meeting.deleted" }

// Params of the server.meeting.deleted event
type serverMeetingDeletedParams struct {
	// Meeting info
	MeetingIds []string `json:"meeting_id"`
}
