package tdws

import "github.com/tada-team/tdproto"

// TODO: remove empty
func NewServerTeamCounters(counters []tdproto.TeamCounter, badge uint) ServerTeamCounters {
	r := ServerTeamCounters{}
	r.Name = r.GetName()
	r.Params.Teams = counters
	r.Params.Badge = badge
	return r
}

// Counters form other teams
type ServerTeamCounters struct {
	BaseEvent
	Params serverTeamCountersParams `json:"params"`
}

func (p ServerTeamCounters) GetName() string { return "server.team.counters" }

// Params of the server.team.counters event
type serverTeamCountersParams struct {
	// Counters
	Teams []tdproto.TeamCounter `json:"teams"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
