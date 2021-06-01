package openapi

type Content struct {
	Schema Schema `json:"schema"`
}

type Contents struct {
	ApplicationJSON               *Content `json:"application/json,omitempty"`
	ApplicationXWWWFormUrlencoded *Content `json:"application/x-www-form-urlencoded,omitempty"`
}

func (c Contents) Iter() (res []Content) {
	if c.ApplicationJSON != nil {
		res = append(res, *c.ApplicationJSON)
	}
	if c.ApplicationXWWWFormUrlencoded != nil {
		res = append(res, *c.ApplicationXWWWFormUrlencoded)
	}
	return
}
