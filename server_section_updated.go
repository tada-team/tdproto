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

type serverSectionUpdatedParams struct {
	ChatType ChatType  `json:"chat_type"`
	Gentime  int64     `json:"gentime"`
	Sections []Section `json:"sections"`
}
