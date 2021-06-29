package tdws

func NewServerTagUpdated(tags ...Tag) (r ServerTagUpdated) {
	r.Name = r.GetName()
	r.Params.Tags = tags
	return r
}

// Tag created or changed
type ServerTagUpdated struct {
	BaseEvent
	Params serverTagUpdatedParams `json:"params"`
}

func (p ServerTagUpdated) GetName() string { return "server.tag.updated" }

// Params of the server.tag.updated event
type serverTagUpdatedParams struct {
	// Tags info
	Tags []Tag `json:"tags"`
}
