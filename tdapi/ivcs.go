package tdapi

// IVCSAuthResponse result of auth in IVCS
type IVCSAuthResponse struct {
	// SessionID IVA user session ID
	SessionID string `json:"session_id,omitempty"`
}
