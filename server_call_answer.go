package tdproto

func NewServerCallAnswer(jid JID, sdp string, onliners []CallOnliner, uid string) (r ServerCallAnswer) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.JSEP.Type = "answer"
	r.Params.JSEP.SDP = sdp
	r.Params.Onliners = onliners
	r.Params.Uid = uid
	return r
}

// Call parameters
type ServerCallAnswer struct {
	BaseEvent
	Params serverCallAnswerParams `json:"params"`
}

func (p ServerCallAnswer) GetName() string { return "server.call.answer" }

type serverCallAnswerParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// List of ICE candidates (when trickle = false)
	Candidates []serverCallAnswerCandidate `json:"candidates"`

	// Current call participants
	Onliners []CallOnliner `json:"onliners"`

	// SDP data
	JSEP JSEP `json:"jsep"`

	// Call id
	Uid string `json:"uid"`
}

type JSEP struct {
	SDP  string `json:"sdp"`
	Type string `json:"type"`
}

type serverCallAnswerCandidate = struct {
	Candidate     string `json:"candidate"`
	SdpMLineIndex int    `json:"sdpMLineIndex"`
}
