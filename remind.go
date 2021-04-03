package tdproto

// Remind
type Remind struct {
	// Remind id
	Uid string `json:"uid"`

	// Chat id
	Chat *JID `json:"chat"`

	// Activation time, iso
	FireAt ISODateTimeString `json:"fire_at"`

	// Comment, if any
	Comment string `json:"comment,omitempty"`
}

// Remind deleted message
type DeletedRemind struct {
	// Remind id
	Uid string `json:"uid"`
}
