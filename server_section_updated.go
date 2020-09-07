package tdproto

func NewServerSectionUpdated(ct ChatType, sections ...Section) (r ServerSectionUpdated) {
	r.BaseEvent.Name = "server.section.updated"
	r.Params.ChatType = ct
	r.Params.Gentime = Gentime() // XXX
	r.Params.Sections = sections
	return r
}

type ServerSectionUpdated struct {
	BaseEvent
	Params ServerSectionUpdatedParams `json:"params"`
}

type ServerSectionUpdatedParams struct {
	ChatType ChatType  `json:"chat_type"`
	Gentime  int64     `json:"gentime"`
	Sections []Section `json:"sections"`
}
