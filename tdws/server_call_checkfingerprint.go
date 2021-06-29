package tdws

func NewServerCallCheckFingerprint(fingerprint string) (r ServerCallCheckFingerprint) {
	r.Name = r.GetName()
	r.Params.Fingerprint = fingerprint
	return r
}

// Experimental function
type ServerCallCheckFingerprint struct {
	BaseEvent
	Params serverCallCheckFingerprintParams `json:"params"`
}

func (p ServerCallCheckFingerprint) GetName() string { return "server.call.checkfingerprint" }

// Params of the server.call.checkfingerprint event
type serverCallCheckFingerprintParams struct {
	Fingerprint string `json:"fingerprint"`
}
