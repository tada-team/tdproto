package tdproto

func NewServerUiSettings(v *UiSettings) (r ServerUiSettings) {
	r.Name = r.GetName()
	r.Params = v
	return r
}

// Part of UI settings changed
type ServerUiSettings struct {
	BaseEvent
	Params *UiSettings `json:"params"`
}

func (p ServerUiSettings) GetName() string { return "server.uisettings" }
