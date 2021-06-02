package api_paths

import "github.com/tada-team/tdproto"

var MiscPaths = []PathSpec{
	{
		Path: "/api/v4/addr",
		Get: &HttpSpec{
			Responce:            "",
			Description:         "Returns client address for debuging purposes.",
			ResponceDescription: "Address of the server.",
		},
	},
	{
		Path: "/api/v4/ping",
		Get: &HttpSpec{
			Responce:            "",
			Description:         "Ping the server.",
			ResponceDescription: "Set to ``\"pong\"``.",
		},
	},
	{
		Path: "/features.json",
		Get: &HttpSpec{
			Responce:    tdproto.Features{},
			Description: "Get the server features information.",
		},
	},
}
