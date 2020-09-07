package tdproto

//type ShortWikiPage struct {
//	Updated string `json:"updated"`
//	Editor  JID    `json:"editor"`
//}

type WikiPage struct {
	Gentime int64  `json:"gentime"`
	Updated string `json:"updated"`
	Editor  JID    `json:"editor"`
	Text    string `json:"text"`
}
