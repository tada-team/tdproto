package tdapi

// Simple api for integrations /api/message or /tasks/[team]/[num]/message
type EasyApiMessage struct {
	// Comma separated api keys (for /api/message calls only)
	Key string `json:"key"`

	// Message text. Required.
	Text string `json:"message"`

	// Message uuid. Optional
	MessageUid string `json:"message_id"`

	// Disable links preview
	Nopreview bool `json:"nopreview"`

	// Mark message as important
	Important bool `json:"important"`

	// Disable counters and push notifications
	Silently bool `json:"silently"`

	// Convert "\\n" to "\n"
	ConvertLinebreaks bool `json:"convert_linebreaks"`

}
