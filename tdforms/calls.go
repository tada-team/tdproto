package tdforms

import "github.com/tada-team/tdproto"

type ClientCallTrickle struct {
	Jid           tdproto.JID `json:"jid"`
	Candidate     string      `json:"candidate"`
	SdpMid        string      `json:"sdp_mid"`
	SdpMlineIndex int         `json:"sdp_mline_index"`
}

type ClientCallRecord struct {
	Jid         tdproto.JID `json:"jid"`
	Uid         string      `json:"uid"`
	Audiorecord bool        `json:"audiorecord,omitempty"`
}

type ClientCallSound struct {
	Jid   tdproto.JID `json:"jid"`
	Muted bool        `json:"muted"`
}

type ClientCallMuteAll struct {
	Jid tdproto.JID `json:"jid"`
}

type ClientCallReject struct {
	Jid    tdproto.JID `json:"jid"`
	Reason string      `json:"reason"`
}

type ClientCallLeave struct {
	Jid    tdproto.JID `json:"jid"`
	Reason string      `json:"reason"`
}
