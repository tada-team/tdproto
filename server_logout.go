package tdproto

func NewServerLogout(reason string) (r ServerLogout) {
	r.Name = r.GetName()
	r.Params.Reason = reason
	return r
}

type ServerLogout struct {
	BaseEvent
	Params serverLogoutParams `json:"params"`
}

func (p ServerLogout) GetName() string { return "server.logout" }

// Params of the server.logout event
type serverLogoutParams struct {
	// Reason
	Reason string `json:"reason"`
}
