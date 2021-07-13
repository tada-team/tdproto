package tdapi

// Server responce
type Resp struct {
	DebugTime string            `json:"_time,omitempty"`
	Ok        bool              `json:"ok"`
	Result    interface{}       `json:"result,omitempty"`
	Error     Err               `json:"error,omitempty"`
	Details   map[string]string `json:"details,omitempty"`
}
