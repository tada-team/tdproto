package tdws

import "github.com/tada-team/tdproto"

func NewServerTeamDeleted(uid string, gentime int64) (r ServerTeamDeleted) {
	r.Name = r.GetName()
	r.Params.Teams = []tdproto.DeletedTeam{{
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

// Params of the server.team.deleted event
type serverTeamDeletedParams struct {
	// Teams info
	Teams []tdproto.DeletedTeam `json:"teams"`
}
