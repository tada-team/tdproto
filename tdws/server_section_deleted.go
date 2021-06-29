package tdws

import "github.com/tada-team/tdproto"

func NewServerSectionDeleted(ct tdproto.ChatType, section tdproto.DeletedSection) (r ServerSectionDeleted) {
	section.Gentime = tdproto.Gentime() // XXX

	r.Name = r.GetName()
	r.Params.ChatType = ct
	r.Params.Gentime = section.Gentime
	r.Params.Sections = []tdproto.DeletedSection{section}
	return r
}

// Contact section or task project deleted
type ServerSectionDeleted struct {
	BaseEvent
	Params serverSectionDeletedParams `json:"params"`
}

func (p ServerSectionDeleted) GetName() string { return "server.section.deleted" }

// Params of the server.section.deleted event
type serverSectionDeletedParams struct {
	// Chat type
	ChatType tdproto.ChatType `json:"chat_type"`

	// Section/project info
	Sections []tdproto.DeletedSection `json:"sections"`

	// Deprecated
	Gentime int64 `json:"gentime"`
}
