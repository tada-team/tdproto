package tdproto

type GroupAccessRequest struct {
	Uid     string            `json:"uid"`
	Created ISODateTimeString `json:"created"`
	Subject *JID              `json:"subject"`
}
