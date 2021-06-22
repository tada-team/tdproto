package tdproto

func NewServerLogin(deviceName string) (r ServerLogin) {
	r.Name = r.GetName()
	r.Params.DeviceName = deviceName
	return r
}

// Login from other device
type ServerLogin struct {
	BaseEvent
	Params serverLoginParams `json:"params"`
}

func (p ServerLogin) GetName() string { return "server.login" }

// Params of the server.login event
type serverLoginParams struct {
	// Device name
	DeviceName string `json:"device_name"`
}
