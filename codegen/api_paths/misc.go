package api_paths

var MiscPaths = []PathSpec{
	{
		Path: "/api/v4/addr",
		Get: &HttpSpec{
			Responce:            "",
			Description:         "Returns client address for debuging purposes.",
			ResponceDescription: "Address of the server.",
		},
	},
}
