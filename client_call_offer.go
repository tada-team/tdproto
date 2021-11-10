package tdproto

// Start a call
type ClientCallOffer struct {
	BaseEvent
	Params clientCallOfferParams `json:"params"`
}

func (p ClientCallOffer) GetName() string { return "client.call.offer" }

// Params of the client.call.offer event
type clientCallOfferParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Mute state
	// Deprecated: use EnabledAudio
	Muted bool `json:"muted"`

	// Is trickle mode enabled
	Trickle bool `json:"trickle"`

	// SDP (session description protocol) data
	Sdp string `json:"sdp"`

	// CallType is a type of call("audio" - audio room, "video" - video room). default = "audio"
	CallType CallType `json:"call_type,omitempty"`

	// Audio state
	EnabledAudio bool `json:"enabled_audio,omitempty"`

	// Video state
	EnabledVideo bool `json:"enabled_video,omitempty"`
}
