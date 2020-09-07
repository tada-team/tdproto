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
	Params ServerChatDraftParams `json:"params"`
}

type ServerChatDraftParams struct {
	Jid      *JID   `json:"jid"`
	Draft    string `json:"draft"`
	DraftNum int64  `json:"draft_num"`
}
