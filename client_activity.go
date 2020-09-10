package tdproto

func NewClientActivity(afk bool) (r ClientActivity) {
	r.Name = "client.activity"
	r.Params.Afk = afk
	return r
}

type ClientActivity struct {
	BaseEvent
	Params clientActivityParams `json:"params"`
}

type clientActivityParams struct {
	Afk bool `json:"afk"`
}
