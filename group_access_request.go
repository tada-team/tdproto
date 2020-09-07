package tdproto

type GroupAccessRequest struct {
	Uid     string `json:"uid"`
	Created string `json:"created"`
	Subject *JID   `json:"subject"`
}
