package tdproto

// TODO: remove empty
func NewServerTeamCounters(counters []TeamCounter, badge uint) ServerTeamCounters {
	r := ServerTeamCounters{}
	r.Name = r.GetName()
	r.Params.Teams = counters
	r.Params.Badge = badge
	return r
}

type ServerTeamCounters struct {
	BaseEvent
	Params serverTeamCountersParams `json:"params"`
}

func (p ServerTeamCounters) GetName() string { return "server.team.counters" }

type serverTeamCountersParams struct {
	Teams []TeamCounter `json:"teams"`
	Badge uint          `json:"badge"`
}
