package tdproto

func NewServerUploadUpdated(uploads ...Upload) (r ServerUploadUpdated) {
	r.Name = r.GetName()
	r.Params.Uploads = uploads
	return r
}

type ServerUploadUpdated struct {
	BaseEvent
	Params serverUploadUpdatedParams `json:"params"`
}

func (p ServerUploadUpdated) GetName() string { return "server.upload.updated" }

type serverUploadUpdatedParams struct {
	Uploads []Upload `json:"uploads"`
}
