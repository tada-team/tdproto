package tdproto

import "time"

func NewServerCallState(chat HasJid, startCall, finishCall *time.Time, onliners []CallOnliner, audiorecord, buzz bool, uid string, timestamp int64) (r ServerCallState) {
	r.Name = r.GetName()
	r.Params.Jid = *chat.JID()
	r.Params.Onliners = onliners
	r.Params.Buzz = buzz
	r.Params.Uid = uid
	r.Params.Audiorecord = audiorecord
	r.Params.Timestamp = timestamp

	if startCall != nil {
		r.Params.Start = NullableIsoDatetime(startCall)
	} else {
		r.Params.Buzz = true
	}

	if finishCall != nil {
		r.Params.Finish = NullableIsoDatetime(startCall)
	}

	return r
}

// Call participant number or parameters changed
type ServerCallState struct {
	BaseEvent
	Params serverCallStateParams `json:"params"`
}

func (p ServerCallState) GetName() string { return "server.call.state" }

type serverCallStateParams struct {
	Jid         JID           `json:"jid"`
	Onliners    []CallOnliner `json:"onliners"`
	Start       *string       `json:"start"`
	Finish      *string       `json:"finish"`
	Audiorecord bool          `json:"audiorecord"`
	Buzz        bool          `json:"buzz,omitempty"`
	Uid         string        `json:"uid"`
	Timestamp   int64         `json:"timestamp"`
}
