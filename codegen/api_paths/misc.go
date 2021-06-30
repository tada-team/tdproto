package api_paths

import "github.com/tada-team/tdproto"

var MiscPaths = []PathSpec{
	{
		Path: "/api/v4/addr",
		Get: &OperationSpec{
			Response:            "",
			Description:         "Returns client address for debuging purposes.",
			ResponseDescription: "Address of the server.",
			SecurityIsOptional:  true,
		},
	},
	{
		Path: "/api/v4/ping",
		Get: &OperationSpec{
			Response:            "",
			Description:         "Ping the server.",
			ResponseDescription: "Set to ``\"pong\"``.",
			SecurityIsOptional:  true,
		},
	},
	{
		Path: "/features.json",
		Get: &OperationSpec{
			Response:           tdproto.Features{},
			Description:        "Get the server features information.",
			SecurityIsOptional: true,
		},
	},
}
