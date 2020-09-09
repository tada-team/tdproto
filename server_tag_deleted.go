package tdproto

func NewServerTagDeleted(tags ...DeletedTag) (r ServerTagDeleted) {
	r.BaseEvent.Name = "server.tag.deleted"
	r.Params.Tags = tags
	return r
}

type ServerTagDeleted struct {
	BaseEvent
	Params serverTagDeletedParams `json:"params"`
}

type serverTagDeletedParams struct {
	Tags []DeletedTag `json:"tags"`
}
