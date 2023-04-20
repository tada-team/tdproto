package tdproto

// Call information
type CallEvent struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// Call buzzing
	Buzz bool `json:"buzz,omitempty"`

	// Creation date, iso datetime
	Created ISODateTimeString `json:"created"`

	// Call start. For direct calls can be empty when buzzing
	Start ISODateTimeString `json:"start,omitempty"`

	// Call finish
	Finish ISODateTimeString `json:"finish,omitempty"`

	// Call record enabled
	Audiorecord bool `json:"audiorecord"`

	// Call participants
	Onliners []CallOnliner `json:"onliners,omitempty"`

	// Version
	Gentime int64 `json:"gentime"`

	// Deprecated: use gentime or created
	Timestamp int64 `json:"timestamp"`
}

// Call participant
type CallOnliner struct {
	// Contact id
	Jid JID `json:"jid"`

	// Contact name
	DisplayName string `json:"display_name"`

	// Contact role
	Role string `json:"role"`

	// Contact icon
	Icon string `json:"icon"`

	// Microphone muted. Computed from devices muted states
	Muted bool `json:"muted"`

	// Video state
	EnabledVideo bool `json:"enabled_video"`

	// Screenshare state
	EnabledScreenshare bool `json:"enabled_screenshare"`

	// HandUp - status of hand in VCS
	HandUp bool `json:"hand_up"`

	// Member devices, strictly one for now
	Devices []CallDevice `json:"devices"`
}

// Call participant device
type CallDevice struct {
	// Device muted
	Muted bool `json:"muted"`

	// Device description
	Useragent string `json:"useragent"`
}
