package tdproto

func NewServerSectionUpdated(ct ChatType, sections ...Section) (r ServerSectionUpdated) {
	r.Name = r.GetName()
	r.Params.ChatType = ct
	r.Params.Gentime = Gentime() // XXX
	r.Params.Sections = sections
	return r
}

// Contact section or task project created or changed
type ServerSectionUpdated struct {
	BaseEvent
	Params serverSectionUpdatedParams `json:"params"`
}

func (p ServerSectionUpdated) GetName() string { return "server.section.updated" }

// Params of the server.section.updated event
type serverSectionUpdatedParams struct {
	// Chat type
	ChatType ChatType `json:"chat_type"`

	// Section/project info
	Sections []Section `json:"sections"`

	// deprecated
	Gentime int64 `json:"gentime"`
}
