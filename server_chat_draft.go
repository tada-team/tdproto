package tdproto

func NewServerChatDraft(jid JID, draft string, gentime int64, revision int64) (r ServerChatDraft) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Draft = draft
	r.Params.DraftGentime = gentime
	r.Params.Revision = revision
	return r
}

// Changed draft message in chat
type ServerChatDraft struct {
	BaseEvent
	Params serverChatDraftParams `json:"params"`
}

func (p ServerChatDraft) GetName() string { return "server.chat.draft" }

// Params of the server.chat.draft event
type serverChatDraftParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Draft text
	Draft string `json:"draft"`

	// Draft version
	// Deprecated: use Revision instead
	DraftGentime int64 `json:"draft_gentime"`

	// Revision Unixtime(ms)
	Revision int64 `json:"revision"`

	// Deprecated: use Revision instead
	DraftNum int64 `json:"draft_num"`
}
