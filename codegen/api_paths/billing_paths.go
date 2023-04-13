package api_paths

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/codegen/openapi/tags"
)

var BillingPaths = []PathSpec{
	{
		Path: "/api/v4/tariffs",
		Get: &OperationSpec{
			Response:    []tdproto.TariffBilling{},
			Description: "Get all tariffs.",
		},
		Tags: []string{tags.BillingTag},
	},
}
