package tdproto

func NewServerSectionDeleted(ct ChatType, section DeletedSection) (r ServerSectionDeleted) {
	r.Name = r.GetName()
	r.Params.ChatType = ct
	r.Params.Gentime = Gentime() // XXX
	r.Params.Sections = []DeletedSection{section}
	return r
}

type ServerSectionDeleted struct {
	BaseEvent
	Params serverSectionDeletedParams `json:"params"`
}

func (p ServerSectionDeleted) GetName() string { return "server.section.deleted" }

type serverSectionDeletedParams struct {
	ChatType ChatType         `json:"chat_type"`
	Gentime  int64            `json:"gentime"`
	Sections []DeletedSection `json:"sections"`
}
