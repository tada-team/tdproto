package tdproto

// Push message over websockets. Readonly.
type MessagePush struct {
	// Push title
	Title string `json:"title"`

	// Push subtitle
	Subtitle string `json:"subtitle"`

	// Push body
	Message string `json:"message"`

	// Absolute url to push icon
	IconUrl string `json:"icon_url"`

	// Url opened on click
	ClickAction string `json:"click_action"`

	// Push tag (for join pushes)
	Tag string `json:"tag"`

	// Team uid
	Team string `json:"team"`

	// Sender contact id
	Sender *JID `json:"sender"`

	// Chat id
	Chat *JID `json:"chat"`

	// Message id
	MessageId string `json:"message_id"`

	// Message creation iso datetime
	Created string `json:"created"`
}
