package tdevents

import "github.com/tada-team/tdproto"

func NewServerUiSettings(v *tdproto.UiSettings) (r ServerUiSettings) {
	r.Name = r.GetName()
	r.Params = v
	return r
}

// Part of UI settings changed
type ServerUiSettings struct {
	BaseEvent
	Params *tdproto.UiSettings `json:"params"`
}

func (p ServerUiSettings) GetName() string { return "server.uisettings" }
