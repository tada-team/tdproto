package tdproto

func NewServerUiSettings(v *UiSettings) (r ServerUiSettings) {
	r.BaseEvent.Name = "server.uisettings"
	r.Params = v
	return r
}

type ServerUiSettings struct {
	BaseEvent
	Params *UiSettings `json:"params"`
}
