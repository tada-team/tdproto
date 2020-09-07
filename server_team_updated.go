package tdproto

func NewServerTeamUpdated(team Team) (r ServerTeamUpdated) {
	r.BaseEvent.Name = "server.team.updated"
	r.Params.Teams = []Team{team}
	return r
}

type ServerTeamUpdated struct {
	BaseEvent
	Params ServerTeamUpdatedParams `json:"params"`
}

type ServerTeamUpdatedParams struct {
	Teams []Team `json:"teams"`
}
