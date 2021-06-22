package tdproto

// JavaScript Session Establishment Protocol
type JSEP struct {
	// Session Description Protocol information
	SDP string `json:"sdp"`
	// See https://rtcweb-wg.github.io/jsep/#rfc.section.4.1.8
	Type string `json:"type"`
}
