package tdproto

func NewServerMeetingUpdated(meeting Meeting, countTeam, countUser int32) (r ServerMeetingUpdated) {
	return NewServerMeetingsUpdated([]Meeting{meeting}, countTeam, countUser)
}

func NewServerMeetingsUpdated(meetings []Meeting, countTeam, countUser int32) (r ServerMeetingUpdated) {
	r.Name = r.GetName()
	r.Params.Meetings = meetings
	r.Params.TeamMeetingsCount = countTeam
	r.Params.UserMeetingsCount = countUser
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
	// Team Meetings count
	TeamMeetingsCount int32 `json:"team_meetings_count"`
	// User Meetings count
	UserMeetingsCount int32 `json:"user_meetings_count"`
}
