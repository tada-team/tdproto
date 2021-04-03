package tdproto

//type ShortWikiPage struct {
//	Updated string `json:"updated"`
//	Editor  JID    `json:"editor"`
//}

// Wiki page. Experimental
type WikiPage struct {
	// Object version
	Gentime int64 `json:"gentime"`

	// Update time
	Updated ISODateTimeString `json:"updated"`

	// Last editor contact id
	Editor JID `json:"editor"`

	// Page text
	Text string `json:"text"`
}
