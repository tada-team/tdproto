package tdproto

type Tag struct {
	Name string `json:"name"`
	Uid  string `json:"uid"`
}

type DeletedTag struct {
	Uid string `json:"uid"`
}
