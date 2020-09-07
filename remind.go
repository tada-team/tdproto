package tdproto

type Remind struct {
	Uid     string `json:"uid"`
	Chat    *JID   `json:"chat"`
	FireAt  string `json:"fire_at"`
	Comment string `json:"comment,omitempty"`
}

type DeletedRemind struct {
	Uid string `json:"uid"`
}
