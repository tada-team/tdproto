package tdproto

// Invitation to team
type Invitation struct {
	Uid     string            `json:"uid"`
	Token   string            `json:"token"`
	Created ISODateTimeString `json:"created"`
	Qr      string            `json:"qr"`
}

// InvitationToMeeting Preview
type InvitationToMeeting struct {
	Description string            `json:"description"`
	DateStart   ISODateTimeString `json:"date_start"`
	MeetingJid  string            `json:"meeting_jid"`
}

// Guest Invitation Form
type MeetingNewGuest struct {
	FullName string `json:"full_name"`
}
