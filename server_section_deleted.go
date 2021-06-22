package tdproto

func NewServerSectionDeleted(ct ChatType, section DeletedSection) (r ServerSectionDeleted) {
	section.Gentime = Gentime() // XXX

	r.Name = r.GetName()
	r.Params.ChatType = ct
	r.Params.Gentime = section.Gentime
	r.Params.Sections = []DeletedSection{section}
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
	ChatType ChatType `json:"chat_type"`

	// Section/project info
	Sections []DeletedSection `json:"sections"`

	// Deprecated
	Gentime int64 `json:"gentime"`
}
