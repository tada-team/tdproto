package tdws

import "github.com/tada-team/tdproto"

func NewServerChatDraft(jid tdproto.JID, draft string, gentime int64) (r ServerChatDraft) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Draft = draft
	r.Params.DraftGentime = gentime
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
	Jid tdproto.JID `json:"jid"`

	// Draft text
	Draft string `json:"draft"`

	// Draft version
	DraftGentime int64 `json:"draft_gentime"`

	// Deprecated
	DraftNum int64 `json:"draft_num"`
}
