package tdapi

// Device ping response
type DevicePing struct {
	// Unique value sent to device
	PingId string `json:"ping_id"`
}
