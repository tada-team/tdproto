package tdproto

func NewServerCallCheckFingerprint(fingerprint string) (r ServerCallCheckFingerprint) {
	r.BaseEvent.Name = "server.call.checkfingerprint"
	r.Params.Fingerprint = fingerprint
	return r
}

type ServerCallCheckFingerprint struct {
	BaseEvent
	Params ServerCallCheckFingerprintParams `json:"params"`
}

type ServerCallCheckFingerprintParams struct {
	Fingerprint string `json:"fingerprint"`
}
