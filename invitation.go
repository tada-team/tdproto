package tdproto

// Invitation to team
type Invitation struct {
	Uid          string            `json:"uid"`
	Token        string            `json:"token"`
	Created      ISODateTimeString `json:"created"`
	Qr           string            `json:"qr"`
	InitialGroup string            `json:"initial_group,omitempty"`
}

// InvitationToMeeting Preview
type InvitationToMeeting struct {
	Description  string            `json:"description"`
	DateStart    ISODateTimeString `json:"date_start"`
	MeetingGroup string            `json:"meeting_group"`
}

type MeetingNewGuest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
