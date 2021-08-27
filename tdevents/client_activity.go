package tdevents

func NewClientActivity(afk bool) (r ClientActivity) {
	r.Name = r.GetName()
	r.Params.Afk = afk
	return r
}

// Change AFK (away from keyboard) status
type ClientActivity struct {
	BaseEvent
	Params clientActivityParams `json:"params"`
}

func (p ClientActivity) GetName() string { return "client.activity" }

// Params of the client.activity event
type clientActivityParams struct {
	// Is away from keyboard
	Afk bool `json:"afk"`
}
