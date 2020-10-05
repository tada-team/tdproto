package tdproto

func NewServerTagUpdated(tags ...Tag) (r ServerTagUpdated) {
	r.Name = r.GetName()
	r.Params.Tags = tags
	return r
}

type ServerTagUpdated struct {
	BaseEvent
	Params serverTagUpdatedParams `json:"params"`
}

func (p ServerTagUpdated) GetName() string { return "server.tag.updated" }

type serverTagUpdatedParams struct {
	Tags []Tag `json:"tags"`
}
