package tdproto

func NewServerTeamDeleted(uid string, gentime int64) (r ServerTeamDeleted) {
	r.Name = r.GetName()
	r.Params.Teams = []DeletedTeam{{
		Uid:       uid,
		IsArchive: true,
		Gentime:   gentime,
	}}
	return r
}

// Team archived
type ServerTeamDeleted struct {
	BaseEvent
	Params serverTeamDeletedParams `json:"params"`
}

func (p ServerTeamDeleted) GetName() string { return "server.team.deleted" }

type serverTeamDeletedParams struct {
	// Teams info
	Teams []DeletedTeam `json:"teams"`
}
