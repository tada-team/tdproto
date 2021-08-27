package tdevents

import "github.com/tada-team/tdproto"

func NewServerUiSettings(namespace string, data tdproto.UiSettingsData) (r ServerUiSettings) {
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
	Data tdproto.UiSettingsData `json:"data"`
}

func (p ServerUiSettings) GetName() string { return "server.uisettings" }
