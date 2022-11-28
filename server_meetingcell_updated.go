package tdproto

func NewServerMeetingCellUpdated(meetingID string, startAt ISODateTimeString, duration int32) (r ServerMeetingCellUpdated) {
	r.Name = r.GetName()
	r.Params.MeetingID = meetingID
	r.Params.StartAt = startAt
	r.Params.Duration = duration
	return r
}

// Meeting Cell updated
type ServerMeetingCellUpdated struct {
	BaseEvent
	Params serverMeetingCellUpdatedParams `json:"params"`
}

func (p ServerMeetingCellUpdated) GetName() string { return "server.meetingcell.updated" }

// Params of the server.meetingcell.updated event
type serverMeetingCellUpdatedParams struct {
	// Meeting Cell info
	MeetingID string            `json:"meeting_id"`
	StartAt   ISODateTimeString `json:"start_at"`
	Duration  int32             `json:"duration"`
}
