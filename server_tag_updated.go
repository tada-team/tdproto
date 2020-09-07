package tdproto

func NewServerTagUpdated(tags ...Tag) (r ServerTagUpdated) {
	r.BaseEvent.Name = "server.tag.updated"
	r.Params.Tags = tags
	return r
}

type ServerTagUpdated struct {
	BaseEvent
	Params ServerTagUpdatedParams `json:"params"`
}

type ServerTagUpdatedParams struct {
	Tags []Tag `json:"tags"`
}
