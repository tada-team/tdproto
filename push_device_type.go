package tdproto

type PushDeviceType int

const (
	PushDeviceAndroid = PushDeviceType(1)
	PushDeviceIOS     = PushDeviceType(2)
	PushDeviceWeb     = PushDeviceType(3)
	PushDeviceSafari  = PushDeviceType(4)
	PushDeviceIOSFB   = PushDeviceType(5)
)

var MobilePushDevices = []PushDeviceType{
	PushDeviceAndroid,
	PushDeviceIOSFB,
	PushDeviceIOS,
}

var IOSPushDevices = []PushDeviceType{
	PushDeviceIOS,
	PushDeviceIOSFB,
}

var AndroidPushDevices = []PushDeviceType{
	PushDeviceAndroid,
}

func (t PushDeviceType) DbValue() int16 {
	return int16(t)
}

func (t PushDeviceType) String() string {
	switch t {
	case PushDeviceAndroid:
		return "Android"
	case PushDeviceIOS:
		return "iOS"
	case PushDeviceWeb:
		return "web"
	case PushDeviceSafari:
		return "safari"
	case PushDeviceIOSFB:
		return "iOS-firebase"
	default:
		return "?"
	}
}

func (t PushDeviceType) Mobile() bool {
	for _, v := range MobilePushDevices {
		if v == t {
			return true
		}
	}
	return false
}

func (t PushDeviceType) Name() string {
	switch t {
	case PushDeviceAndroid:
		return "android"
	case PushDeviceIOS:
		return "ios"
	case PushDeviceWeb:
		return "web"
	case PushDeviceSafari:
		return "safari"
	case PushDeviceIOSFB:
		return "iosfb"
	default:
		return "?"
	}
}
