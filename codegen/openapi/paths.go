package main

import (
	"github.com/tada-team/tdproto"
)

var TdPaths = map[string]OpenApiPath{
	"/ping": {
		Get: &OpenApiOperation{
			Summary: "Ping server",
			Responses: map[string]OpenApiResponse{
				"200": createResponseRefJson(tdproto.ClientPing{}),
			},
		},
	},
}
