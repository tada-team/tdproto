package tdproto

// CallType ...
type CallType string

const (
	// CallTypeAudio is a audio room
	CallTypeAudio CallType = "audio"

	// CallTypeVideo is a video room
	CallTypeVideo CallType = "video"
)

func (ct CallType) IsVideo() bool { return ct == CallTypeVideo }
func (ct CallType) IsAudio() bool { return ct == CallTypeAudio }

func (ct CallType) String() string { return string(ct) }
