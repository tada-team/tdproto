package tdproto

func NewServerUiSettings(namespace string, data UiSettingsData) (r ServerUiSettings) {
	r.Name = r.GetName()
	r.Params = ServerUiSettingsParams{
		Namespace: namespace,
		Data:      data,
	}
	return r
}

// Part of UI settings changed
type ServerUiSettings struct {
	BaseEvent
	Params ServerUiSettingsParams `json:"params"`
}

type ServerUiSettingsParams struct {
	// Namespace. For example: web, app
	Namespace string `json:"namespace"`

	// UiSettingsData
	Data UiSettingsData `json:"data"`
}

func (p ServerUiSettings) GetName() string { return "server.uisettings" }
