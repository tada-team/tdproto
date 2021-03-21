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
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// Call participants
	Onliners []CallOnliner `json:"onliners"`

	// Call start, if any
	Start *ISODateTimeString `json:"start"`

	// Call finish, if any
	Finish *ISODateTimeString `json:"finish"`

	// Call record enabled
	Audiorecord bool `json:"audiorecord"`

	// Call buzzing
	Buzz bool `json:"buzz,omitempty"`

	// Event start. FIXME: why not gentime?
	Timestamp int64 `json:"timestamp"`
}
