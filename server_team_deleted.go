package tdproto

func NewServerTeamDeleted(uid string, gentime int64) (r ServerTeamDeleted) {
	r.BaseEvent.Name = "server.team.deleted"
	r.Params.Teams = []DeletedTeam{{
		Uid:       uid,
		IsArchive: true,
		Gentime:   gentime,
	}}
	return r
}

type ServerTeamDeleted struct {
	BaseEvent
	Params serverTeamDeletedParams `json:"params"`
}

type serverTeamDeletedParams struct {
	Teams []DeletedTeam `json:"teams"`
}
