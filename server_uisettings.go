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
	Params ServerUiSettingsParams `json:"params"`
	BaseEvent
}

type ServerUiSettingsParams struct {
	// UiSettingsData
	Data UiSettingsData `json:"data"`

	// Namespace. For example: web, app
	Namespace string `json:"namespace"`
}

func (p ServerUiSettings) GetName() string { return "server.uisettings" }
