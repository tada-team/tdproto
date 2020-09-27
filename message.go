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

type MessageContent struct {
	Text          string       `json:"text"`
	Type          Mediatype    `json:"type"`
	Subtype       Mediasubtype `json:"subtype,omitempty"`
	Upload        string       `mediatype:"audiomsg,image,video,file" json:"upload,omitempty"`
	MediaUrl      string       `mediatype:"audiomsg,image,video,file" json:"mediaURL,omitempty"`
	Size          int          `mediatype:"audiomsg,image,video,file" json:"size,omitempty"`
	Duration      *uint        `mediatype:"audiomsg,video" json:"duration,omitempty"`
	Processing    bool         `mediatype:"video" json:"processing,omitempty"`
	PreviewHeight int          `mediatype:"image,video" json:"previewHeight,omitempty"`
	PreviewWidth  int          `mediatype:"image,video" json:"previewWidth,omitempty"`
	PreviewUrl    string       `mediatype:"image,video" json:"previewURL,omitempty"`
	Preview2xUrl  string       `mediatype:"image,video" json:"preview2xURL,omitempty"`
	Name          string       `mediatype:"image,video,file" json:"name,omitempty"`
	Animated      bool         `mediatype:"image" json:"animated,omitempty"`
	Title         string       `mediatype:"change" json:"title,omitempty"`
	New           *string      `mediatype:"change" json:"new,omitempty"`
	Old           *string      `mediatype:"change" json:"old,omitempty"`
	Actor         *JID         `mediatype:"change" json:"actor,omitempty"`
	Comment       string       `mediatype:"progress" json:"comment,omitempty"`
	I             uint         `mediatype:"progress" json:"i,omitempty"`
	Total         uint         `mediatype:"progress" json:"total,omitempty"`
	GivenName     *string      `mediatype:"contact" json:"given_name,omitempty"`
	FamilyName    *string      `mediatype:"contact" json:"family_name,omitempty"`
	Patronymic    *string      `mediatype:"contact" json:"patronymic,omitempty"`
	Phones        *[]string    `mediatype:"contact" json:"phones,omitempty"`
	Emails        *[]string    `mediatype:"contact" json:"emails,omitempty"`
	Stickerpack   string       `mediasubtype:"sticker" json:"stickerpack,omitempty"`
	PdfVersion    *PdfVersion  `json:"pdf_version,omitempty"`
	//Deadline      *time.Time   `mediasubtype:"newtask" json:"deadline,omitempty"`
}

type Message struct {
	// message content struct
	Content MessageContent `json:"content"`

	// simple plaintext message representation
	PushText string `json:"push_text,omitempty"`

	// sender jid
	From JID `json:"from"`

	// recipient jid
	To JID `json:"to"`

	// message uid
	MessageId string `json:"message_id"`

	// message creation datetime (set by server side)
	Created string `json:"created"`

	// object version
	Gentime int64 `json:"gentime"`

	// chat type
	ChatType ChatType `json:"chat_type"`

	// chat jid
	Chat JID `json:"chat,omitempty"`

	// external/internals links
	Links MessageLinks `json:"links,omitempty"`

	// importance flag
	Important bool `json:"important,omitempty"`

	// datetime of message modification or deletion
	Edited string `json:"edited,omitempty"`

	// message was seen by anybody in chat. True or null
	Received bool `json:"received,omitempty"`

	// unused yet
	NumReceived int `json:"num_received,omitempty"`

	// disable link previews. True or null
	Nopreview bool `json:"nopreview,omitempty"`

	// has link previews. True or null
	HasPreviews bool `json:"has_previews,omitempty"`

	// previous message id in this chat. Uid or null
	Prev string `json:"prev,omitempty"`

	// this message is first in this chat. True or null
	IsFirst bool `json:"is_first,omitempty"`

	// this message is first in this chat. True or null
	IsLast bool `json:"is_last,omitempty"`

	// message reactions struct. Can be null
	Reactions []MessageReaction `json:"reactions,omitempty"`

	// message that was replied to, if any
	ReplyTo *Message `json:"reply_to,omitempty"`

	// forwarded messages. Can be null. Also contains double of ReplyTo for backward compatibility
	LinkedMessages []Message `json:"linked_messages,omitempty"`

	// has mention (@). True or null
	Notice bool `json:"notice,omitempty"`

	// message has no pushes and did not affect any counters
	Silently bool `json:"silently,omitempty"`

	// author can change this message until date. Can be null
	EditableUntil string `json:"editable_until,omitempty"`

	// index number of this message. Starts from 0. Null for deleted messages. Changes when any previous message wad deleted.
	Num *int `json:"num,omitempty"`

	// debug information if any
	Debug string `json:"_debug,omitempty"`
}

type MessageLinkPreview struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type MessageLink struct {
	Pattern   string              `json:"pattern"`
	Url       string              `json:"url"`
	Text      string              `json:"text"`
	Preview   *MessageLinkPreview `json:"preview,omitempty"`
	Uploads   []Upload            `json:"uploads,omitempty"`
	NoPreview bool                `json:"nopreview,omitempty"`
	YoutubeId string              `json:"youtube_id,omitempty"`
}

type MessageLinks []MessageLink

func (links MessageLinks) Sort() {
	sort.Slice(links, func(i, j int) bool {
		return links[i].Pattern < links[j].Pattern
	})
}

type MessageReaction struct {
	Name    string                  `json:"name"`
	Counter int                     `json:"counter"`
	Details []MessageReactionDetail `json:"details"`
}

type MessageReactionDetail struct {
	Created string `json:"created"`
	Sender  *JID   `json:"sender"`
	Name    string `json:"name"`
}
