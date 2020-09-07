package tdproto

type CallEvent struct {
	Start       *string       `json:"start"`
	Finish      *string       `json:"finish"`
	Audiorecord bool          `json:"audiorecord"`
	Onliners    []CallOnliner `json:"onliners"`
}

type CallOnliner struct {
	Jid         JID          `json:"jid"`
	Muted       bool         `json:"muted"`
	Devices     []CallDevice `json:"devices"`
	DisplayName string       `json:"display_name"`
	Icon        string       `json:"icon"`
}

type CallDevice struct {
	Muted     bool   `json:"muted"`
	Useragent string `json:"useragent"`
}
