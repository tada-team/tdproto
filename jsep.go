package tdproto

type SDPType string

const (
	AnswerType SDPType = "answer"
	OfferType  SDPType = "offer"
)

type JSEP struct {
	SDP  string  `json:"sdp"`
	Type SDPType `json:"type"`
}
