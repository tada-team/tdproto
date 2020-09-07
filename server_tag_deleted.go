package tdproto

func NewServerTagDeleted(tags ...DeletedTag) (r ServerTagDeleted) {
	r.BaseEvent.Name = "server.tag.deleted"
	r.Params.Tags = tags
	return r
}

type ServerTagDeleted struct {
	BaseEvent
	Params ServerTagDeletedParams `json:"params"`
}

type ServerTagDeletedParams struct {
	Tags []DeletedTag `json:"tags"`
}
