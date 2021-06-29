package tdws

// deprecated
func NewServerRoster(roster Roster) (r ServerRoster) {
	r.Name = r.GetName()
	r.Params = roster
	return r
}

// deprecated
type ServerRoster struct {
	BaseEvent
	Params Roster `json:"params"`
}

func (p ServerRoster) GetName() string { return "server.roster" }
