package tdevents

import "github.com/tada-team/tdproto"

func NewServerTagDeleted(tags ...tdproto.DeletedTag) (r ServerTagDeleted) {
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

// Params of the server.tag.deleted event
type serverTagDeletedParams struct {
	// Tags info
	Tags []tdproto.DeletedTag `json:"tags"`
}
