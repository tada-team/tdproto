package tdproto

// Audio call information
type CallEvent struct {
	// Call start
	Start *ISODateTimeString `json:"start"`

	// Call finish
	Finish *ISODateTimeString `json:"finish"`

	// Call record enabled
	Audiorecord bool `json:"audiorecord"`

	// Call participants
	Onliners []CallOnliner `json:"onliners,omitempty"`
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
