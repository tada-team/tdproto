package tdproto

// Websocket session
type Session struct {
	// Session id
	Uid string `json:"uid"`

	// Creation datetime
	Created ISODateTimeString `json:"created"`

	// Language code
	Lang string `json:"lang,omitempty"`

	// Team id
	Team string `json:"team,omitempty"`

	// User agent
	Useragent string `json:"useragent,omitempty"`

	// IP address
	Addr string `json:"addr,omitempty"`

	// Mobile
	IsMobile bool `json:"is_mobile,omitempty"`

	// Away from keyboard
	Afk bool `json:"afk,omitempty"`
}
