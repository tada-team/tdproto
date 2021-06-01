package openapi

type Response struct {
	Description string   `json:"description"`
	Content     Contents `json:"content"`
}

type Responses struct {
	Status200 *Response `json:"200,omitempty"`
	Status404 *Response `json:"404,omitempty"`
	Status403 *Response `json:"403,omitempty"`
	Status422 *Response `json:"422,omitempty"`
}

func (r Responses) Iter() (res []Response)  {
	if r.Status200 != nil {
		res = append(res, *r.Status200)
	}
	if r.Status404 != nil {
		res = append(res, *r.Status404)
	}
	if r.Status403 != nil {
		res = append(res, *r.Status403)
	}
	if r.Status422 != nil {
		res = append(res, *r.Status422)
	}
	return
}
