package tdproto

// TODO: remove empty
func NewServerTeamCounters(counters []TeamCounter, badge uint) ServerTeamCounters {
	r := ServerTeamCounters{}
	r.Name = "server.team.counters"
	r.Params.Teams = counters
	r.Params.Badge = badge
	return r
}

type ServerTeamCounters struct {
	BaseEvent
	Params serverTeamCountersParams `json:"params"`
}

type serverTeamCountersParams struct {
	Teams []TeamCounter `json:"teams"`
	Badge uint          `json:"badge"`
}
