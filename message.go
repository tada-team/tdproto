package tdproto

import "sort"

type Mediatype string

const (
	MediatypePlain    Mediatype = "plain"
	MediatypeChange   Mediatype = "change"
	MediatypeDeleted  Mediatype = "deleted"
	MediatypeFile     Mediatype = "file"
	MediatypeImage    Mediatype = "image"
	MediatypeVideo    Mediatype = "video"
	MediatypeAudiomsg Mediatype = "audiomsg"
	MediatypeContact  Mediatype = "contact"
	MediatypePdf      Mediatype = "pdf"
	//MediatypeNewtask  Mediatype = "newtask"
	//MediatypeProgress Mediatype = "progress"
)

type Mediasubtype string

const (
	MediaSubtypeSticker Mediasubtype = "sticker"
	MediaSubtypeNewtask Mediasubtype = "newtask"
	//MediaSubtypeSpeech  Mediasubtype = "speech"
)

// Chat message content
type MessageContent struct {
	// Text repesentation of message
	Text string `json:"text"`

	// Message type
	Type Mediatype `json:"type"`

	// Message subtype, if any
	Subtype Mediasubtype `json:"subtype,omitempty"`

	// Upload id, if any
	Upload string `mediatype:"audiomsg,image,video,file" json:"upload,omitempty"`

	// Upload url, if any
	MediaUrl string `mediatype:"audiomsg,image,video,file" json:"mediaURL,omitempty"`

	// Upload size, if any
	Size int `mediatype:"audiomsg,image,video,file" json:"size,omitempty"`

	// Upload duration, if any
	Duration *uint `mediatype:"audiomsg,video" json:"duration,omitempty"`

	// Upload stil processing, if any
	Processing bool `mediatype:"video" json:"processing,omitempty"`

	// Upload preview height, in pixels, if any
	PreviewHeight int `mediatype:"image,video" json:"previewHeight,omitempty"`

	// Upload width, in pixels, if any
	PreviewWidth int `mediatype:"image,video" json:"previewWidth,omitempty"`

	// Upload preview absolute url, if any
	PreviewUrl string `mediatype:"image,video" json:"previewURL,omitempty"`

	// Upload high resolution preview absolute url, if any
	Preview2xUrl string `mediatype:"image,video" json:"preview2xURL,omitempty"`

	// Upload name, if any
	Name string `mediatype:"image,video,file" json:"name,omitempty"`

	// Upload is animated image, if any
	Animated bool `mediatype:"image" json:"animated,omitempty"`

	// Change title (for "change" mediatype)
	Title string `mediatype:"change" json:"title,omitempty"`

	// Change old value (for "change" mediatype)
	Old *string `mediatype:"change" json:"old,omitempty"`

	// Change new value (for "change" mediatype)
	New *string `mediatype:"change" json:"new,omitempty"`

	// Change actor contact id (for "change" mediatype)
	Actor *JID `mediatype:"change" json:"actor,omitempty"`

	// Comment. For audimessage.
	Comment string `mediatype:"progress" json:"comment,omitempty"`

	// Given name (for "contact"  mediatype)
	GivenName *string `mediatype:"contact" json:"given_name,omitempty"`

	// Family name (for "contact"  mediatype)
	FamilyName *string `mediatype:"contact" json:"family_name,omitempty"`

	// Patronymic name (for "contact"  mediatype)
	Patronymic *string `mediatype:"contact" json:"patronymic,omitempty"`

	// Contact phones list (for "contact"  mediatype)
	Phones *[]string `mediatype:"contact" json:"phones,omitempty"`

	// Emails list (for "contact"  mediatype)
	Emails *[]string `mediatype:"contact" json:"emails,omitempty"`

	// Stickerpack name (for "sticker" subtype)
	Stickerpack string `mediasubtype:"sticker" json:"stickerpack,omitempty"`

	// Pdf version, if any
	PdfVersion *PdfVersion `json:"pdf_version,omitempty"`

	//Deadline *time.Time   `mediasubtype:"newtask" json:"deadline,omitempty"`
}

// Chat message
type Message struct {
	// Message content struct
	Content MessageContent `json:"content"`

	// Simple plaintext message representation
	PushText string `json:"push_text,omitempty" tdproto:"readonly"`

	// Sender contact id
	From JID `json:"from" tdproto:"readonly"`

	// Recipient id (group, task or contact)
	To JID `json:"to"`

	// Message uid
	MessageId string `json:"message_id"`

	// Message creation datetime (set by server side)
	Created string `json:"created" tdproto:"readonly"`

	// Object version
	Gentime int64 `json:"gentime" tdproto:"readonly"`

	// Chat type
	ChatType ChatType `json:"chat_type" tdproto:"readonly"`

	// Chat id
	Chat JID `json:"chat" tdproto:"readonly"`

	// External/internals links
	Links MessageLinks `json:"links,omitempty" tdproto:"readonly"`

	// Markup entities. Experimental
	Markup []MarkupEntity `json:"markup,omitempty" tdproto:"readonly"`

	// Importance flag
	Important bool `json:"important,omitempty"`

	// Datetime of message modification or deletion
	Edited string `json:"edited,omitempty" tdproto:"readonly"`

	// Message was seen by anybody in chat. True or null
	Received bool `json:"received,omitempty" tdproto:"readonly"`

	// Unused yet
	NumReceived int `json:"num_received,omitempty" tdproto:"readonly"`

	// Disable link previews. True or null
	Nopreview bool `json:"nopreview,omitempty"`

	// Has link previews. True or null
	HasPreviews bool `json:"has_previews,omitempty" tdproto:"readonly"`

	// Previous message id in this chat. Uid or null
	Prev string `json:"prev,omitempty" tdproto:"readonly"`

	// This message is first in this chat. True or null
	IsFirst bool `json:"is_first,omitempty" tdproto:"readonly"`

	// This message is first in this chat. True or null
	IsLast bool `json:"is_last,omitempty" tdproto:"readonly"`

	// Message reactions struct. Can be null
	Reactions []MessageReaction `json:"reactions,omitempty" tdproto:"readonly"`

	// Message that was replied to, if any
	ReplyTo *Message `json:"reply_to,omitempty"`

	// Forwarded messages. Can be null. Also contains double of ReplyTo for backward compatibility
	LinkedMessages []Message `json:"linked_messages,omitempty"`

	// Has mention (@). True or null
	Notice bool `json:"notice,omitempty" tdproto:"readonly"`

	// Message has no pushes and did not affect any counters
	Silently bool `json:"silently,omitempty" tdproto:"readonly"`

	// Author can change this message until date. Can be null
	EditableUntil string `json:"editable_until,omitempty" tdproto:"readonly"`

	// Index number of this message. Starts from 0. Null for deleted messages. Changes when any previous message wad deleted.
	Num *int `json:"num,omitempty" tdproto:"readonly"`

	// Debug information, if any
	Debug string `json:"_debug,omitempty" tdproto:"readonly"`
}

// Website title and description
type MessageLinkPreview struct {
	// Website title or og:title content
	Title string `json:"title"`

	// Website description
	Description string `json:"description,omitempty"`
}

// Checked message links. In short: "Click here: {link.Pattern}" => "Click here: <a href='{link.Url}'>{link.Text}</a>"
type MessageLink struct {
	// Text fragment that should be replaced by link
	Pattern string `json:"pattern"`

	// Internal (tadateam://) or external link
	Url string `json:"url"`

	// Text replacement.
	Text string `json:"text"`

	// Optional preview info, for websites
	Preview *MessageLinkPreview `json:"preview,omitempty"`

	// Optional upload info
	Uploads []Upload `json:"uploads,omitempty"`

	// Website previews disabled
	NoPreview bool `json:"nopreview,omitempty"`

	// Optional youtube movie id
	YoutubeId string `json:"youtube_id,omitempty"`
}

type MessageLinks []MessageLink

func (links MessageLinks) Sort() {
	sort.Slice(links, func(i, j int) bool {
		return links[i].Pattern < links[j].Pattern
	})
}

// Message emoji reaction
type MessageReaction struct {
	// Emoji
	Name string `json:"name"`

	// Number of reactions
	Counter int `json:"counter"`

	// Details
	Details []MessageReactionDetail `json:"details"`
}

// Message reaction detail
type MessageReactionDetail struct {
	// When reaction added, iso datetime
	Created string `json:"created"`

	// Reaction author
	Sender *JID `json:"sender"`

	// Reaction emoji
	Name string `json:"name"`
}
