package tdproto

type ClientActivity struct {
	BaseEvent
	Params clientActivityParams `json:"params"`
}

type clientActivityParams struct {
	Afk bool `json:"afk"`
}
