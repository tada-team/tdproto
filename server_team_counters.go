package tdproto

// TODO: remove empty
func NewServerTeamCounters(counters []TeamCounter, badge uint) ServerTeamCounters {
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

type serverTeamCountersParams struct {
	// Counters
	Teams []TeamCounter `json:"teams"`

	// Total number of unreads
	Badge uint `json:"badge"`
}
