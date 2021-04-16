package openapi

type Path struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
}

func (ps Path) Operations() (res []*Operation) {
	if ps.Get != nil {
		res = append(res, ps.Get)
	}
	if ps.Post != nil {
		res = append(res, ps.Post)
	}
	if ps.Put != nil {
		res = append(res, ps.Put)
	}
	if ps.Delete != nil {
		res = append(res, ps.Delete)
	}
	return
}

