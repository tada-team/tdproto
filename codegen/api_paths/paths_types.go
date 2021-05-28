package api_paths

type HttpSpec struct {
	Request             interface{}
	Responce            interface{}
	Description         string
	ResponceDescription string
}

type PathSpec struct {
	Get    *HttpSpec
	Put    *HttpSpec
	Delete *HttpSpec
	Post   *HttpSpec
}
