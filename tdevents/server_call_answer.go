package tdevents

import "github.com/tada-team/tdproto"

func NewServerCallAnswer(jid tdproto.JID, sdp string, onliners []tdproto.CallOnliner, uid string) (r ServerCallAnswer) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.JSEP.Type = "answer"
	r.Params.JSEP.SDP = sdp
	r.Params.Onliners = onliners
	r.Params.Uid = uid
	return r
}

// Call parameters (deprecated: use `ServerCallSdp`)
type ServerCallAnswer struct {
	BaseEvent
	Params serverCallAnswerParams `json:"params"`
}

func (p ServerCallAnswer) GetName() string { return "server.call.answer" }

// Params of the server.call.answer event
type serverCallAnswerParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// List of ICE candidates (when trickle = false)
	Candidates []serverCallAnswerCandidate `json:"candidates,omitempty"`

	// Current call participants
	Onliners []tdproto.CallOnliner `json:"onliners,omitempty"`

	// SDP data
	JSEP tdproto.JSEP `json:"jsep"`

	// Call id
	Uid string `json:"uid"`
}

// ICE candidate for call answer
type serverCallAnswerCandidate = struct {
	Candidate     string `json:"candidate"`
	SdpMLineIndex int    `json:"sdpMLineIndex"`
}
