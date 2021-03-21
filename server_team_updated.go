package tdproto

func NewServerTeamUpdated(team Team) (r ServerTeamUpdated) {
	r.Name = r.GetName()
	r.Params.Teams = []Team{team}
	return r
}

// Team created or changed
type ServerTeamUpdated struct {
	BaseEvent
	Params serverTeamUpdatedParams `json:"params"`
}

func (p ServerTeamUpdated) GetName() string { return "server.team.updated" }

type serverTeamUpdatedParams struct {
	Teams []Team `json:"teams"`
}
