package tdproto

func NewServerMeetingUpdated(meeting Meeting, count int32) (r ServerMeetingUpdated) {
	return NewServerMeetingsUpdated([]Meeting{meeting}, count)
}

func NewServerMeetingsUpdated(meetings []Meeting, count int32) (r ServerMeetingUpdated) {
	r.Name = r.GetName()
	r.Params.Meetings = meetings
	r.Params.MeetingsCount = count
	return r
}

// Meeting created or updated
type ServerMeetingUpdated struct {
	BaseEvent
	Params serverMeetingUpdatedParams `json:"params"`
}

func (p ServerMeetingUpdated) GetName() string { return "server.meeting.updated" }

// Params of the server.meeting.updated event
type serverMeetingUpdatedParams struct {
	// Meeting info
	Meetings []Meeting `json:"meetings"`
	// Meetings count
	MeetingsCount int32 `json:"meetings_count"`
}
