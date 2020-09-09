package tdproto

func NewServerCallAnswer(jid JID, sdp string, onliners []CallOnliner, uid string) (r ServerCallAnswer) {
	r.Name = "server.call.answer"
	r.Params.Jid = jid
	r.Params.JSEP.Type = "answer"
	r.Params.JSEP.SDP = sdp
	r.Params.Onliners = onliners
	r.Params.Uid = uid
	return r
}

type ServerCallAnswer struct {
	BaseEvent
	Params serverCallAnswerParams `json:"params"`
}

type serverCallAnswerParams struct {
	Jid        JID                         `json:"jid"`
	Candidates []serverCallAnswerCandidate `json:"candidates"`
	Onliners   []CallOnliner               `json:"onliners"`
	JSEP       JSEP                        `json:"jsep"`
	Uid        string                      `json:"uid"`
}

type JSEP struct {
	SDP  string `json:"sdp"`
	Type string `json:"type"`
}

type serverCallAnswerCandidate = struct {
	Candidate     string `json:"candidate"`
	SdpMLineIndex int    `json:"sdpMLineIndex"`
}
