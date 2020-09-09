package tdproto

func NewServerTagUpdated(tags ...Tag) (r ServerTagUpdated) {
	r.BaseEvent.Name = "server.tag.updated"
	r.Params.Tags = tags
	return r
}

type ServerTagUpdated struct {
	BaseEvent
	Params serverTagUpdatedParams `json:"params"`
}

type serverTagUpdatedParams struct {
	Tags []Tag `json:"tags"`
}
