package tdproto

func NewServerCallCheckFingerprint(fingerprint string) (r ServerCallCheckFingerprint) {
	r.BaseEvent.Name = "server.call.checkfingerprint"
	r.Params.Fingerprint = fingerprint
	return r
}

type ServerCallCheckFingerprint struct {
	BaseEvent
	Params serverCallCheckFingerprintParams `json:"params"`
}

type serverCallCheckFingerprintParams struct {
	Fingerprint string `json:"fingerprint"`
}
