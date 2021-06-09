package api_paths

import "github.com/tada-team/tdproto"

var MiscPaths = []PathSpec{
	{
		Path: "/api/v4/addr",
		Get: &OperationSpec{
			Responce:            "",
			Description:         "Returns client address for debuging purposes.",
			ResponceDescription: "Address of the server.",
			SecurityIsOptional:  true,
		},
	},
	{
		Path: "/api/v4/ping",
		Get: &OperationSpec{
			Responce:            "",
			Description:         "Ping the server.",
			ResponceDescription: "Set to ``\"pong\"``.",
			SecurityIsOptional:  true,
		},
	},
	{
		Path: "/features.json",
		Get: &OperationSpec{
			Responce:           tdproto.Features{},
			Description:        "Get the server features information.",
			SecurityIsOptional: true,
		},
	},
}
