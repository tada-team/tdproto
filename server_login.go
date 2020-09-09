package tdproto

func NewServerLogin(deviceName string) (r ServerLogin) {
	r.BaseEvent.Name = "server.login"
	r.Params.DeviceName = deviceName
	return r
}

type ServerLogin struct {
	BaseEvent
	Params serverLoginParams `json:"params"`
}

type serverLoginParams struct {
	DeviceName string `json:"device_name"`
}
