package tdproto

type Invitation struct {
	Uid     string   `json:"uid"`
	Token   string   `json:"token"`
	Created DateTime `json:"created"`
	Qr      string   `json:"qr"`
}
