package tdproto

func NewServerChatDraft(jid JID, draft string, gentime int64) (r ServerChatDraft) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Draft = draft
	r.Params.Gentime = gentime
	return r
}

// Changed draft message in chat
type ServerChatDraft struct {
	BaseEvent
	Params serverChatDraftParams `json:"params"`
}

func (p ServerChatDraft) GetName() string { return "server.chat.draft" }

type serverChatDraftParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Draft version
	Gentime int64 `json:"gentime" tdproto:"readonly"`

	// Draft text
	Draft string `json:"draft"`

	// Deprecated
	DraftNum int64 `json:"draft_num"`
}
