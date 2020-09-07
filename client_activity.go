package tdproto

type ClientActivity struct {
	BaseEvent
	Params ClientActivityParams `json:"params"`
}

type ClientActivityParams struct {
	Afk bool `json:"afk"`
}
