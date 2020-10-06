package tdproto

func NewServerTagDeleted(tags ...DeletedTag) (r ServerTagDeleted) {
	r.Name = r.GetName()
	r.Params.Tags = tags
	return r
}

type ServerTagDeleted struct {
	BaseEvent
	Params serverTagDeletedParams `json:"params"`
}

func (p ServerTagDeleted) GetName() string { return "server.tag.deleted" }

type serverTagDeletedParams struct {
	Tags []DeletedTag `json:"tags"`
}
