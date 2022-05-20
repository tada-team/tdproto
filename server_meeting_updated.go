package tdproto

func NewServerMeetingUpdated(meeting Meeting) (r ServerMeetingUpdated) {
	return NewServerMeetingsUpdated([]Meeting{meeting})
}

func NewServerMeetingsUpdated(meetings []Meeting) (r ServerMeetingUpdated) {
	r.Name = r.GetName()
	r.Params.Meetings = meetings
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
}
