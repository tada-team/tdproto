package tdproto

// CallType ...
type CallType string

const (
	// CallTypeAudio is a audio room
	CallTypeAudio CallType = "audio"

	// CallTypeVideo is a video room
	CallTypeVideo CallType = "video"

	// CallTypeVideoMultistream is a video room in multistream mode
	CallTypeVideoMultistream CallType = "video_multistream"

	// CallTypeIVCS is a vcs using IVA
	CallTypeIVCS CallType = "ivcs"
)

func (ct CallType) IsVideo() bool { return ct == CallTypeVideo || ct == CallTypeVideoMultistream }
func (ct CallType) IsAudio() bool { return ct == CallTypeAudio }

func (ct CallType) String() string { return string(ct) }
