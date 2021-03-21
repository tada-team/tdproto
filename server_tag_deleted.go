package tdproto

func NewServerTagDeleted(tags ...DeletedTag) (r ServerTagDeleted) {
	r.Name = r.GetName()
	r.Params.Tags = tags
	return r
}

// Tag deleted
type ServerTagDeleted struct {
	BaseEvent
	Params serverTagDeletedParams `json:"params"`
}

func (p ServerTagDeleted) GetName() string { return "server.tag.deleted" }

type serverTagDeletedParams struct {
	// Tags info
	Tags []DeletedTag `json:"tags"`
}
