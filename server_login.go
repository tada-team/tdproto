package tdproto

func NewServerLogin(deviceName string) (r ServerLogin) {
	r.Name = r.GetName()
	r.Params.DeviceName = deviceName
	return r
}

type ServerLogin struct {
	BaseEvent
	Params serverLoginParams `json:"params"`
}

func (p ServerLogin) GetName() string { return "server.login" }

type serverLoginParams struct {
	DeviceName string `json:"device_name"`
}
