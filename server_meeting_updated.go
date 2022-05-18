package tdproto

func NewServerMeetingUpdated(meeting Meeting) (r ServerMeetingUpdated) {
	r.Name = r.GetName()
	r.Params.Meetings = []Meeting{}
	return r
}

// Meeting created or updated
type ServerMeetingUpdated struct {
	BaseEvent
	Params serverMeetingUpdatedParams `json:"params"`
}

func (p serverMeetingUpdatedParams) GetName() string { return "server.meeting.updated" }

// Params of the server.meeting.updated event
type serverMeetingUpdatedParams struct {
	// Meeting info
	Meetings []Meeting `json:"meetings"`
}
