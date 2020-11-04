package tdapi

type Paginator struct {
	Limit  int `schema:"limit"`
	Offset int `schema:"offset"`
}
