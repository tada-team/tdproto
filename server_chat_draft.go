package tdproto

func NewServerChatDraft(jid *JID, draft string, draftNum int64) (r ServerChatDraft) {
	r.BaseEvent.Name = "server.chat.draft"
	r.Params.Jid = jid
	r.Params.Draft = draft
	r.Params.DraftNum = draftNum
	return r
}

type ServerChatDraft struct {
	BaseEvent
	Params serverChatDraftParams `json:"params"`
}

type serverChatDraftParams struct {
	Jid      *JID   `json:"jid"`
	Draft    string `json:"draft"`
	DraftNum int64  `json:"draft_num"`
}
