package tdproto

func NewServerCallCheckFingerprint(fingerprint string) (r ServerCallCheckFingerprint) {
	r.Name = r.GetName()
	r.Params.Fingerprint = fingerprint
	return r
}

type ServerCallCheckFingerprint struct {
	BaseEvent
	Params serverCallCheckFingerprintParams `json:"params"`
}

func (p ServerCallCheckFingerprint) GetName() string { return "server.call.checkfingerprint" }

type serverCallCheckFingerprintParams struct {
	Fingerprint string `json:"fingerprint"`
}
