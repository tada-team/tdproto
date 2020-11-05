package tdapi

type Paginator struct {
	Limit  int `schema:"limit"`
	Offset int `schema:"offset"`
}

type UserParams struct {
	Lang     string `schema:"lang"`
	Timezone string `schema:"timezone"`
}
