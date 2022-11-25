package tdproto

func NewServerMeetingCellUpdated(meetingCell Meeting) (r ServerMeetingCellUpdated) {
	r.Name = r.GetName()
	r.Params.MeetingCell = meetingCell
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
	MeetingCell Meeting `json:"meeting_cell"`
}
