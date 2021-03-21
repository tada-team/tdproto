package tdproto

func NewClientActivity(afk bool) (r ClientActivity) {
	r.Name = r.GetName()
	r.Params.Afk = afk
	return r
}

// Change AFG (away from keyboard) status
type ClientActivity struct {
	BaseEvent
	Params clientActivityParams `json:"params"`
}

func (p ClientActivity) GetName() string { return "client.activity" }

type clientActivityParams struct {
	Afk bool `json:"afk"`
}
