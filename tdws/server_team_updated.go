package tdws

import "github.com/tada-team/tdproto"

func NewServerTeamUpdated(team tdproto.Team) (r ServerTeamUpdated) {
	r.Name = r.GetName()
	r.Params.Teams = []tdproto.Team{team}
	return r
}

// Team created or changed
type ServerTeamUpdated struct {
	BaseEvent
	Params serverTeamUpdatedParams `json:"params"`
}

func (p ServerTeamUpdated) GetName() string { return "server.team.updated" }

// Params of the server.team.updated event
type serverTeamUpdatedParams struct {
	Teams []tdproto.Team `json:"teams"`
}
