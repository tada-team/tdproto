package tdproto

type BaseEvent struct {
	Name      string `json:"event"`
	ConfirmId string `json:"confirm_id,omitempty"`
}

type AnyEvent struct {
	BaseEvent
	Params *map[string]interface{} `json:"params"`
}

func (ev BaseEvent) GetName() string      { return ev.Name }
func (ev BaseEvent) GetConfirmId() string { return ev.ConfirmId }

func (ev BaseEvent) String() string {
	if ev.ConfirmId != "" {
		return "*" + ev.Name
	}
	return ev.Name
}

func NewAnyEvent(name string) AnyEvent {
	params := make(map[string]interface{})
	return AnyEvent{
		BaseEvent: BaseEvent{Name: name},
		Params:    &params,
	}
}

type Event interface {
	GetName() string
	GetConfirmId() string
}
