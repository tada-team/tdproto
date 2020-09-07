package tdproto

type Country struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Default bool   `json:"default,omitempty"`
	Popular bool   `json:"popular,omitempty"`
}
