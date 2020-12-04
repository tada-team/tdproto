package tdproto

type GroupAccessRequest struct {
	Uid     string   `json:"uid"`
	Created DateTime `json:"created"`
	Subject *JID     `json:"subject"`
}
