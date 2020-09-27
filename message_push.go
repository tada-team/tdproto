package tdproto

type MessagePush struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Message     string `json:"message"`
	IconUrl     string `json:"icon_url"`
	ClickAction string `json:"click_action"`
	Tag         string `json:"tag"`
	Team        string `json:"team"`
	Sender      *JID   `json:"sender"`
	Chat        *JID   `json:"chat"`
	MessageId   string `json:"message_id"`
	Created     string `json:"created"`
}
