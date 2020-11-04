package tdapi

type Paginator struct {
	Limit  int `schema:"limit"`
	Offset int `schema:"offset"`
}

type UserParams struct {
	Lang     int `schema:"lang"`
	Timezone int `schema:"timezone"`
}
