package tdproto

type UploadMediaType string

const (
	MediaTypeFile      UploadMediaType = "file"
	MediaTypeImage     UploadMediaType = "image"
	MediaTypeVideo     UploadMediaType = "video"
	MediaTypeAudio     UploadMediaType = "audio"
	MediaTypeFileImage UploadMediaType = "imagefile"
)

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
	case MediaTypeFileImage:
		return "fileimage"
	default:
		return ""
	}
}
