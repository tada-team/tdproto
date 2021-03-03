package tdproto

type UploadMediaType int16

const (
	MediaTypeFile  = UploadMediaType(0)
	MediaTypeImage = UploadMediaType(1)
	MediaTypeVideo = UploadMediaType(2)
	MediaTypeAudio = UploadMediaType(3)
)

func (u UploadMediaType) DbValue() int16 {
	return int16(u)
}

func (u UploadMediaType) String() string {
	switch u {
	case MediaTypeFile:
		return "file"
	case MediaTypeImage:
		return "image"
	case MediaTypeVideo:
		return "video"
	case MediaTypeAudio:
		return "audio"
	default:
		return ""
	}
}

func (u UploadMediaType) Ref() *UploadMediaType {
	return &u
}
