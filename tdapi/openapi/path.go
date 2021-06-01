package openapi

type Path struct {
	Parameters []Parameter `json:"parameters,omitempty"`
	Get        *Operation  `json:"get,omitempty"`
	Post       *Operation  `json:"post,omitempty"`
	Put        *Operation  `json:"put,omitempty"`
	Delete     *Operation  `json:"delete,omitempty"`
}

func (p Path) Iter() (res []Operation) {
	if p.Get != nil {
		res = append(res, *p.Get)
	}
	if p.Post != nil {
		res = append(res, *p.Post)
	}
	if p.Put != nil {
		res = append(res, *p.Put)
	}
	if p.Delete != nil {
		res = append(res, *p.Delete)
	}
	return
}
