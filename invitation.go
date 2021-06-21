package tdproto

// Invitation to team
type Invitation struct {
	Uid     string            `json:"uid"`
	Token   string            `json:"token"`
	Created ISODateTimeString `json:"created"`
	Qr      string            `json:"qr"`
}
