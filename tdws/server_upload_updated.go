package tdws

import "github.com/tada-team/tdproto"

func NewServerUploadUpdated(uploads ...tdproto.Upload) (r ServerUploadUpdated) {
	r.Name = r.GetName()
	r.Params.Uploads = uploads
	return r
}

// Upload object created or changed
type ServerUploadUpdated struct {
	BaseEvent
	Params serverUploadUpdatedParams `json:"params"`
}

func (p ServerUploadUpdated) GetName() string { return "server.upload.updated" }

// Params of the server.upload.updated event
type serverUploadUpdatedParams struct {
	// Uploads data
	Uploads []tdproto.Upload `json:"uploads"`
}
