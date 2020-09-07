package tdproto

// deprecated
func NewServerRoster(roster Roster) (r ServerRoster) {
	r.BaseEvent.Name = "server.roster"
	r.Params = roster
	return r
}

// deprecated
type ServerRoster struct {
	BaseEvent
	Params Roster `json:"params"`
}
