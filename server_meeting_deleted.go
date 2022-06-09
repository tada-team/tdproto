package tdproto

func NewServerMeetingDeleted(meetingId string, countTeam, countUser int32) (r ServerMeetingDeleted) {
	r.Name = r.GetName()
	r.Params.MeetingId = meetingId
	r.Params.TeamMeetingsCount = countTeam
	r.Params.UserMeetingsCount = countUser
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
	MeetingId string `json:"meeting_id"`
	// Team Meetings count
	TeamMeetingsCount int32 `json:"team_meetings_count"`
	// User Meetings count
	UserMeetingsCount int32 `json:"user_meetings_count"`
	// Dates of team meetings
	TeamMeetingsDates []string `json:"team_meetings_dates"`
	// Dates of user meetings
	UserMeetingsDates []string `json:"user_meetings_dates"`
}
